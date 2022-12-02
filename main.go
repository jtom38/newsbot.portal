package main

import (
	"context"
	"log"
	"net/http"

	//"github.com/jtom38/newsbot/portal/routes"
	"github.com/jtom38/newsbot/portal/services"
	"github.com/jtom38/newsbot/portal/web"
)

func main() {
	ctx := context.Background()

	c := services.NewConfigClient()
	apiAddress := c.GetConfig(services.Config_API_Address)

	//server := routes.NewServer(&ctx, apiAddress)
	server := web.NewServer(ctx, apiAddress)

	log.Print("Starting portal on http://localhost:8080")
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		panic(err)
	}
}
