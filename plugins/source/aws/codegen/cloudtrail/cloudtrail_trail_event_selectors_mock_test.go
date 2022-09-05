// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package cloudtrail

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func buildCloudtrailTrailEventSelectors(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudtrailClient(ctrl)

	item := types.EventSelector{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []types.EventSelector{item},
		}, nil)

	return client.Services{
		Cloudtrail: mock,
	}
}

func TestCloudtrailTrailEventSelectors(t *testing.T) {
	client.AwsMockTestHelper(t, CloudtrailTrailEventSelectors(), buildCloudtrailTrailEventSelectors, client.TestOptions{})
}