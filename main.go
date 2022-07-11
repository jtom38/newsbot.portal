package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jtom38/newsbot/portal/routes"
)

func main() {
	ctx := context.Background()
	server := routes.NewServer(&ctx, "http://localhost:8081")

	log.Print("Starting portal on :8080")
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		panic(err)
	}
}