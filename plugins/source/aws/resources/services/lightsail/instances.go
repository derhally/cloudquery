// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_instances",
		Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Instance.html",
		Resolver:    fetchLightsailInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "access_details",
				Type:     schema.TypeJSON,
				Resolver: resolveLightsailInstanceAccessDetails,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "add_ons",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AddOns"),
			},
			{
				Name:     "blueprint_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BlueprintId"),
			},
			{
				Name:     "blueprint_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BlueprintName"),
			},
			{
				Name:     "bundle_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BundleId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "hardware",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Hardware"),
			},
			{
				Name:     "ip_address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddressType"),
			},
			{
				Name:     "ipv6_addresses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Ipv6Addresses"),
			},
			{
				Name:     "is_static_ip",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsStaticIp"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "networking",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Networking"),
			},
			{
				Name:     "private_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateIpAddress"),
			},
			{
				Name:     "public_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIpAddress"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "ssh_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SshKeyName"),
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Username"),
			},
		},

		Relations: []*schema.Table{
			InstancePortStates(),
		},
	}
}
