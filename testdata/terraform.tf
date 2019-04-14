terraform {
  backend "s3" {
    bucket = "xxxk8s.terraform.state"
    key    = "packer-terraform-state-post-processor/test.tfstate"
    region = "eu-central-1"
    profile = "rad-admin"
  }
}