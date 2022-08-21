package uuid

import (
	"github.com/google/uuid"
)

func New() (string, error) {
	u, err := generate()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func generate() (*uuid.UUID, error) {
	u := uuid.New()
	return &u, nil
}
