package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	greetv1 "github.com/nakaaaa/connect/go/gen/greet/v1"
	"github.com/nakaaaa/connect/go/gen/greet/v1/greetv1connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
