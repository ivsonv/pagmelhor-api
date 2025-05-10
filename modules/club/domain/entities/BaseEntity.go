package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt time.Time  `gorm:"autoCreateTime" name:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime;null" name:"updated_at"`
	DeletedAt *time.Time `gorm:"index;null" name:"deleted_at"`
	ID        int        `gorm:"primary_key" name:"id"`
}

func (p *PartnerEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	p.UpdatedAt = &now
	return
}

func (p *PartnerEntity) BeforeDelete(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	p.DeletedAt = &now
	return
}
