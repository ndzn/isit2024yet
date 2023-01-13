package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		currentYear := time.Now().Year()
		title := "Is it 2024 yet?"
		color := "red"
		message := "it isn't."

		if currentYear == 2024 {
			title = "Is it 2024 yet?"
			color = "green"
			message = "It is!"
		}

		html := fmt.Sprintf(`
			<html>
			<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>%s</title>
				<style>
					body {
						background-color: black;
						color: white;
						font-family: sans-serif;
						text-align: center;
					}
					h1 {
						color: %s;
						font-size: 50px;
						margin: 50px 0;
					}
				</style>
			</head>
			<body>
				<h1>%s</h1>
			</body>
			<script async src="https://vim.screenshot.download/_.js" data-ackee-server="https://vim.screenshot.download" data-ackee-domain-id="1ce83615-e193-42eb-96cf-ade09f75cae3" data-ackee-opts='{ "detailed": true }'></script>
			</html>
		`, title, color, message)
		c.Set("Content-Type", "text/html")
		c.Send(html)
	})
	app.Listen(":8080")
}
