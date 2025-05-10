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
	Name          string              `gorm:"not null" name:"name"`
	Description   string              `gorm:"type:text;null" name:"description"`
	OriginalPrice float64             `gorm:"type:decimal(10,2);null" name:"original_price"`
	DiscountType  enums.DiscountType  `gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')" name:"discount_type"`
	DiscountValue float64             `gorm:"type:decimal(10,2);null" name:"discount_value"`
	Status        int16               `gorm:"not null;default:1" name:"status"`
	Notes         string              `gorm:"type:text;null" name:"notes"`
	ContractorID  *int                `gorm:"null" name:"contractor_id"`
	Contractor    *ContractorEntity   `gorm:"foreignKey:ContractorID" name:"contractor,omitempty"`
	PartnerID     *int                `gorm:"null" name:"partner_id"`
	Partner       *PartnersEntity     `gorm:"foreignKey:PartnerID" name:"partner,omitempty"`
	BenefitItems  []BenefitItemEntity `gorm:"foreignKey:BenefitID" name:"benefit_items"`
}
