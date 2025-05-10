package entities

import "time"

type BaseEntity struct {
	CreatedAt time.Time  `gorm:"autoCreateTime" name:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" name:"updated_at"`
	DeletedAt *time.Time `gorm:"index" name:"deleted_at"`
	ID        int        `gorm:"primary_key" name:"id"`
}
