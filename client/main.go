package main

import (
	"fmt"
	"gsm/client/swagger"
	"gsm/client/task"
	"gsm/config"
	"log"

	_ "gsm/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title           GSA API
// @version         1.0
// @description     Golang Swagger Architectures endpoints

// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Println("client starting...")

	engine := gin.Default()

	conn, err := newClient()
	if err != nil {
		log.Println(err)
		return
	}

	task.MakeConnection(conn)

	group := engine.Group("api")
	swagger.AddRoutes(group)
	task.AddRoutes(group)

	engine.Run(config.CLIENT_HTTP_ADDRESS)
}

func newClient() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(config.SERVER_GRPC_ADDRESS,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
