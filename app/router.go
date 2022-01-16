package app

import (
	"github.com/gin-contrib/cors"
	"auth-jwt/controller/users"
	"time"
)

func mapUrls() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	  }))
	router.POST("/api/register", users.Register)
	router.POST("/api/login", users.Login)
	router.GET("api/user", users.Get)
	router.GET("api/logout", users.Logout)
}