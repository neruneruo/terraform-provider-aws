// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package licensemanager

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	licensemanager_sdkv2 "github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDistributedGrants,
			TypeName: "aws_licensemanager_grants",
		},
		{
			Factory:  DataSourceReceivedLicense,
			TypeName: "aws_licensemanager_received_license",
		},
		{
			Factory:  DataSourceReceivedLicenses,
			TypeName: "aws_licensemanager_received_licenses",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAssociation,
			TypeName: "aws_licensemanager_association",
			Name:     "Association",
		},
		{
			Factory:  ResourceGrant,
			TypeName: "aws_licensemanager_grant",
		},
		{
			Factory:  ResourceGrantAccepter,
			TypeName: "aws_licensemanager_grant_accepter",
		},
		{
			Factory:  ResourceLicenseConfiguration,
			TypeName: "aws_licensemanager_license_configuration",
			Name:     "License Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LicenseManager
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*licensemanager_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return licensemanager_sdkv2.NewFromConfig(cfg,
		licensemanager_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
