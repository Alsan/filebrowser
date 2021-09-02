package main

import (
	"log"

	pb "github.com/alsan/filebrowser/proto"
	utils "github.com/alsan/filebrowser/utils"
)

func Main() {
	resp, err := Login("127.0.0.1:8080", "alsan", "KyHS4s77t1")
	utils.ExitIfError("Unable connect to server, %v", err)

	if resp.GetStatus() == pb.ResponseStatus_Ok {
		log.Printf("login successful, token: %s", resp.GetToken())
	} else {
		log.Fatalf("Login failed, reason: %s", resp.GetMessage())
	}
}
