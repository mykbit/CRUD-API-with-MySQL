package models

import (
	"database/sql"
	"time"
)

type Vinyl struct {
	ID          int       `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	ReleaseDate time.Time `json:"ReleaseDate"`
}

type VinylModel struct {
	DB *sql.DB
}
