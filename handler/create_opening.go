package handler

import (
	"net/http"

	"github.com/DaniloJNS/vagas-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context)  {
  request := CreateOpeningRequest{}

  if ctx.BindJSON(&request); ctx.IsAborted() {
    return
  }

  logger.Infof("request received: %+v", request)

  if err := request.Validate(); err != nil {
    logger.Errorf("validation error: %v", err.Error())
    sendError(ctx, http.StatusBadRequest, err.Error())
    return
  }


  opening := schemas.Opening {
    Role: request.Role,
    Company: request.Company,
    Location: request.Location,
    Remote: *request.Remote,
    Link: request.Link,
    Salary: request.Salary,
  }

  if err := db.Create(&opening).Error; err != nil {
    logger.Errorf("Error creating opening: %v", err.Error())
    sendError(ctx, http.StatusInternalServerError, "Error in persist operation")
  }

  sendSucess(ctx, "create-opening", opening)
}
