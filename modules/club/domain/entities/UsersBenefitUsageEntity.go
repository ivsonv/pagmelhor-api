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
	VoucherCode   string                   `gorm:"not null" name:"voucher_code"`
	Description   string                   `gorm:"type:text" name:"description"`
	DiscountType  enums.DiscountType       `gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')" name:"discount_type"`
	DiscountValue float64                  `gorm:"type:decimal(10,2)" name:"discount_value"`
	Notes         string                   `gorm:"type:text" name:"notes"`
	Status        enums.BenefitUsageStatus `gorm:"type:varchar(10);check:status IN ('used', 'reserved', 'cancelled')" name:"status"`
	UserID        int                      `gorm:"not null;index" name:"user_id"`
	User          UserEntity               `gorm:"foreignKey:UserID"`
	PartnerID     int                      `gorm:"not null;index" name:"partner_id"`
	Partner       PartnerEntity            `gorm:"foreignKey:PartnerID"`
	BenefitID     int                      `gorm:"not null;index" name:"benefit_id"`
	Benefit       BenefitEntity            `gorm:"foreignKey:BenefitID"`
	BenefitItemID *int                     `gorm:"index" name:"benefit_item_id"`
	BenefitItem   *BenefitItemEntity       `gorm:"foreignKey:BenefitItemID"`
}
