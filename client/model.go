package client

import (
	"context"
	"log"
	"time"

	pb "github.com/alsan/filebrowser/proto"
	utils "github.com/alsan/filebrowser/utils"

	"google.golang.org/grpc"
)

func createConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	utils.ExitIfError("Cannot create connetion to server, %v", err)
	defer conn.Close()

	return conn
}

func createRequestContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return ctx
}

func Login(address string, username string, password string) (*pb.LoginResponse, error) {
	log.Println("creating connection")
	conn := createConnection(address)
	client := pb.NewFileBrowserServiceClient(conn)
	ctx := createRequestContext()

	return client.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})
}

func ListFiles(address string, token string, path string) (*pb.ListFilesResponse, error) {
	// TODO: implementation
	return nil, nil
}

func UploadFile(address string, token string, path string, filename string, content []byte) error {
	// TODO: implementation
	return nil
}

func DownloadFile(address string, token string, path string) ([]byte, error) {
	// TODO: implementation
	return nil, nil
}
