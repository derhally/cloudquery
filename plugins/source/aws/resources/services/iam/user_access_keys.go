// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserAccessKeys() *schema.Table {
	return &schema.Table{
		Name:                 "aws_iam_user_access_keys",
		Resolver:             fetchIamUserAccessKeys,
		PostResourceResolver: postIamUserAccessKeyResolver,
		Multiplex:            client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("user_id"),
			},
			{
				Name:     "access_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKeyId"),
			},
			{
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
			{
				Name:     "last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastRotated"),
			},
		},
	}
}