package aadapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/aad/mgmt/2020-07-01-preview/aad"
	"github.com/Azure/go-autorest/autorest"
)

// AzureADMetricsClientAPI contains the set of methods on the AzureADMetricsClient type.
type AzureADMetricsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, azureADMetricsName string, azureADMetricsConfig aad.AzureADMetricsConfig) (result aad.AzureADMetricsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, azureADMetricsName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, azureADMetricsName string) (result aad.AzureADMetricsConfig, err error)
	List(ctx context.Context, resourceGroupName string) (result aad.AzureADMetricsListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result aad.AzureADMetricsListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result aad.AzureADMetricsListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result aad.AzureADMetricsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, azureADMetricsName string, azureADMetricsConfig *aad.AzureADMetricsUpdateParameter) (result aad.AzureADMetricsConfig, err error)
}

var _ AzureADMetricsClientAPI = (*aad.AzureADMetricsClient)(nil)

// PrivateLinkForAzureAdClientAPI contains the set of methods on the PrivateLinkForAzureAdClient type.
type PrivateLinkForAzureAdClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy aad.PrivateLinkPolicy) (result aad.PrivateLinkForAzureAdCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result aad.PrivateLinkPolicyListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result aad.PrivateLinkPolicyListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result aad.PrivateLinkPolicyListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result aad.PrivateLinkPolicyListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy *aad.PrivateLinkPolicyUpdateParameter) (result aad.PrivateLinkPolicy, err error)
}

var _ PrivateLinkForAzureAdClientAPI = (*aad.PrivateLinkForAzureAdClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, policyName string, groupName string) (result aad.PrivateLinkResource, err error)
	ListByPrivateLinkPolicy(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkResourceListResultPage, err error)
	ListByPrivateLinkPolicyComplete(ctx context.Context, resourceGroupName string, policyName string) (result aad.PrivateLinkResourceListResultIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*aad.PrivateLinkResourcesClient)(nil)