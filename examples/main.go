package main

import (
	burst "github.com/eininst/fiber-middleware-burst"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	app := fiber.New()

	app.Use(burst.New(burst.Config{
		//每秒产生100个令牌, 最多存储200个令牌。
		Limiter: rate.NewLimiter(rate.Every(time.Second*100), 200),
		//拿不到令牌时等待3秒，仍拿不到则返回 "too many requests"
		Timeout: time.Second * 3,
	}))

	_ = app.Listen(":8080")
}
