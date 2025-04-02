package handler

import (
	"github.com/Lenin0/opportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler", config.InfoLevel)
	db = config.GetSQlite()
}
