package route

import "github.com/gofiber/fiber/v2"

func LoginMiddleware(c *fiber.Ctx) error {
	if c.Path() == "/login" || c.Path() == "/api/login" {
		return c.Next()
	}
	t := c.Cookies("token", "")
	if t == "" {
		return c.Redirect("/login")
	}
	user, ok := jwtUserParse(t)
	if !ok {
		return c.Redirect("/login")
	}
	c.Locals("user", user)
	return c.Next()
}
