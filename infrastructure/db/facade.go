package db

import (
	"github.com/bekha-io/document-generator/domain/repository"
	"github.com/jmoiron/sqlx"
)

func NewSqlxRepositoryFacade(db *sqlx.DB) *repository.RepositoryFacade {
	return &repository.RepositoryFacade{
		Templates: NewTemplateRepository(db),
		Documents: NewDocumentRepository(db),
	}
}
