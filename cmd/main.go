package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"

	"test_hackers/internal/core/hackerservice"
	"test_hackers/internal/handlers/hackerhandler"
	"test_hackers/internal/repositories"
	pkg "test_hackers/pkg/storage"
)

func main() {
	conn, err := redis.Dial("tcp", os.Getenv("REDIS_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}
	pkg.FillStorage(conn)
	storage := repositories.NewStorage(conn)
	service := hackerservice.NewService(storage)
	handler := hackerhandler.NewHandler(service)

	app := fiber.New()
	app.Get("/json/hackers", handler.GetAllHackers)

	log.Fatal(app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
