package entities

import (
	"app/modules/club/domain"
	"app/modules/club/domain/enums"
)

func (b BenefitEntity) TableName() string {
	return domain.SchemaClubName + ".benefits"
}

type BenefitEntity struct {
	BaseEntity
	Name          string             `json:"name" gorm:"not null"`
	Description   string             `json:"description" gorm:"type:text;null"`
	OriginalPrice float64            `json:"original_price" gorm:"type:decimal(10,2);null"`
	DiscountType  enums.DiscountType `json:"discount_type" gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')"`
	DiscountValue float64            `json:"discount_value" gorm:"type:decimal(10,2);null"`
	Status        int16              `json:"status" gorm:"not null;default:1"`
	Notes         string             `json:"notes" gorm:"type:text;null"`

	ContractorID *int              `json:"contractor_id" gorm:"null"`
	Contractor   *ContractorEntity `json:"contractor,omitempty" gorm:"foreignKey:ContractorID"`

	PartnerID *int            `json:"partner_id" gorm:"null"`
	Partner   *PartnersEntity `json:"partner,omitempty" gorm:"foreignKey:PartnerID"`

	BenefitItems []BenefitItemEntity `json:"benefit_items" gorm:"foreignKey:BenefitID"`
}
