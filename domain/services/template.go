package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/errs"
	"github.com/bekha-io/document-generator/domain/repository"
	"github.com/bekha-io/document-generator/domain/types"
)

type TemplateService struct {
	Log        *slog.Logger
	Repository *repository.RepositoryFacade
}

func NewTemplateService(repo *repository.RepositoryFacade) *TemplateService {
	logger := slog.With("service", "TemplateService")
	return &TemplateService{
		Repository: repo,
		Log:        logger,
	}
}

func (s *TemplateService) GetTemplate(ctx context.Context, id types.TemplateID) (*entities.Template, error) {
	tmpl, err := s.Repository.Templates.Get(ctx, id)
	if err != nil {
		return nil, errors.Join(err, errs.ErrTemplateNotFound)
	}
	return tmpl, nil
}
