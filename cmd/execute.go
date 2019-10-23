/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"executor/internal/grpc/domain"
	"executor/internal/grpc/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute the file",
	Long: `This CMD is used for executing the file with a name

	If no blog post is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, Args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}
		req := domain.Name{
			FileName: filename,
		}
		serverAddress := "localhost:7000"
		conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
		if e != nil {
			panic(e)
		}
		client := service.NewExecutorServiceClient(conn)
		res, err := client.Execute(context.Background(), &req)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	executeCmd.Flags().String("filename", "", "The name of the file")
	executeCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(executeCmd)
}
