package entities

import (
	"time"
)

type Book struct {
	Id          string
	Title       string
	Author      string
	Published   string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
