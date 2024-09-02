package MiddleWare

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/thisisamr/tiny-trail/db"
)

func Rate_Limiter(c *fiber.Ctx) error {
	// implement the rate limiting
	ip := c.IP()
	r_client := db.Redis_Client

	value, err := r_client.Get(db.Ctx, ip).Result()
	if err == redis.Nil {
		r_client.Set(db.Ctx, ip, 5, (1 * time.Minute)).Err()
	} else {
		v, _ := strconv.Atoi(value)
		if v < 1 {
			resettime, _ := r_client.TTL(db.Ctx, ip).Result()
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("limit exceeded try again in %.0f S", resettime.Seconds())})
		} else {
			r_client.Decr(db.Ctx, ip)
		}
	}
	return c.Next()
}
