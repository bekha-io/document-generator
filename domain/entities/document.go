package entities

import (
	"encoding/json"
	"errors"

	"github.com/bekha-io/document-generator/domain/errs"
	"github.com/bekha-io/document-generator/domain/types"
	"github.com/google/uuid"
)

var _ Entity = (*Document)(nil)

type Document struct {
	ID         types.DocumentID // ID документа
	TemplateID types.TemplateID // Шаблон документа

	TemplateData map[string]interface{} // Входные данные для рендеринга шаблона
	Content      []byte                 // Результат рендеринга

	SignInformation []types.DocumentSignInformation // Информация о лицах, подписавших документ
}

func NewDocument(template *Template, templateData map[string]interface{}) *Document {
	return &Document{
		ID:           types.DocumentID(uuid.New()),
		TemplateID:   template.ID,
		TemplateData: templateData,
	}
}

// Map implements Entity.
func (d *Document) Map() map[string]interface{} {

	templateData, _ := json.Marshal(d.TemplateData)

	return map[string]interface{}{
		"id":               d.ID.String(),
		"template_id":      d.TemplateID.String(),
		"template_data":    templateData,
		"content":          d.Content,
		"sign_information": nil, // TODO: Add field
	}
}

// ParseMap implements Entity.
func (d *Document) ParseMap(input map[string]interface{}) error {
	id, ok := input["id"].(string)
	if !ok {
		return errors.New("empty id")
	}
	d.ID = types.DocumentID(uuid.MustParse(id))

	templateId, ok := input["template_id"].(string)
	if !ok {
		return errors.New("empty template_id")
	}
	d.TemplateID = types.TemplateID(uuid.MustParse(templateId))

	templateData, ok := input["template_data"].([]byte)
	if !ok {
		return errors.New("empty template_data")
	}
	err := json.Unmarshal(templateData, &d.TemplateData)
	if err != nil {
		return err
	}

	content, ok := input["content"].([]byte)
	if !ok {
		return errors.New("empty content")
	}
	d.Content = content

	// TODO: Add sign information support

	return nil
}

func (d *Document) Sign(sign types.DocumentSignInformation) error {
	for _, info := range d.SignInformation {
		if info.SignatureHash() == sign.SignatureHash() {
			return errs.ErrDocumentAlreadySigned
		}
	}
	d.SignInformation = append(d.SignInformation, sign)
	return nil
}

func (d *Document) SetInputValue(key string, value interface{}) {
	d.TemplateData[key] = value
}

func (d *Document) RemoveInputValue(key string) error {
	if _, ok := d.TemplateData[key]; !ok {
		return errors.New("key not found")
	}
	delete(d.TemplateData, key)
	return nil
}
