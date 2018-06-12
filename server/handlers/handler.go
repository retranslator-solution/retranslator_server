package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/retranslator-solution/retranslator_server/models"
	"github.com/retranslator-solution/retranslator_server/storage"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	*application.Application
}

func (h *Handler) GetResource(c *gin.Context) {
	resources, err := h.Storage.Get(c.Param("name"))

	switch err {
	case storage.NotFound:
		c.AbortWithStatus(http.StatusNotFound)
	case nil:
		c.JSON(http.StatusOK, resources)
	default:
		log.Errorln("can not GetResource", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *Handler) ListResources(c *gin.Context) {
	resources, err := h.Storage.GetResourceNames()

	switch err {
	case nil:
		c.JSON(http.StatusOK, resources)
	default:
		log.Errorln("can not ListResources", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateOrCreate(c *gin.Context) {
	var resource models.Resource

	if err := c.BindJSON(&resource); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err := h.Storage.UpdateOrCreate(&resource)

	switch err {
	case nil:
		c.JSON(http.StatusOK, &resource)
	default:
		log.Errorln("can not UpdateOrCreate", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}


func (h *Handler) Delete(c *gin.Context) {
	err := h.Storage.Delete(c.Param("name"))

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	default:
		log.Errorln("can not Delete", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
