package entities

import (
	"app/modules/club/domain"
	"app/modules/club/domain/enums"
)

func (u UsersBenefitUsageEntity) TableName() string {
	return domain.SchemaClubName + ".users_benefit_usage"
}

type UsersBenefitUsageEntity struct {
	BaseEntity
	VoucherCode   string                   `json:"voucher_code" gorm:"not null"`
	Description   string                   `json:"description" gorm:"type:text"`
	DiscountType  enums.DiscountType       `json:"discount_type" gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')"`
	DiscountValue float64                  `json:"discount_value" gorm:"type:decimal(10,2)"`
	Notes         string                   `json:"notes" gorm:"type:text"`
	Status        enums.BenefitUsageStatus `json:"status" gorm:"type:varchar(10);check:status IN ('used', 'reserved', 'cancelled')"`
	UserID        int                      `json:"user_id" gorm:"not null;index"`
	User          UserEntity               `json:"user" gorm:"foreignKey:UserID"`
	PartnerID     int                      `json:"partner_id" gorm:"not null;index"`
	Partner       PartnersEntity           `json:"partner" gorm:"foreignKey:PartnerID"`
	BenefitID     int                      `json:"benefit_id" gorm:"not null;index"`
	Benefit       BenefitEntity            `json:"benefit" gorm:"foreignKey:BenefitID"`
	BenefitItemID *int                     `json:"benefit_item_id" gorm:"index"`
	BenefitItem   *BenefitItemEntity       `json:"benefit_item,omitempty" gorm:"foreignKey:BenefitItemID"`
}
