package api

import (
	"github.com/bekha-io/document-generator/domain/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *API) GenerateDocument(c *gin.Context) {
	ctx := c.Request.Context()
	templateIdRaw := c.Param("template_id")
	templateId, err := uuid.Parse(templateIdRaw)
	if err != nil {
		c.String(400, "Invalid template id")
		return
	}

	var templateInput map[string]interface{}
	c.BindJSON(&templateInput)

	tmpl, err := a.TemplateService.GetTemplate(ctx, types.TemplateID(templateId))
	if err != nil {
		c.String(400, err.Error())
		return
	}

	doc, err := a.DocumentService.GenerateDocument(c.Request.Context(), tmpl, templateInput)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Data(200, "text/html", doc.Content)
}
