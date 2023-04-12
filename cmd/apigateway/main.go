package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/apigateway/server"
	"github.com/dietzy1/chatapp/services/apigateway/websocket"
)

func main() {

	config.ReadEnvfile()

	log.Println("Starting Websocket Server")
	go websocket.Start()

	log.Println("Starting API Gateway")
	server.Start()

}

//Currently I have a GRPC gateway server which calls a remote method

//The GRPC gateway server is running on port 8090 and the GRPC server is running on port 8080
//it should perform a call to the authentication service and return the token
//token should then be stored as a cookie as a header in the response
//the client should then be able to access the other services with the token as a header
