package events

import "github.com/bekha-io/document-generator/domain/entities"

var _ Event = (*DocumentIssuedEvent)(nil)
type DocumentIssuedEvent struct {
	Document *entities.Document
}

func (e *DocumentIssuedEvent) Name() string { return "document.issued" }

var _ Event = (*DocumentSignedEvent)(nil)

type DocumentSignedEvent struct {
	Document *entities.Document
}
func (e *DocumentSignedEvent) Name() string { return "document.signed" }
