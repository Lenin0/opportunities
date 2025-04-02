package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlit:: %v", err)
	}

	return nil
}

func GetSQlite() *gorm.DB {
	return db
}

func GetLogger(prefix string, level LogLevel, outputs ...string) *Logger {
	if logger == nil {
		logger = NewLogger(prefix, level, outputs...)
	}
	return logger
}
