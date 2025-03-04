package main

import (
	"fmt"

	"github.com/clevim/ConnectMais/router"
)

func main() {
	r := router.SetupRouter()

	router.ListRoutes(r)

	fmt.Println("Server is running at http://localhost:8080")
	r.Run()
}
