package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ClusterUpgradesClient is the client for the ClusterUpgrades methods of the Servicefabric service.
type ClusterUpgradesClient struct {
	BaseClient
}

// NewClusterUpgradesClient creates an instance of the ClusterUpgradesClient client.
func NewClusterUpgradesClient(timeout *int32) ClusterUpgradesClient {
	return NewClusterUpgradesClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewClusterUpgradesClientWithBaseURI creates an instance of the ClusterUpgradesClient client.
func NewClusterUpgradesClientWithBaseURI(baseURI string, timeout *int32) ClusterUpgradesClient {
	return ClusterUpgradesClient{NewWithBaseURI(baseURI, timeout)}
}

// Resume resume cluster upgrades
//
// resumeClusterUpgrade is the upgrade of the cluster
func (client ClusterUpgradesClient) Resume(ctx context.Context, resumeClusterUpgrade ResumeClusterUpgrade) (result String, err error) {
	req, err := client.ResumePreparer(ctx, resumeClusterUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Resume", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Resume", resp, "Failure sending request")
		return
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Resume", resp, "Failure responding to request")
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client ClusterUpgradesClient) ResumePreparer(ctx context.Context, resumeClusterUpgrade ResumeClusterUpgrade) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/MoveToNextUpgradeDomain"),
		autorest.WithJSON(resumeClusterUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterUpgradesClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client ClusterUpgradesClient) ResumeResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Rollback rollback cluster upgrades
func (client ClusterUpgradesClient) Rollback(ctx context.Context) (result String, err error) {
	req, err := client.RollbackPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Rollback", nil, "Failure preparing request")
		return
	}

	resp, err := client.RollbackSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Rollback", resp, "Failure sending request")
		return
	}

	result, err = client.RollbackResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Rollback", resp, "Failure responding to request")
	}

	return
}

// RollbackPreparer prepares the Rollback request.
func (client ClusterUpgradesClient) RollbackPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/RollbackUpgrade"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RollbackSender sends the Rollback request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterUpgradesClient) RollbackSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RollbackResponder handles the response to the Rollback request. The method always
// closes the http.Response Body.
func (client ClusterUpgradesClient) RollbackResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start start cluster upgrades
//
// startClusterUpgrade is the description of the start cluster upgrade
func (client ClusterUpgradesClient) Start(ctx context.Context, startClusterUpgrade StartClusterUpgrade) (result String, err error) {
	req, err := client.StartPreparer(ctx, startClusterUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client ClusterUpgradesClient) StartPreparer(ctx context.Context, startClusterUpgrade StartClusterUpgrade) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/Upgrade"),
		autorest.WithJSON(startClusterUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterUpgradesClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ClusterUpgradesClient) StartResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update cluster upgrades
//
// updateClusterUpgrade is the description of the update cluster upgrade
func (client ClusterUpgradesClient) Update(ctx context.Context, updateClusterUpgrade UpdateClusterUpgrade) (result String, err error) {
	req, err := client.UpdatePreparer(ctx, updateClusterUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterUpgradesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ClusterUpgradesClient) UpdatePreparer(ctx context.Context, updateClusterUpgrade UpdateClusterUpgrade) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/UpdateUpgrade"),
		autorest.WithJSON(updateClusterUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterUpgradesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClusterUpgradesClient) UpdateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
