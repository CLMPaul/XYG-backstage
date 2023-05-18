package db

import (
	"time"
)

type VersionRecord struct {
	ModuleName string `gorm:"primaryKey"`
	Version    string
	UpdatedAt  time.Time
}
