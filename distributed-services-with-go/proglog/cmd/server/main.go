package main

import (
	"log"
	"golang/go/distributed-services-with-go/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
