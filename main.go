package main

import (
	"github.com/bekha-io/document-generator/domain/services"
	"github.com/bekha-io/document-generator/infrastructure/db"
	"github.com/bekha-io/document-generator/presentation/api"
)

func main() {
	database := db.NewSqlxConnection("postgres", "postgres://planet9:planet9@localhost:5432/planet9?sslmode=disable")
	repoFacade := db.NewSqlxRepositoryFacade(database)

	tmplService := services.NewTemplateService(repoFacade)
	documentService := services.NewDocumentService(repoFacade)

	api := api.NewAPI(tmplService, documentService)
	api.E.Run(":8080")
}
