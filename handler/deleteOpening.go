package handler

import (
	"fmt"
	"net/http"

	"github.com/Lenin0/opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendErro(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Openning{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendErro(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendErro(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)

}
