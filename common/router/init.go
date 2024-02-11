package router

import (
	"kakeru-pro-web/common/config"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(corsConfig())

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	r.Run(":" + strconv.Itoa(config.Config.PORT))
}

func corsConfig() gin.HandlerFunc {
	corsConfig := cors.Config{}
	corsConfig.AllowOrigins = []string{config.Config.FRONT_URL}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	return cors.New(corsConfig)
}
