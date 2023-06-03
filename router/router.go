package router

import "github.com/gin-gonic/gin"

func Initializer() {
  r := gin.Default()
  // define an routes to GET /ping
  r.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "pong",
    })
  })

  // Initialize WEB SERVER
  r.Run(":4000")
}
