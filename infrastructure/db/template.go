package db

import (
	"context"
	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/repository"
	"github.com/bekha-io/document-generator/domain/types"
	"github.com/jmoiron/sqlx"
)

var _ repository.TemplateRepository = (*TemplateRepository)(nil)

type TemplateRepository struct {
	db *sqlx.DB
}

func NewTemplateRepository(db *sqlx.DB) *TemplateRepository {
	return &TemplateRepository{
		db: db,
	}
}

func (t *TemplateRepository) Get(ctx context.Context, id types.TemplateID) (*entities.Template, error) {
	e := &entities.Template{}
	row := t.db.QueryRowxContext(ctx, "SELECT * FROM templates WHERE id = $1", id.String())
	if row.Err() != nil {
		return nil, row.Err()
	}

	res := make(map[string]interface{})
	err := row.MapScan(res) 
	if err != nil {
		return nil, err
	}

	err = e.ParseMap(res)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (t *TemplateRepository) Save(ctx context.Context, template *entities.Template) error {
	_, err := t.db.NamedExecContext(ctx, `
	INSERT INTO templates (id, name, template_file) VALUES (:id, :name, :template_file)
	ON CONFLICT (id) DO UPDATE SET name = :name, template_file = :template_file;`, template.Map())
	return err
}
