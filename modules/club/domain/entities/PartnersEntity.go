package entities

import (
	"app/modules/club/domain"
)

func (p PartnersEntity) TableName() string {
	return domain.SchemaClubName + ".partners"
}

type PartnersEntity struct {
	BaseEntity
	Name            string `json:"name" gorm:"not null"`
	CpfOrCnpj       string `json:"cpf_cnpj" gorm:"not null;unique"`
	Email           string `json:"email" gorm:"not null;unique"`
	Slug            string `json:"slug" gorm:"not null;unique"`
	Status          int16  `json:"status" gorm:"not null;default:1"`
	Phone           string `json:"phone" gorm:"null"`
	Password        string `json:"password" gorm:"null"`
	Image           string `json:"image" gorm:"null"`
	Address         string `json:"address" gorm:"null"`
	City            string `json:"city" gorm:"null"`
	State           string `json:"state" gorm:"null"`
	ZipCode         string `json:"zip_code" gorm:"null"`
	Description     string `json:"description" gorm:"type:text;null"`
	MetaTitle       string `json:"meta_title" gorm:"null"`
	MetaDescription string `json:"meta_description" gorm:"type:text;null"`
	MetaKeywords    string `json:"meta_keywords" gorm:"type:text;null"`
}
