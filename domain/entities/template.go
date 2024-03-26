package entities

import (
	"errors"

	"github.com/bekha-io/document-generator/domain/types"
	"github.com/google/uuid"
)

var _ Entity = (*Template)(nil)

type Template struct {
	ID   types.TemplateID
	Name string

	TemplateFile types.TemplateFile
}

func (t *Template) Map() map[string]interface{} {
	return map[string]interface{}{
		"id":            t.ID,
		"name":          t.Name,
		"template_file": t.TemplateFile.Data,
	}
}

// ParseMap implements Entity.
func (t *Template) ParseMap(i map[string]interface{}) error {
	id, ok := i["id"].(string)
	if !ok {
		return errors.New("empty id")
	}
	t.ID = types.TemplateID(uuid.MustParse(id))

	templateFileRaw, ok := i["template_file"].(string)
	if !ok {
		return errors.New("empty template_file")
	}
	t.TemplateFile = types.TemplateFile{Data: []byte(templateFileRaw)}

	name, ok := i["name"].(string)
	if !ok {
		return errors.New("empty name")
	}
	t.Name = name

	return nil
}
