package utils

import (
	"github.com/google/uuid"
	"strings"
)

// MakeUid 生成唯一用户id
func MakeUid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// MakeIdentity 生成唯一标识符
func MakeIdentity() string {
	return uuid.New().String()
}
