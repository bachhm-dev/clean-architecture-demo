package main

import (
	"log"
	"net/http"

	"github.com/bachhm.dev/clean-architecture-service/controller/httpapi"
	openMeteoRepo "github.com/bachhm.dev/clean-architecture-service/infrastructure/open-meteo"
	redisRepo "github.com/bachhm.dev/clean-architecture-service/infrastructure/redis"
	weatherService "github.com/bachhm.dev/clean-architecture-service/service"
	"github.com/go-chi/chi/v5"

	// "github.com/go-redis/redis/v8"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Ensure Redis is running
	})

	// Set up the repository and service
	repo := redisRepo.NewRedisRepository(rdb)
	openMeteo := openMeteoRepo.New()
	service := weatherService.NewService(repo, openMeteo)

	controller := httpapi.NewAPIController(service)

	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		controller.SetUpRoute(r)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
