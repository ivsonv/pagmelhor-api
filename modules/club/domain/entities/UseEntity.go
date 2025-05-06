package entities

import (
	"app/modules/club/domain"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func (u UserEntity) TableName() string {
	return domain.SchemaClubName + ".users"
}

type UserEntity struct {
	BaseEntity
	Email    string `json:"email" gorm:"index;unique;null"`
	Name     string `json:"name" gorm:"not null"`
	Phone    string `json:"phone" gorm:"null"`
	Password string `json:"-" gorm:"null"`
}

func (u *UserEntity) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *UserEntity) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *UserEntity) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
