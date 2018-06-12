package handlers

import (
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/sirupsen/logrus"
)

func GetRouter(app *application.Application) *gin.Engine {

	h := Handler{Application: app}

	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false),
	)

	resourceRouter := engine.Group("/retranslator/v1/resources")

	resourceRouter.GET("/", h.ListResources)
	resourceRouter.POST("/", h.UpdateOrCreate)
	resourceRouter.GET("/:name", h.GetResource)
	resourceRouter.DELETE("/:name", h.Delete)
	return engine
}
