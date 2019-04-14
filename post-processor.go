package main

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/svchost/disco"
	"github.com/hashicorp/terraform/terraform"

	tfBackendLocal "github.com/hashicorp/terraform/backend/local"
	tfBackendRemote "github.com/hashicorp/terraform/backend/remote"

	tfBackendAzure "github.com/hashicorp/terraform/backend/remote-state/azure"
	tfBackendConsul "github.com/hashicorp/terraform/backend/remote-state/consul"
	tfBackendEtcdv3 "github.com/hashicorp/terraform/backend/remote-state/etcdv3"
	tfBackendGcs "github.com/hashicorp/terraform/backend/remote-state/gcs"
	tfBackendInmem "github.com/hashicorp/terraform/backend/remote-state/inmem"
	tfBackendManta "github.com/hashicorp/terraform/backend/remote-state/manta"
	tfBackendS3 "github.com/hashicorp/terraform/backend/remote-state/s3"
	tfBackendSwift "github.com/hashicorp/terraform/backend/remote-state/swift"

	tfConfig "github.com/hashicorp/terraform/config"
)

var (
	irregularBackends = map[string]bool{
		"atlas": true,
	}
)

// Config is the post-processor configuration with interpolation supported.
type Config struct {
	TerraformFilePath string `mapstructure:"terraform_file_path"`
	State             string `mapstructure:"state"`
	ctx               interpolate.Context
}

func (c *Config) handlingTerraformConfig() bool {
	return filepath.Base(c.TerraformFilePath) == "terraform.tf"
}

func (c *Config) validate() (*tfConfig.Config, error) {
	if !c.handlingTerraformConfig() {
		return nil, fmt.Errorf("Only terraform.tf files asre supported")
	}
	if c.TerraformFilePath == "" {
		return nil, fmt.Errorf("Terraform file path missing. terraform_file_path not set?")
	}
	cfg, err := tfConfig.LoadFile(c.TerraformFilePath)
	if err != nil {
		return nil, fmt.Errorf("Terraform configuration file not found at: '%s'", c.TerraformFilePath)
	}
	return cfg, nil
}

// PostProcessor holds the Config object.
type PostProcessor struct {
	config          Config
	terraformConfig *tfConfig.Config
	terraformState  string
}

func (p *PostProcessor) configureBackend(b backend.Backend, c map[string]interface{}) (backend.Backend, error) {
	rc, err := tfConfig.NewRawConfig(c)
	if err != nil {
		return backend.Nil{}, err
	}
	conf := terraform.NewResourceConfig(rc)
	_, errs := b.Validate(conf)
	if len(errs) > 0 {
		return backend.Nil{}, fmt.Errorf("Error while configuring Terraform backend: '%+v'", errs)
	}
	err = b.Configure(conf)
	if err != nil {
		return backend.Nil{}, err
	}
	return b, nil
}

func (p *PostProcessor) getBackend(bt string) (backend.Backend, error) {
	backends := map[string]interface{}{
		"azure":  func() backend.Backend { return tfBackendAzure.New() },
		"consul": func() backend.Backend { return tfBackendConsul.New() },
		"etcdv3": func() backend.Backend { return tfBackendEtcdv3.New() },
		"gcs":    func() backend.Backend { return tfBackendGcs.New() },
		"inmem":  func() backend.Backend { return tfBackendInmem.New() },
		"local":  func() backend.Backend { return tfBackendLocal.New() },
		"manta":  func() backend.Backend { return tfBackendManta.New() },
		"remote": func() backend.Backend { return tfBackendRemote.New(disco.New()) },
		"s3":     func() backend.Backend { return tfBackendS3.New() },
		"swift":  func() backend.Backend { return tfBackendSwift.New() },
	}
	b, ok := backends[bt]
	if !ok {
		return backend.Nil{}, fmt.Errorf("Unknown backend type '%s'", bt)
	}
	return b.(func() backend.Backend)(), nil
}

// Configure sets the Config object with configuration values from the Packer
// template.
func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)

	if err != nil {
		return err
	}

	cfg, err := p.config.validate()
	if err != nil {
		return err
	}

	if cfg.Terraform != nil {
		if cfg.Terraform.Backend != nil {

			errs := cfg.Terraform.Backend.Validate()
			if len(errs) > 0 {
				return fmt.Errorf("Error while validating the Terraform backend configuration: '%+v'", errs)
			}

			p.terraformConfig = cfg
			p.terraformState = "default"
			if p.config.State != "" {
				p.terraformState = p.config.State
			}

		} else {
			return fmt.Errorf("Failed to load Terraform backend from terraform.tf file")
		}
	} else {
		return fmt.Errorf("Failed to load Terraform backend from terraform.tf file")
	}

	return nil
}

// PostProcess parses the AMI ID from the artifact ID, and then passes the AMI ID
// to UpdateJSONFile to be set as the new value of the JSON paths properties in
// Packer template.
// AWS artifact ID output has the format of <region>:<ami_id>,
// for example: ap-southeast-2:ami-4f8fae2c
func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {

	ui.Say(fmt.Sprintf("%s", artifact.String()))
	ui.Say(fmt.Sprintf("Builder ID ========> %s", artifact.BuilderId()))

	rawMap := p.terraformConfig.Terraform.Backend.RawConfig.RawMap()

	if _, ok := irregularBackends[p.terraformConfig.Terraform.Backend.Type]; !ok {

		b, err := p.getBackend(p.terraformConfig.Terraform.Backend.Type)
		if err != nil {
			return artifact, false, err
		}
		configuredBackend, err := p.configureBackend(b, rawMap)
		if err != nil {
			return artifact, false, err
		}
		state, err := configuredBackend.State(p.terraformState)
		if err != nil {
			return artifact, false, err
		}
		state.RefreshState()
		state.State().RootModule().Outputs[artifact.BuilderId()] = &terraform.OutputState{
			Sensitive: false,
			Type:      "string",
			Value:     artifact.Id(),
		}

		if persistErr := state.PersistState(); persistErr != nil {
			return artifact, false, persistErr
		}

	} else {
		if p.terraformConfig.Terraform.Backend.Type == "atlas" {
		} else {
			return artifact, false, fmt.Errorf("Backend type '%s' not supported", p.terraformConfig.Terraform.Backend.Type)
		}
	}

	return artifact, true, nil
}
