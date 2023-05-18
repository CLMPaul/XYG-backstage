package internal

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64     `json:"id"  gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ModelLogically struct {
	ID        int64          `json:"id"  gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

type ViewModel struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}
