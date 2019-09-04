// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
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
			{{- range .Resources }}
			case "{{ . }}":
			r := &{{ GetType . }}{{ GetResourceTitle . }}{
				subscriptionID: subscriptionID,
				groupName:      groupName,
				resourceName:   resource.Name,
				authorizer:     authorizer,
			}
			resources = append(resources, r)
			{{ end -}}
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
