package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/sirupsen/logrus"
)

func GetRouter(app *application.Application) *gin.Engine {

	h := Handler{Application: app}

	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false),
		cors.Default(),
	)

	indexBox := packr.NewBox("../../static/dist")
	engine.GET("/", func(c *gin.Context) {
		c.Writer.Write(indexBox.Bytes("index.html"))
	})

	staticBox := packr.NewBox("../../static/dist/static")
	engine.StaticFS("/static", staticBox)

	resourceRouter := engine.Group("/retranslator/v1/resources")

	resourceRouter.GET("/", h.ListResources)
	resourceRouter.POST("/", h.UpdateOrCreate)
	resourceRouter.GET("/:name", h.GetResource)
	resourceRouter.DELETE("/:name", h.Delete)
	return engine
}
