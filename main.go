package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		currentYear := time.Now().Year()
		title := "Is it 2024 yet?"
		color := "red"
		message := "it isn't."
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		if currentYear == 2024 {
			title = "Is it 2024 yet?"
			color = "green"
			message = "It is!"
			currentTime = time.Now().Format("2006-01-02 15:04:05")
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
					h2 {
						color: rgba(255, 255, 255, 0.2);
						font-size: 10px;
						margin: 0;
					}
					a {
						color: #1a1a1a;
					  }
				</style>
			</head>
			<body>
				<h1>%s</h1>
				<h2>checked on %s (EST)</h2>
				<h2>source: <a href="https://github.com/ndzn/isit2024yet">github</a> </h2>
			</body>
			<script async src="https://vim.screenshot.download/_.js" data-ackee-server="https://vim.screenshot.download" data-ackee-domain-id="1ce83615-e193-42eb-96cf-ade09f75cae3" data-ackee-opts='{ "detailed": true }'></script>
			</html>
		`, title, color, message, currentTime)
		c.Set("Content-Type", "text/html")
		return c.SendString(html)
	})
	app.Listen(":8080")
}
