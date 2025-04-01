package handler

import (
	"net/http"

	"github.com/Lenin0/opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	var request CreateOpeningRequest

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("invalid json: %v", err.Error())
		sendErro(ctx, http.StatusBadRequest, "invalid JSON")
		return
	}

	if validationErrors := request.Validate(); len(validationErrors) > 0 {
		logger.Errorf("validation errors: %v", validationErrors)
		sendErro(ctx, http.StatusBadRequest, validationErrors)
		return
	}

	opening := schemas.Openning{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendErro(ctx, http.StatusInternalServerError, "error crreating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
