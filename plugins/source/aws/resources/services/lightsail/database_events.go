// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DatabaseEvents() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_database_events",
		Resolver:  fetchLightsailDatabaseEvents,
		Multiplex: client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "database_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "event_categories",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EventCategories"),
			},
			{
				Name:     "message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Message"),
			},
			{
				Name:     "resource",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Resource"),
			},
		},
	}
}