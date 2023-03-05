package main

import (
	"fmt"
	"log"

	"github.com/vbrenister/go-rest-example/api/server"
)

const (
	port = ":3000"
)

func main() {
	fmt.Println("Server is running at localhost:", port)
	log.Fatal(server.Run(port))
}
