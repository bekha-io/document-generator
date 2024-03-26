package types

import (
	"crypto/sha256"
	"time"

	"github.com/google/uuid"
)

type DocumentID uuid.UUID

func (d DocumentID) String() string {
	return uuid.UUID(d).String()
}

type DocumentSignInformation struct {
	PersonSinged string
	SignedAt     time.Time
}

func NewDocumentSignInformation(personSigned string) DocumentSignInformation {
	return DocumentSignInformation{
		PersonSinged: personSigned,
		SignedAt:     time.Now().UTC(),
	}
}

func (d DocumentSignInformation) SignatureHash() string {
	s := sha256.New()
	res := s.Sum([]byte(d.PersonSinged))
	return string(res)
}
