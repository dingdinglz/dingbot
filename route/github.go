package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GithubWebhookRoute(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Body())
	return JsonMessage(ctx, 0, "ok")
}
