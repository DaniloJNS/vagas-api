package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK,  gin.H{
    "message": "GET Opening",
  })
}

func IndexOpeningHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK,  gin.H{
    "message": "GET Opening",
  })
}

func ShowOpeningHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK,  gin.H{
    "message": "GET Opening",
  })
}

func DeleteOpeningHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK,  gin.H{
    "message": "GET Opening",
  })
}
