package main

import (
	"log"

	"github.com/Yesterday17/pug-backend/auth"
	"github.com/Yesterday17/pug-backend/config"
	"github.com/Yesterday17/pug-backend/controllers"
	"github.com/Yesterday17/pug-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()
	db := models.InitModels(&cfg.ModelSettings)

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Set("config", cfg)
		ctx.Next()
	})

	// frontend
	r.StaticFile("/", "./public")

	// Session
	r.POST("/session", controllers.SessionCreate, controllers.SessionUpdate)
	r.PUT("/session", controllers.UserRegister, controllers.SessionUpdate)
	r.GET("/session/key", controllers.SessionGetKey)
	r.DELETE("/session", auth.Authorize, controllers.SessionRevoke)

	// User
	r.GET("/user", auth.Authorize, nil, controllers.SessionUpdate)

	if err := r.Run(cfg.Listen); err != nil {
		log.Fatal(err)
	}
}
