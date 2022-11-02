// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func NetworkAcls() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_network_acls",
		Description: "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html",
		Resolver:    fetchEc2NetworkAcls,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveNetworkAclArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "associations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Associations"),
			},
			{
				Name:     "entries",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Entries"),
			},
			{
				Name:     "is_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDefault"),
			},
			{
				Name:     "network_acl_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkAclId"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
