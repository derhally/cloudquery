// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

//go:generate mockgen -package=mocks -destination=../mocks/xray.go -source=xray.go XrayClient
type XrayClient interface {
	BatchGetTraces(context.Context, *xray.BatchGetTracesInput, ...func(*xray.Options)) (*xray.BatchGetTracesOutput, error)
	GetEncryptionConfig(context.Context, *xray.GetEncryptionConfigInput, ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error)
	GetGroup(context.Context, *xray.GetGroupInput, ...func(*xray.Options)) (*xray.GetGroupOutput, error)
	GetGroups(context.Context, *xray.GetGroupsInput, ...func(*xray.Options)) (*xray.GetGroupsOutput, error)
	GetInsight(context.Context, *xray.GetInsightInput, ...func(*xray.Options)) (*xray.GetInsightOutput, error)
	GetInsightEvents(context.Context, *xray.GetInsightEventsInput, ...func(*xray.Options)) (*xray.GetInsightEventsOutput, error)
	GetInsightImpactGraph(context.Context, *xray.GetInsightImpactGraphInput, ...func(*xray.Options)) (*xray.GetInsightImpactGraphOutput, error)
	GetInsightSummaries(context.Context, *xray.GetInsightSummariesInput, ...func(*xray.Options)) (*xray.GetInsightSummariesOutput, error)
	GetSamplingRules(context.Context, *xray.GetSamplingRulesInput, ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error)
	GetSamplingStatisticSummaries(context.Context, *xray.GetSamplingStatisticSummariesInput, ...func(*xray.Options)) (*xray.GetSamplingStatisticSummariesOutput, error)
	GetSamplingTargets(context.Context, *xray.GetSamplingTargetsInput, ...func(*xray.Options)) (*xray.GetSamplingTargetsOutput, error)
	GetServiceGraph(context.Context, *xray.GetServiceGraphInput, ...func(*xray.Options)) (*xray.GetServiceGraphOutput, error)
	GetTimeSeriesServiceStatistics(context.Context, *xray.GetTimeSeriesServiceStatisticsInput, ...func(*xray.Options)) (*xray.GetTimeSeriesServiceStatisticsOutput, error)
	GetTraceGraph(context.Context, *xray.GetTraceGraphInput, ...func(*xray.Options)) (*xray.GetTraceGraphOutput, error)
	GetTraceSummaries(context.Context, *xray.GetTraceSummariesInput, ...func(*xray.Options)) (*xray.GetTraceSummariesOutput, error)
	ListTagsForResource(context.Context, *xray.ListTagsForResourceInput, ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error)
}
