# Burst

[![Build Status](https://travis-ci.org/ivpusic/grpool.svg?branch=master)](https://github.com/infinitasx/easi-go-aws)

> 令牌桶限流器
>
> 基于 Golang 官方扩展包 `golang.org/x/time/rate`

## ⚙ Installation

```text
go get -u github.com/eininst/fiber-middleware-burst
```

## ⚡ Quickstart

```go
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
        //每秒产生200个令牌, 最多存储500个令牌。
        Limiter: rate.NewLimiter(200, 500),
        //拿不到令牌时等待3秒，仍拿不到则返回 "too many requests"
        Timeout: time.Second * 3,
    }))

    _ = app.Listen(":8080")
}
```

> See [examples](/examples)

## License

*MIT*