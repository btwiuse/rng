package rng

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	randid := uuid.New()
	return randid.String()
}

func NewPetUUID(s string) string {
	petid := uuid.NewSHA1(uuid.Nil, []byte(s))
	return petid.String()
}
