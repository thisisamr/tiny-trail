package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thisisamr/tiny-trail/Middleware"
	"time"
)

type request struct {
	Url     string        `json:"url"`
	Expirey time.Duration `json:"expirey"`
}

type response struct {
	Url             string        `json:"url"`
	Expirey         time.Duration `json:"expirey"`
	XrateLimitReset time.Duration `json:"xrate_reset"`
	XrateRemaining  int           `json:"xrate_ramaining"`
}

func ClipUrl(c *fiber.Ctx) error {
	body := new(request)
	e := c.BodyParser(body)
	if e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	err := MiddleWare.ValidateURL(body.Url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}
	expirey_err := MiddleWare.ValidateExpiryTime(body.Expirey)
	if expirey_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error validatig expiry date"})
	}

	return c.Status(fiber.StatusOK).JSON(body)
}
