package types

import (
	"github.com/google/uuid"
)

type TemplateID uuid.UUID

func (t TemplateID) String() string {
	return uuid.UUID(t).String()
}
