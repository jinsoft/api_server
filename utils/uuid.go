package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GenUUID() string {
	return uuid.NewString()
}

func GenUUID32() string {
	uuidStr := uuid.NewString()
	return strings.ReplaceAll(uuidStr, "-", "")
}
