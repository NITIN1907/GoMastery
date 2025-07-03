package main

import (
	"context"
	proto "grpc/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.HelloServiceClient

func main() {
	//connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	//register client
	client = proto.NewHelloServiceClient(conn)

	//implement rest api
	r := gin.Default()
	r.GET("/sent-message-to-server/:message", clientConnectionServer)
	r.Run(":8000")

	//creating request
	// req := &proto.HelloRequest{Name: "hello from client"}

	//calling server
	// client.SayHello(context.TODO(), req)
}

func clientConnectionServer(c *gin.Context) {
	message := c.Param("message")
	req := &proto.HelloRequest{Name: message}
	client.SayHello(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully to server" + message,
	})
}
