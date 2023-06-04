package router

import (
	"github.com/DaniloJNS/vagas-api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouters(router *gin.Engine) {
  handler.InitializeHandler()
  v1 := router.Group("api/v1")
  {
    v1.GET("/openings", handler.IndexOpeningHandler)

    v1.GET("/opening", handler.ShowOpeningHandler)

    v1.POST("/opening", handler.CreateOpeningHandler)

    v1.DELETE("/opening", handler.DeleteOpeningHandler)
  }
}
