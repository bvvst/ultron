package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Post("/newsms", func(c *fiber.Ctx) error {
		msg := new(TwilioIncomingMessageBody)

		if err := c.BodyParser(msg); err != nil {
			return err
		}

		log.Println(msg.Body)
		gptResp := QueryChatGPT(msg.Body)

		resp := new(Response)
		resp.Message.Body = gptResp

		return c.XML(resp)
	})

	log.Fatal(app.Listen(":3000"))
}
