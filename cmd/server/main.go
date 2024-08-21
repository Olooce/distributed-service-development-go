package main

import (
	"log"

	"github.com/Olooce/distributed-service-development-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
