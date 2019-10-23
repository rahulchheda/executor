package main

import (
	//"context"
	"fmt"
	//"github.com/litmuschaos/go-util/internal/grpc"
	"executor/internal/grpc/domain"
	"executor/internal/grpc/service"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
)

type ExecuteFileServiceGrpc struct{}

func (serviceExecute *ExecuteFileServiceGrpc) Execute(ctx context.Context, in *domain.Name) (*service.Status, error) {
	log.Println("Recieved Request for executing file with name" + in.FileName)
	log.Println("Execute the file vefore this and fill up the status")

	response := &service.Status{
		CompletedStatus: true,
		AnyError:        nil,
	}
	return response, nil
}
func getNetListerner(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))

	}

	return lis
}

func main() {
	netListener := getNetListerner(7000)
	opts := []grpc.ServerOption{}
	gRPCServer := grpc.NewServer(opts...)
	executor := &ExecuteFileServiceGrpc{}
	service.RegisterExecutorServiceServer(gRPCServer, executor)
	go func() {
		if err := gRPCServer.Serve(netListener); err != nil {
			log.Fatalf("Failed to Serve: %v", err)
		}
	}()
	fmt.Println("gRPC In Action!")

	// Bad way to stop the server
	// if err := s.Serve(listener); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
	// Right way to stop the server using a SHUTDOWN HOOK

	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed
	// And our main routine keeps running
	// If the main routine were to shutdown so would the child routine that is Serving the server
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	gRPCServer.Stop()
	netListener.Close()
}
