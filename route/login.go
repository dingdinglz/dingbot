package route

import (
	"time"

	"github.com/dingdinglz/dingbot/database"
	"github.com/gofiber/fiber/v2"
)

func ApiLoginRoute(c *fiber.Ctx) error {
	if c.FormValue("username", "") == "" || c.FormValue("password", "") == "" {
		return JsonMessage(c, 1, "用户名或密码不能为空!")
	}
	ok := database.UserLogin(c.FormValue("username"), c.FormValue("password"))
	if !ok {
		return JsonMessage(c, 2, "用户名或密码错误!")
	}
	c.Cookie(&fiber.Cookie{Name: "token",
		Value:   jwtUserGenerate(c.FormValue("username")),
		Expires: time.Now().Add(24 * time.Hour)})
	return JsonMessage(c, 0, "ok")
}
