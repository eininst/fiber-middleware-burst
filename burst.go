package burst

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"time"
)

type Config struct {
	Limiter        *rate.Limiter
	Timeout        time.Duration
	TimeoutHandler fiber.Handler
}

func New(cfg Config) fiber.Handler {
	if cfg.Timeout == 0 {
		panic("Timeout cannot be empty")
	}
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), cfg.Timeout)
		defer cancel()
		err := cfg.Limiter.Wait(ctx)
		if err != nil {
			if cfg.TimeoutHandler == nil {
				return c.SendStatus(fiber.StatusTooManyRequests)
			}
			return cfg.TimeoutHandler(c)
		}
		return c.Next()
	}
}
