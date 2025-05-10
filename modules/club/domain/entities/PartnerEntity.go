package entities

import (
	"app/modules/club/domain"
)

func (p PartnerEntity) TableName() string {
	return domain.SchemaClubName + ".partners"
}

type PartnerEntity struct {
	BaseEntity
	Name            string `gorm:"not null" name:"name"`
	CpfCnpj         string `gorm:"not null;unique" name:"cpf_cnpj"`
	Email           string `gorm:"not null;unique" name:"email"`
	Slug            string `gorm:"not null;unique" name:"slug"`
	Status          int16  `gorm:"not null;default:1" name:"status"`
	Phone           string `gorm:"null" name:"phone"`
	Password        string `gorm:"null" name:"password"`
	Image           string `gorm:"null" name:"image"`
	Address         string `gorm:"null" name:"address"`
	City            string `gorm:"null" name:"city"`
	State           string `gorm:"null" name:"state"`
	ZipCode         string `gorm:"null" name:"zip_code"`
	Description     string `gorm:"type:text;null" name:"description"`
	MetaTitle       string `gorm:"null" name:"meta_title"`
	MetaDescription string `gorm:"type:text;null" name:"meta_description"`
	MetaKeywords    string `gorm:"type:text;null" name:"meta_keywords"`
}
