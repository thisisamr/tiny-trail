package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/thisisamr/tiny-trail/db"
)

func ResolveUrl(c *fiber.Ctx) error {
	client := db.Redis_Client
	value, err := client.Get(db.Ctx, c.Params("url")).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Err": "not found"})
	}
	if err != redis.Nil && err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		return c.Status(fiber.StatusOK).JSON(value)
	}
}
