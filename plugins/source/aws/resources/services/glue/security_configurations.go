// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityConfigurations() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_security_configurations",
		Resolver:  fetchGlueSecurityConfigurations,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time_stamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTimeStamp"),
			},
			{
				Name:     "encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionConfiguration"),
			},
		},
	}
}
