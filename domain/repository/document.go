package repository

import (
	"context"

	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/types"
)

type DocumentRepository interface {
	Get(ctx context.Context, id types.DocumentID) (*entities.Document, error)
	Save(ctx context.Context, document *entities.Document) error
}