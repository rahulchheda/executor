package impl

import (
	"context"
	"github.com/litmuschaos/go-util/internal/grpc/domain"
	"github.com/litmuschaos/go-util/internal/grpc/service"
	"log"
)

type ExecuteFileServiceGrpc struct{}

func NewExecuteFileServiceGrpc() *ExecuteFileServiceGrpc {
	return &ExecuteFileServiceGrpc{}
}

func (serviceExecute *ExecuteFileServiceGrpc) Execute(ctx context.Context, in *domain.Name) (*service.Status, error) {
	log.Println("Recieved Request for executing file with name" + in.FileName)
	log.Println("Execute the file vefore this and fill up the status")

	return &service.Status{
		CompletedStatus: true,
		AnyError:        nil,
	}, nil
}
