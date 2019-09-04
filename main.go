// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-09-04 17:16:07.577221396 -0500 CDT m=+0.000380580
package azure_resources

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

var authorizer autorest.Authorizer

type resource interface {
	GetProperties() ([]byte, error)
}

func GetAllByGroupName(subscriptionID, groupName string) []resource {
	client := resources.NewClient(subscriptionID)
	client.Authorizer = authorizer

	results, err := client.ListByResourceGroup(context.Background(), groupName, "", "", nil)
	if err != nil {
		log.Fatalln(err)
	}

	var resources []resource

	for _, resource := range results.Values() {
		switch *resource.Type {
		case "Microsoft.Compute/virtualMachines":
			r := &ComputeVirtualMachines{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Compute/virtualMachines/Extensions":
			r := &ComputeVirtualMachineExtensions{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Network/publicIPAddresses":
			r := &NetworkPublicIPAddresses{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Compute/disks":
			r := &ComputeDisks{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)

		case "Microsoft.Network/virtualNetworks":
			r := &NetworkVirtualNetworks{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)
		}
	}

	return resources
}

func SetAuthorizer(tenantID, clientID, clientSecret string) error {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return err
	}
	spToken, err := adal.NewServicePrincipalToken(*oauthConfig, clientID, clientSecret, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return err
	}
	authorizer = autorest.NewBearerAuthorizer(spToken)
	return nil
}
