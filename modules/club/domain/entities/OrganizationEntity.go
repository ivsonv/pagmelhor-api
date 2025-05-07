package entities

import (
	"app/modules/club/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u OrganizationEntity) TableName() string {
	return domain.SchemaClubName + ".organizations"
}

type OrganizationEntity struct {
	BaseEntity
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Phone    string `json:"phone" gorm:"null"`
	Slug     string `json:"slug" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

func (o *OrganizationEntity) BeforeSave(tx *gorm.DB) error {
	if o.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		o.Password = string(hashedPassword)
	}
	return nil
}

func (o *OrganizationEntity) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(password))
	return err == nil
}

func (o *OrganizationEntity) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	o.Password = string(hashedPassword)
	return nil
}
