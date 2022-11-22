package util

import "github.com/google/uuid"

func CreateUUID() string {
	id := uuid.New()
	return id.String()
}
