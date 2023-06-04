package handler

import (
	"net/http"

	"github.com/DaniloJNS/vagas-api/schemas"
	"github.com/gin-gonic/gin"
)

func IndexOpeningHandler(ctx *gin.Context) {
  openings := []schemas.Opening{}

  if err := db.Find(&openings).Error; err!= nil {
    sendError(ctx, http.StatusInternalServerError, "err list openings")
  }
  
  sendSucess(ctx, "list-openings", openings)
}
