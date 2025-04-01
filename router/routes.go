package router

import (
	docs "github.com/Lenin0/opportunities/docs"
	"github.com/Lenin0/opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/opening/list", handler.ListOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	// Initialize Swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
