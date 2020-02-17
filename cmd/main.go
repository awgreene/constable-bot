package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func handle(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Hey!")
}

func main() {
	bot := slacker.NewClient("SomeToken")
	definition := &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
