package main

import (
	"log"

	"github.com/Yesterday17/pug-backend/controllers"
	"github.com/Yesterday17/pug-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.InitModels(models.InitSettings{
		Debug: true,
	})

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	// put frontend here
	r.StaticFile("/", "./public")

	// methods does not need authorization
	r.POST("/user/login", controllers.UserLogin)
	r.PUT("/user/register", controllers.UserRegister)

	// methods need authorization
	auth := Authorize("test_token")
	r.POST("/user/logout", auth, controllers.UserLogout)
	r.GET("/user", auth, nil)

	if err := r.Run(":14514"); err != nil {
		log.Fatal(err)
	}
}