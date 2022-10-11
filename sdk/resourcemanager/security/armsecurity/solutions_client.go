//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// SolutionsClient contains the methods for the SecuritySolutions group.
// Don't use this type directly, use NewSolutionsClient() instead.
type SolutionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSolutionsClient creates a new instance of SolutionsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSolutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SolutionsClient, error) {
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
	client := &SolutionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a specific Security Solution.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// securitySolutionName - Name of security solution.
// options - SolutionsClientGetOptions contains the optional parameters for the SolutionsClient.Get method.
func (client *SolutionsClient) Get(ctx context.Context, resourceGroupName string, ascLocation string, securitySolutionName string, options *SolutionsClientGetOptions) (SolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ascLocation, securitySolutionName, options)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, securitySolutionName string, options *SolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/securitySolutions/{securitySolutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if securitySolutionName == "" {
		return nil, errors.New("parameter securitySolutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySolutionName}", url.PathEscape(securitySolutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SolutionsClient) getHandleResponse(resp *http.Response) (SolutionsClientGetResponse, error) {
	result := SolutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Security Solutions for the subscription.
// Generated from API version 2020-01-01
// options - SolutionsClientListOptions contains the optional parameters for the SolutionsClient.List method.
func (client *SolutionsClient) NewListPager(options *SolutionsClientListOptions) *runtime.Pager[SolutionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SolutionsClientListResponse]{
		More: func(page SolutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SolutionsClientListResponse) (SolutionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SolutionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SolutionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SolutionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SolutionsClient) listCreateRequest(ctx context.Context, options *SolutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SolutionsClient) listHandleResponse(resp *http.Response) (SolutionsClientListResponse, error) {
	result := SolutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionList); err != nil {
		return SolutionsClientListResponse{}, err
	}
	return result, nil
}