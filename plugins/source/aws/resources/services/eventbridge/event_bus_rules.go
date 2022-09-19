// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventBusRules() *schema.Table {
	return &schema.Table{
		Name:      "aws_eventbridge_event_bus_rules",
		Resolver:  fetchEventbridgeEventBusRules,
		Multiplex: client.ServiceAccountRegionMultiplexer("events"),
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
				Name:     "event_bus_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEventbridgeEventBusRuleTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "event_bus_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventBusName"),
			},
			{
				Name:     "event_pattern",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventPattern"),
			},
			{
				Name:     "managed_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagedBy"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "schedule_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduleExpression"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
		},
	}
}