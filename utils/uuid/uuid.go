package uuid

import uuid "github.com/satori/go.uuid"

func GetUUid() string {
	return uuid.NewV4().String()
}
