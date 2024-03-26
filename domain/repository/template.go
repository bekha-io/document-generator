package repository

import (
	"context"

	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/types"
)

type TemplateRepository interface {
	Get(ctx context.Context, id types.TemplateID) (*entities.Template, error)
	Save(ctx context.Context, template *entities.Template) error
}