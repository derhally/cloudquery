// Code generated by codegen using template resource_list_describe_mock_test.go.tpl; DO NOT EDIT.

package cloudfront

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func buildCloudfrontDistributions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudfrontClient(ctrl)

	var item types.DistributionSummary
	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDistributions(
		gomock.Any(),
		&cloudfront.ListDistributionsInput{},
		gomock.Any(),
	).Return(
		&cloudfront.ListDistributionsOutput{
			// &types.DistributionList
			DistributionList: &types.DistributionList{
				Items: []types.DistributionSummary{item},
			},
		},
		nil,
	)

	var detail types.Distribution
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}

	detail.Id = item.Id

	mock.EXPECT().GetDistribution(
		gomock.Any(),
		&cloudfront.GetDistributionInput{

			Id: item.Id,
		},
		gomock.Any(),
	).Return(

		&cloudfront.GetDistributionOutput{
			Distribution: &detail,
		},

		nil,
	)

	return client.Services{
		Cloudfront: mock,
	}
}

func TestCloudfrontDistributions(t *testing.T) {
	client.AwsMockTestHelper(t, CloudfrontDistributions(), buildCloudfrontDistributions, client.TestOptions{})
}