package entities

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	ID        int       `gorm:"primary_key" json:"id"`
}
