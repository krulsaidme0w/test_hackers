package main

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"os"
	pkg "test_hackers/pkg/storage"

	"github.com/gofiber/fiber/v2"

	"test_hackers/internal/core/hackerservice"
	"test_hackers/internal/handlers/hackerhandler"
	"test_hackers/internal/repositories"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})
	pkg.FillStorage(rdb, &ctx)
	storage := repositories.NewStorage(rdb, &ctx)
	service := hackerservice.NewService(storage)
	handler := hackerhandler.NewHandler(service)

	app := fiber.New()
	app.Get("/json/hackers", handler.GetAllHackers)

	log.Fatal(app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
