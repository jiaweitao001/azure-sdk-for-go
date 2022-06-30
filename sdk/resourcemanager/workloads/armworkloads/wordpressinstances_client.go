//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WordpressInstancesClient contains the methods for the WordpressInstances group.
// Don't use this type directly, use NewWordpressInstancesClient() instead.
type WordpressInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWordpressInstancesClient creates a new instance of WordpressInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWordpressInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WordpressInstancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &WordpressInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or updated WordPress instance resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// phpWorkloadName - Php workload name
// wordpressInstanceResource - Resource create or update request payload
// options - WordpressInstancesClientBeginCreateOrUpdateOptions contains the optional parameters for the WordpressInstancesClient.BeginCreateOrUpdate
// method.
func (client *WordpressInstancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, phpWorkloadName string, wordpressInstanceResource WordpressInstanceResource, options *WordpressInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[WordpressInstancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, phpWorkloadName, wordpressInstanceResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WordpressInstancesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WordpressInstancesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or updated WordPress instance resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *WordpressInstancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, phpWorkloadName string, wordpressInstanceResource WordpressInstanceResource, options *WordpressInstancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, phpWorkloadName, wordpressInstanceResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WordpressInstancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, phpWorkloadName string, wordpressInstanceResource WordpressInstanceResource, options *WordpressInstancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/phpWorkloads/{phpWorkloadName}/wordpressInstances/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if phpWorkloadName == "" {
		return nil, errors.New("parameter phpWorkloadName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{phpWorkloadName}", url.PathEscape(phpWorkloadName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, wordpressInstanceResource)
}

// Delete - Delete WordPress instance resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// phpWorkloadName - Php workload name
// options - WordpressInstancesClientDeleteOptions contains the optional parameters for the WordpressInstancesClient.Delete
// method.
func (client *WordpressInstancesClient) Delete(ctx context.Context, resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientDeleteOptions) (WordpressInstancesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, phpWorkloadName, options)
	if err != nil {
		return WordpressInstancesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WordpressInstancesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WordpressInstancesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WordpressInstancesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WordpressInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/phpWorkloads/{phpWorkloadName}/wordpressInstances/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if phpWorkloadName == "" {
		return nil, errors.New("parameter phpWorkloadName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{phpWorkloadName}", url.PathEscape(phpWorkloadName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the WordPress instance resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// phpWorkloadName - Php workload name
// options - WordpressInstancesClientGetOptions contains the optional parameters for the WordpressInstancesClient.Get method.
func (client *WordpressInstancesClient) Get(ctx context.Context, resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientGetOptions) (WordpressInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, phpWorkloadName, options)
	if err != nil {
		return WordpressInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WordpressInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WordpressInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WordpressInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/phpWorkloads/{phpWorkloadName}/wordpressInstances/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if phpWorkloadName == "" {
		return nil, errors.New("parameter phpWorkloadName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{phpWorkloadName}", url.PathEscape(phpWorkloadName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WordpressInstancesClient) getHandleResponse(resp *http.Response) (WordpressInstancesClientGetResponse, error) {
	result := WordpressInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WordpressInstanceResource); err != nil {
		return WordpressInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists WordPress instance resources under a phpWorkload resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// phpWorkloadName - Php workload name
// options - WordpressInstancesClientListOptions contains the optional parameters for the WordpressInstancesClient.List method.
func (client *WordpressInstancesClient) NewListPager(resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientListOptions) *runtime.Pager[WordpressInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WordpressInstancesClientListResponse]{
		More: func(page WordpressInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WordpressInstancesClientListResponse) (WordpressInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, phpWorkloadName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WordpressInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WordpressInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WordpressInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WordpressInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, phpWorkloadName string, options *WordpressInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/phpWorkloads/{phpWorkloadName}/wordpressInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if phpWorkloadName == "" {
		return nil, errors.New("parameter phpWorkloadName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{phpWorkloadName}", url.PathEscape(phpWorkloadName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WordpressInstancesClient) listHandleResponse(resp *http.Response) (WordpressInstancesClientListResponse, error) {
	result := WordpressInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WordpressInstanceResourceList); err != nil {
		return WordpressInstancesClientListResponse{}, err
	}
	return result, nil
}
