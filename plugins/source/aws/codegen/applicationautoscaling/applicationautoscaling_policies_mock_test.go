// Code generated by codegen; DO NOT EDIT.

package applicationautoscaling

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
)

func buildApplicationAutoscalingPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockApplicationAutoscalingClient(ctrl)

	item := types.ScalingPolicy{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeScalingPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&applicationautoscaling.DescribeScalingPoliciesOutput{
			ScalingPolicies: []types.ScalingPolicy{item},
		}, nil)

	return client.Services{
		ApplicationAutoscaling: mock,
	}
}

func TestApplicationAutoscalingPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ApplicationAutoscalingPolicies(), buildApplicationAutoscalingPolicies, client.TestOptions{})
}
