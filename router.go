package main

import (
	"github.com/gofiber/fiber"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/kit/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

type app struct {
	service Service
}

func (a *app) handleGetLinks(c *fiber.Ctx) {
	links, err := a.service.GetLinks(c.Context())
	if err != nil {
		_ = c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}
	_ = c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    links,
	})
}

func LinkerRouter() *fiber.App {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = NewPostgres(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	s := app{NewService(r)}
	app := fiber.New()
	app.Get("/", s.handleGetLinks)
	return app
}
