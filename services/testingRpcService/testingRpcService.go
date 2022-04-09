package testingRpcService

import (
	"consumerskripsi/constans"
	testingRpc "consumerskripsi/repositories/testingRpc"
	"consumerskripsi/services"
	"context"
	"fmt"
	"log"
)

type testingRpcService struct {
	Service services.UsecaseService
}

func NewTestingRpcService(service services.UsecaseService) testingRpcService {
	return testingRpcService{
		Service: service,
	}
}

func (svc *testingRpcService) GetUserById(ctx context.Context, input *testingRpc.UserRequest) (*testingRpc.UserResult, error) {
	log.Println("Incoming Request GetUserById:", input)
	metaResponse := testingRpc.UserResult{
		Meta: &testingRpc.MetaResponse{
			Success: true,
			Message: constans.EMPTY_VALUE,
		},
	}

	fmt.Println(input)

	return &metaResponse, nil

}
