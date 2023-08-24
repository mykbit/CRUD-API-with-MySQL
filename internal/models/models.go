package models

import (
	"database/sql"
	"time"
)

type Vinyl struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
}

type VinylModel struct {
	DB *sql.DB
}
