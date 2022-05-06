package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Get("authorization")
	token := strings.Replace(cookie, "Bearer ", "", 1)
	if _, err := util.ParseJwt(token); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}
