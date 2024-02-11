package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GreetController struct {
	db *gorm.DB
}

func NewGreetController(db *gorm.DB) GreetController {
	return GreetController{db: db}
}

func (c GreetController) HelloWorld(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}
