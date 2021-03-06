package main

import (
	"github.com/brendonmatos/golive"
	"github.com/brendonmatos/golive/examples/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {

	app := fiber.New()
	liveServer := golive.NewServer()

	loggerbsc := golive.NewLoggerBasic()
	loggerbsc.Level = golive.LogDebug
	liveServer.Log = loggerbsc.Log

	app.Get("/", liveServer.CreateHTMLHandler(components.NewTodo, golive.PageContent{
		Lang:  "us",
		Title: "Hello world",
	}))

	app.Get("/ws", websocket.New(liveServer.HandleWSRequest))

	_ = app.Listen(":3000")
}
