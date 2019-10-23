This Repo is just the starting to developing the go-based executor
What will I implement in this!
 - gRPC server
 - Connecting it to Cobra CLI
 - executing those cobra commands in the executor ansible!# go-util

How to use it?
 - compile and run the main.go in cmd/grpc/server
 - It will run the gRPC server at localhost:7000
 - After that do a `go install` command in the root of this repo
 - Then use the CLI command `executor execute --filename {file_name}`
 - Check the file name in the grpc server log
