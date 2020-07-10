package containerinstance

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ContainerClient is the client for the Container methods of the Containerinstance service.
type ContainerClient struct {
	BaseClient
}

// NewContainerClient creates an instance of the ContainerClient client.
func NewContainerClient(subscriptionID string) ContainerClient {
	return NewContainerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerClientWithBaseURI creates an instance of the ContainerClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewContainerClientWithBaseURI(baseURI string, subscriptionID string) ContainerClient {
	return ContainerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ExecuteCommand executes a command for a specific container instance in a specified resource group and container
// group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
// containerExecRequest - the request for the exec command.
func (client ContainerClient) ExecuteCommand(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (result ContainerExecResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerClient.ExecuteCommand")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecuteCommandPreparer(ctx, resourceGroupName, containerGroupName, containerName, containerExecRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ExecuteCommand", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteCommandSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ExecuteCommand", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteCommandResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ExecuteCommand", resp, "Failure responding to request")
	}

	return
}

// ExecuteCommandPreparer prepares the ExecuteCommand request.
func (client ContainerClient) ExecuteCommandPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/exec", pathParameters),
		autorest.WithJSON(containerExecRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteCommandSender sends the ExecuteCommand request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerClient) ExecuteCommandSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExecuteCommandResponder handles the response to the ExecuteCommand request. The method always
// closes the http.Response Body.
func (client ContainerClient) ExecuteCommandResponder(resp *http.Response) (result ContainerExecResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLogs get the logs for a specified container instance in a specified resource group and container group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
// tail - the number of lines to show from the tail of the container instance log. If not provided, all
// available logs are shown up to 4mb.
func (client ContainerClient) ListLogs(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (result Logs, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerClient.ListLogs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListLogsPreparer(ctx, resourceGroupName, containerGroupName, containerName, tail)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ListLogs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLogsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ListLogs", resp, "Failure sending request")
		return
	}

	result, err = client.ListLogsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerClient", "ListLogs", resp, "Failure responding to request")
	}

	return
}

// ListLogsPreparer prepares the ListLogs request.
func (client ContainerClient) ListLogsPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if tail != nil {
		queryParameters["tail"] = autorest.Encode("query", *tail)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListLogsSender sends the ListLogs request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerClient) ListLogsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListLogsResponder handles the response to the ListLogs request. The method always
// closes the http.Response Body.
func (client ContainerClient) ListLogsResponder(resp *http.Response) (result Logs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
