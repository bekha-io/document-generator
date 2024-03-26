package db

import (
	"context"

	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/repository"
	"github.com/bekha-io/document-generator/domain/types"
	"github.com/jmoiron/sqlx"
)

var _ repository.DocumentRepository = (*DocumentRepository)(nil)

type DocumentRepository struct {
	db *sqlx.DB
}

func NewDocumentRepository(db *sqlx.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

// Get implements repository.DocumentRepository.
func (d *DocumentRepository) Get(ctx context.Context, id types.DocumentID) (*entities.Document, error) {
	panic("unimplemented")
}

// Save implements repository.DocumentRepository.
func (d *DocumentRepository) Save(ctx context.Context, document *entities.Document) error {
	_, err := d.db.NamedExecContext(ctx, `
	INSERT INTO documents (id, template_id, template_data) VALUES (:id, :template_id, :template_data)
	ON CONFLICT (id) DO UPDATE SET template_id = :template_id, template_data = :template_data ;`, document.Map())
	return err
}
