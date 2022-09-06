// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package cloudfront

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

func buildCloudfrontCachePolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudfrontClient(ctrl)

	item := types.CachePolicySummary{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListCachePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&cloudfront.ListCachePoliciesOutput{
			Items: []types.CachePolicySummary{item},
		}, nil)

	return client.Services{
		Cloudfront: mock,
	}
}

func TestCloudfrontCachePolicies(t *testing.T) {
	client.MockTestHelper(t, CloudfrontCachePolicies(), buildCloudfrontCachePolicies, client.TestOptions{})
}
