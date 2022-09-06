// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package cloudtrail

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

func buildCloudtrailTrails(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudtrailClient(ctrl)

	item := types.Trail{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&cloudtrail.DescribeTrailsOutput{
			TrailList: []types.Trail{item},
		}, nil)

	return client.Services{
		Cloudtrail: mock,
	}
}

func TestCloudtrailTrails(t *testing.T) {
	client.MockTestHelper(t, CloudtrailTrails(), buildCloudtrailTrails, client.TestOptions{})
}
