package model

import (
	"database/sql"
	"time"
)

type Link struct {
	ID        int64
	Info      LinkInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type LinkInfo struct {
	URL         string
	Title       string
	Description string
}