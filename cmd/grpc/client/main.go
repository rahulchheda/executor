package main

import (
	cmd "executor/cmd"
	//"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	cmd.Execute()
}

/*import (
	"context"
	"fmt"
	//"github.com/litmuschaos/go-util/internal/grpc"
	"github.com/litmuschaos/go-util/internal/grpc/domain"
	"github.com/litmuschaos/go-util/internal/grpc/service"
	"google.golang.org/grpc"
	//"log"
	//"net"
	//"github.com/spf13/cobra"
)

func main() {
	//cobra.OnInitialize(initConfig)
	serverAddress := "localhost:7000"
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}

	defer conn.Close()

	client := service.NewExecutorServiceClient(conn)

	s := domain.Name{
		FileName: "patch.go",
	}

	if responseMessage, e := client.Execute(context.Background(), &s); e != nil {
		panic(fmt.Sprintln("File Executed.."))
		fmt.Println(responseMessage)
		fmt.Println("========================================")
	}
}
*/
