package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// API v1

	/*

		// GET /api/v1/agent/register/:agentid
		Get the required information to setup the agent, include various information in body, like interface info, etc.
		After sending the data in the body from the agent, the backend (this) receives it, and responds with a token.
		The agent receives the token and saves it, after requested, once a token is generated, registering cannot be done again
		unless it is done from the panel, it would then require the token to be re generated.

		** Agents can also be disabled.

		Get information using agent id (pooling time, etc.) <- From Agent
		Agent Receives information, then generates

	*/

	app.Get("/api/v1/register/:agentid", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// POST /api/v1/agent/register/:agentid
	/*
			The agent would then post after getting the token response, confirming that it has been setup,
			it also includes their public ip, etc. The information becomes permanent

		Only allow if agent ID exists. Throw HTTP codes!
	*/

	// POST /api/v1/agent/put/:agentid
	/*
		The token is included in the request, and it includes the public ip information, etc. This can be used
		to track IP changes and such.

		This is used to send information from data it's collected, or from tools like nmap, etc.
	*/

	// GET /api/v1/agent/queue/:agentid
	/*
		This grabs the queue of tasks to be completed such as running an iperf speed test, or running an nmap test,
		dns lookup, or other general network tasks, these tasks will have an id for each task, that can be correlated
		when posting

	*/

	// GET /api/v1/agent/update/:agentid
	/*
		Update configuration such as the peers for running speed tests and such, it also grabs the destination information
		that will be used for trace routes, and basic network path monitoring, the backend will be processing it
	*/

	// GET /api/v1/
	/*

	 */

	// GET /api/v1/agent/update/:token
	/*
		This is for uploading the data every 500ms, or otherwise specified when
	*/
	app.Get("/api/v1/update/:token", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	log.Fatal(app.Listen(":3000"))
}
