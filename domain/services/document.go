package services

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"log/slog"

	"github.com/bekha-io/document-generator/domain/entities"
	"github.com/bekha-io/document-generator/domain/errs"
	"github.com/bekha-io/document-generator/domain/repository"
)

type DocumentService struct {
	Log        *slog.Logger
	Repository *repository.RepositoryFacade
}

func NewDocumentService(repo *repository.RepositoryFacade) *DocumentService {
	logger := slog.With("service", "DocumentService")
	return &DocumentService{
		Repository: repo,
		Log:        logger,
	}
}

// GenerateDocument creates a new document based on the specified template
func (s *DocumentService) GenerateDocument(ctx context.Context, tmpl *entities.Template, templateData map[string]interface{}) (*entities.Document, error) {
	doc := entities.NewDocument(tmpl, templateData)

	// Parsing html template
	htmlTemplate, err := template.New(tmpl.Name).Parse(string(tmpl.TemplateFile.Data))
	if err != nil {
		return nil, errors.Join(errs.ErrTemplateNotValid, err)
	}

	// Rendering html template
	buf := bytes.NewBuffer([]byte(""))
	err = htmlTemplate.ExecuteTemplate(buf, tmpl.Name, doc.TemplateData)
	if err != nil {
		return nil, errors.Join(errs.ErrDocumentTemplateDataInvalid, err)
	}
	doc.Content = buf.Bytes()

	err = s.Repository.Documents.Save(ctx, doc)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
