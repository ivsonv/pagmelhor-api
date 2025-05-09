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
	Name             string             `json:"name" gorm:"not null"`
	CoveragePercent  float64            `json:"coverage_percent" gorm:"type:decimal(5,2)"`
	LimitPerDay      float64            `json:"limit_per_day" gorm:"type:decimal(10,2)"`
	LimitTotal       float64            `json:"limit_total" gorm:"type:decimal(10,2)"`
	MaxCoverageValue float64            `json:"max_coverage_value" gorm:"type:decimal(10,2)"`
	Unlimited        bool               `json:"unlimited" gorm:"default:false"`
	Notes            string             `json:"notes" gorm:"type:text"`
	Status           int16              `json:"status" gorm:"not null;default:1"`
	OriginalPrice    float64            `json:"original_price" gorm:"type:decimal(10,2)"`
	DiscountType     enums.DiscountType `json:"discount_type" gorm:"type:varchar(20);check:discount_type IN ('percent', 'fixed')"`
	DiscountValue    float64            `json:"discount_value" gorm:"type:decimal(10,2)"`
	BenefitID        int                `json:"benefit_id" gorm:"not null"`
	Benefit          BenefitEntity      `json:"benefit,omitempty" gorm:"foreignKey:BenefitID"`
}
