//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_ListByProfile.json
func ExampleSecretsClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecretsClient().NewListByProfilePager("RG", "profile1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SecretListResult = armcdn.SecretListResult{
		// 	Value: []*armcdn.Secret{
		// 		{
		// 			Name: to.Ptr("secret1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/secrets"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret1"),
		// 			Properties: &armcdn.SecretProperties{
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 				Parameters: &armcdn.CustomerCertificateParameters{
		// 					Type: to.Ptr(armcdn.SecretTypeCustomerCertificate),
		// 					CertificateAuthority: to.Ptr("Symantec"),
		// 					ExpirationDate: to.Ptr("2025-01-01T00:00:00-00:00"),
		// 					SecretSource: &armcdn.ResourceReference{
		// 						ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vaults/keyvaultname/secrets/certificatename"),
		// 					},
		// 					SecretVersion: to.Ptr("abcdef1234578900abcdef1234567890"),
		// 					Subject: to.Ptr("*.contoso.com"),
		// 					SubjectAlternativeNames: []*string{
		// 						to.Ptr("foo.contoso.com"),
		// 						to.Ptr("www3.foo.contoso.com")},
		// 						Thumbprint: to.Ptr("ABCDEF1234567890ABCDEF1234567890ABCDEF12"),
		// 						UseLatestVersion: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("69f05517-6aaf-4a1e-a96d-f8c02f66c756-test12-afdx-test-domains-azfdtest-xyz"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/secrets"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/autogenerated_secret_name"),
		// 				Properties: &armcdn.SecretProperties{
		// 					DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 					ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 					Parameters: &armcdn.ManagedCertificateParameters{
		// 						Type: to.Ptr(armcdn.SecretTypeManagedCertificate),
		// 						ExpirationDate: to.Ptr("2025-01-01T00:00:00-00:00"),
		// 						Subject: to.Ptr("bar.contoso.com"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_Get.json
func ExampleSecretsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecretsClient().Get(ctx, "RG", "profile1", "secret1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Secret = armcdn.Secret{
	// 	Name: to.Ptr("secret1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/secrets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret1"),
	// 	Properties: &armcdn.SecretProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		Parameters: &armcdn.CustomerCertificateParameters{
	// 			Type: to.Ptr(armcdn.SecretTypeCustomerCertificate),
	// 			CertificateAuthority: to.Ptr("Symantec"),
	// 			ExpirationDate: to.Ptr("2025-01-01T00:00:00-00:00"),
	// 			SecretSource: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vaults/keyvaultname/secrets/certificatename"),
	// 			},
	// 			SecretVersion: to.Ptr("abcdef1234578900abcdef1234567890"),
	// 			Subject: to.Ptr("*.contoso.com"),
	// 			SubjectAlternativeNames: []*string{
	// 				to.Ptr("foo.contoso.com"),
	// 				to.Ptr("www3.foo.contoso.com")},
	// 				Thumbprint: to.Ptr("ABCDEF1234567890ABCDEF1234567890ABCDEF12"),
	// 				UseLatestVersion: to.Ptr(true),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_Create.json
func ExampleSecretsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecretsClient().BeginCreate(ctx, "RG", "profile1", "secret1", armcdn.Secret{
		Properties: &armcdn.SecretProperties{
			Parameters: &armcdn.CustomerCertificateParameters{
				Type: to.Ptr(armcdn.SecretTypeCustomerCertificate),
				SecretSource: &armcdn.ResourceReference{
					ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename"),
				},
				SecretVersion:    to.Ptr("abcdef1234578900abcdef1234567890"),
				UseLatestVersion: to.Ptr(false),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Secret = armcdn.Secret{
	// 	Name: to.Ptr("secret1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/secrets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret1"),
	// 	Properties: &armcdn.SecretProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		Parameters: &armcdn.CustomerCertificateParameters{
	// 			Type: to.Ptr(armcdn.SecretTypeCustomerCertificate),
	// 			CertificateAuthority: to.Ptr("Symantec"),
	// 			ExpirationDate: to.Ptr("2025-01-01T00:00:00-00:00"),
	// 			SecretSource: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vaults/keyvaultname/secrets/certificatename"),
	// 			},
	// 			SecretVersion: to.Ptr("abcdef1234578900abcdef1234567890"),
	// 			Subject: to.Ptr("*.contoso.com"),
	// 			SubjectAlternativeNames: []*string{
	// 				to.Ptr("foo.contoso.com"),
	// 				to.Ptr("www3.foo.contoso.com")},
	// 				Thumbprint: to.Ptr("ABCDEF1234567890ABCDEF1234567890ABCDEF12"),
	// 				UseLatestVersion: to.Ptr(true),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_Delete.json
func ExampleSecretsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecretsClient().BeginDelete(ctx, "RG", "profile1", "secret1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
