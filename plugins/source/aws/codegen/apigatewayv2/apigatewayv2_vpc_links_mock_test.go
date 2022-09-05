// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package apigatewayv2

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func buildApigatewayv2VpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockApigatewayv2Client(ctrl)

	item := types.VpcLink{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&apigatewayv2.GetVpcLinksOutput{
			Items: []types.VpcLink{item},
		}, nil)

	return client.Services{
		Apigatewayv2: mock,
	}
}

func TestApigatewayv2VpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, Apigatewayv2VpcLinks(), buildApigatewayv2VpcLinks, client.TestOptions{})
}