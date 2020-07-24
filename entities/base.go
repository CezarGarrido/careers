package entities

import (
	"time"
)

type Base struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}
