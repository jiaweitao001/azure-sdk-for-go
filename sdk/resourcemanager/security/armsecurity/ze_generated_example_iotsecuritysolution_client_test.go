//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHub.json
func ExampleIotSecuritySolutionClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armsecurity.IotSecuritySolutionClientListBySubscriptionOptions{Filter: to.StringPtr("<filter>")})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByRg.json
func ExampleIotSecuritySolutionClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armsecurity.IotSecuritySolutionClientListByResourceGroupOptions{Filter: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<solution-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotSecuritySolutionClientGetResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/CreateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<solution-name>",
		armsecurity.IoTSecuritySolutionModel{
			Tags:     map[string]*string{},
			Location: to.StringPtr("<location>"),
			Properties: &armsecurity.IoTSecuritySolutionProperties{
				DisabledDataSources: []*armsecurity.DataSource{},
				DisplayName:         to.StringPtr("<display-name>"),
				Export:              []*armsecurity.ExportData{},
				IotHubs: []*string{
					to.StringPtr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")},
				RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
					{
						RecommendationType: armsecurity.RecommendationType("IoT_OpenPorts").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					},
					{
						RecommendationType: armsecurity.RecommendationType("IoT_SharedCredentials").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					}},
				Status:                  armsecurity.SecuritySolutionStatus("Enabled").ToPtr(),
				UnmaskedIPLoggingStatus: armsecurity.UnmaskedIPLoggingStatus("Enabled").ToPtr(),
				UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
					Query: to.StringPtr("<query>"),
					QuerySubscriptions: []*string{
						to.StringPtr("075423e9-7d33-4166-8bdf-3920b04e3735")},
				},
				Workspace: to.StringPtr("<workspace>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotSecuritySolutionClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<solution-name>",
		armsecurity.UpdateIotSecuritySolutionData{
			Tags: map[string]*string{
				"foo": to.StringPtr("bar"),
			},
			Properties: &armsecurity.UpdateIoTSecuritySolutionProperties{
				RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
					{
						RecommendationType: armsecurity.RecommendationType("IoT_OpenPorts").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					},
					{
						RecommendationType: armsecurity.RecommendationType("IoT_SharedCredentials").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					}},
				UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
					Query: to.StringPtr("<query>"),
					QuerySubscriptions: []*string{
						to.StringPtr("075423e9-7d33-4166-8bdf-3920b04e3735")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotSecuritySolutionClientUpdateResult)
}

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/DeleteIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<solution-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}