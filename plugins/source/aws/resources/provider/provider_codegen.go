// Code generated by codegen using template provider.go.tpl; DO NOT EDIT.

package provider

import (
	"embed"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/module"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/accessanalyzer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/acm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/apigatewayv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/applicationautoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudfront"
)

var (
	//go:embed moduledata/*
	moduleData embed.FS

	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:             "aws",
		Version:          Version,
		Configure:        client.Configure,
		ErrorClassifier:  client.ErrorClassifier,
		ModuleInfoReader: module.EmbeddedReader(moduleData, "moduledata"),
		ResourceMap: map[string]*schema.Table{
			//"iot.security_profiles": 				 iot.IotSecurityProfiles(), //TODO disabled because of api error NotFoundException: No method found matching route security-profiles for http method GET.

			"accessanalyzer.analyzers":                accessanalyzer.AccessAnalyzerAccessanalyzers(),
			"acm.certificates":                        acm.ACMCertificates(),
			"apigatewayv2.apis":                       apigatewayv2.Apigatewayv2Apis(),
			"apigatewayv2.domain_names":               apigatewayv2.Apigatewayv2DomainNames(),
			"apigatewayv2.vpc_links":                  apigatewayv2.Apigatewayv2VpcLinks(),
			"applicationautoscaling.scaling_policies": applicationautoscaling.ApplicationAutoscalingPolicies(),
			"appsync.graphql_apis":                    appsync.AppSyncGraphqlApis(),
			"athena.data_catalogs":                    athena.AthenaDataCatalogs(),
			"athena.work_groups":                      athena.AthenaWorkGroups(),
			"autoscaling.launch_configurations":       autoscaling.AutoscalingLaunchConfigurations(),
			"autoscaling.scheduled_actions":           autoscaling.AutoscalingScheduledActions(),
			"autoscaling.auto_scaling_groups":         autoscaling.AutoscalingGroups(),
			"backup.global_settings":                  backup.BackupGlobalSettings(),
			"backup.region_settings":                  backup.BackupRegionSettings(),
			"backup.backup_vaults":                    backup.BackupVaults(),
			"backup.backup_plans":                     backup.BackupBackupPlans(),
			"cloudformation.stacks":                   cloudformation.CloudformationStacks(),
			"cloudfront.cache_policies":               cloudfront.CloudfrontCachePolicies(),
			"cloudfront.distributions":                cloudfront.CloudfrontDistributions(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}