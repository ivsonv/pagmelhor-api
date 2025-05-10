package entities

import (
	"app/modules/club/domain"
	"app/modules/club/domain/enums"
)

func (b BenefitItemEntity) TableName() string {
	return domain.SchemaClubName + ".benefit_items"
}

type BenefitItemEntity struct {
	BaseEntity
	Name             string             `gorm:"not null" name:"name"`
	CoveragePercent  float64            `gorm:"type:decimal(5,2)" name:"coverage_percent"`
	LimitPerDay      float64            `gorm:"type:decimal(10,2)" name:"limit_per_day"`
	LimitTotal       float64            `gorm:"type:decimal(10,2)" name:"limit_total"`
	MaxCoverageValue float64            `gorm:"type:decimal(10,2)" name:"max_coverage_value"`
	Unlimited        bool               `gorm:"default:false" name:"unlimited"`
	Notes            string             `gorm:"type:text" name:"notes"`
	Status           int16              `gorm:"not null;default:1" name:"status"`
	OriginalPrice    float64            `gorm:"type:decimal(10,2)" name:"original_price"`
	DiscountType     enums.DiscountType `gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')" name:"discount_type"`
	DiscountValue    float64            `gorm:"type:decimal(10,2)" name:"discount_value"`
	BenefitID        int                `gorm:"not null" name:"benefit_id"`
	Benefit          BenefitEntity      `gorm:"foreignKey:BenefitID" name:"benefit,omitempty"`
}
