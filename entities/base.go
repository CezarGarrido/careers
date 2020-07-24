package entities

import (
	"time"
)

type Base struct {
	ID        int64     `json:"id"`
	UUID      int64     `json:"uuid"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}
