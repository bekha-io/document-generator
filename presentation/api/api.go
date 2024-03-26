package api

import (
	"github.com/bekha-io/document-generator/domain/services"
	"github.com/gin-gonic/gin"
)

type API struct {
	TemplateService *services.TemplateService
	DocumentService *services.DocumentService

	E *gin.Engine
}

func NewAPI(ts *services.TemplateService, ds *services.DocumentService) *API {
	a := &API{
		TemplateService: ts,
		DocumentService: ds,
	}

	e := gin.New()
	e.POST("/documents/:template_id", a.GenerateDocument)

	a.E = e
	return a
}
