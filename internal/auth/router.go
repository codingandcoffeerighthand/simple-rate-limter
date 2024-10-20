package auth

import (
	"learn-redis/internal/auth/adapter/http"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(r fiber.Router, db *gorm.DB, rdb *redis.Client) {
	api := http.NewApi(db, rdb)
	r.Post("/register", api.Register)
	r.Post("/login", api.Login)
	r.Get("/ping/top-user", api.TopUser)
	r.Get("/ping/count", api.CountUser)
	r.Get("/ping", api.Ping)
}
