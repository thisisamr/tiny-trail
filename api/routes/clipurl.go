package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/thisisamr/tiny-trail/Middleware"
	"github.com/thisisamr/tiny-trail/db"
)

type request struct {
	Url        string        `json:"url"`
	Expirey    time.Duration `json:"expirey"`
	Custom_Url string        `json:"custom_url"`
}

type response struct {
	Url             string        `json:"url"`
	Expirey         time.Duration `json:"expirey"`
	XrateLimitReset time.Duration `json:"xrate_reset"`
	XrateRemaining  int           `json:"xrate_ramaining"`
}

func ClipUrl(c *fiber.Ctx) error {
	body := new(request)
	if e := c.BodyParser(body); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	if err := MiddleWare.ValidateURL(body.Url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}
	if expirey_err := MiddleWare.ValidateExpiryTime(body.Expirey); expirey_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error validatig expiry date"})
	}
	// get the ip user data
	result, err := db.Redis_Client.Get(db.Ctx, c.IP()).Result()
	remaining, _ := strconv.Atoi(result)
	v, _ := db.Redis_Client.TTL(db.Ctx, c.IP()).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "ooops"})
	}
	// get the expirytime
	time_toexpire := time.Duration(body.Expirey * time.Minute)
	if body.Custom_Url == "" {
		id := uuid.New().String()[:6]
		err := db.Redis_Client.Set(db.Ctx, id, body.Url, time_toexpire).Err()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		response := response{Url: "http://" + os.Getenv("DOMAIN") + "/" + id, Expirey: time.Duration(time_toexpire.Minutes()), XrateLimitReset: time.Duration(time.Duration(v).Minutes()), XrateRemaining: remaining}
		return c.Status(fiber.StatusCreated).JSON(response)

	} else {
		err := db.Redis_Client.Get(db.Ctx, body.Custom_Url).Err()
		if err == redis.Nil {
			// thats cool no one is using it
			err := db.Redis_Client.Set(db.Ctx, body.Custom_Url, body.Url, time_toexpire).Err()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(err)
			} else {
				response := response{Url: "http://" + os.Getenv("DOMAIN") + "/" + body.Custom_Url, Expirey: time.Duration(time_toexpire.Minutes()), XrateLimitReset: v, XrateRemaining: remaining}
				return c.Status(fiber.StatusCreated).JSON(response)
			}

		} else {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "custom url already exists"})
		}
	}

}
