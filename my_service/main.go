package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"log"
	"runtime"
)

var favorites = map[string]string{
	"color": "blue",
	"food":  "pizza",
	"drink": "coke",
	"movie": "matrix",
}

func main() {
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS server", nc.ConnectedUrl())

	service, err := micro.AddService(nc, micro.Config{
		Name:        "Isaac",
		Version:     "0.0.1",
		Description: "just a service for bookmarks",
	})

	service.AddEndpoint("list_favorites",
		micro.HandlerFunc(listFavorites),
		micro.WithEndpointSubject("isaac.favorite.list"),
		micro.WithEndpointMetadata(map[string]string{
			"description": "list all favorites",
		}),
	)

	runtime.Goexit()
}

func listFavorites(request micro.Request) {
	request.RespondJSON(favorites)
	
}
