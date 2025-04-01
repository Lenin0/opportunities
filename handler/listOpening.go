package handler

import (
	"net/http"

	"github.com/Lenin0/opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Openning{}

	if err := db.Find(&openings).Error; err != nil {
		sendErro(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
