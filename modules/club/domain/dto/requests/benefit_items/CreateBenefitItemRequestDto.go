package benefit_items

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/enums"
)

type CreateBenefitItemRequestDto struct {
	Name             string             `json:"name" validate:"required"`
	CoveragePercent  float64            `json:"coverage_percent"`
	LimitPerDay      float64            `json:"limit_per_day"`
	LimitTotal       float64            `json:"limit_total"`
	MaxCoverageValue float64            `json:"max_coverage_value"`
	Unlimited        bool               `json:"unlimited"`
	Notes            string             `json:"notes"`
	Status           int16              `json:"status" validate:"required"`
	OriginalPrice    float64            `json:"original_price"`
	DiscountType     enums.DiscountType `json:"discount_type" validate:"required"`
	DiscountValue    float64            `json:"discount_value"`
	BenefitID        int                `json:"benefit_id" validate:"required"`
}

func (dto CreateBenefitItemRequestDto) ToMapEntity() *entities.BenefitItemEntity {
	return &entities.BenefitItemEntity{
		Name:             dto.Name,
		CoveragePercent:  dto.CoveragePercent,
		LimitPerDay:      dto.LimitPerDay,
		LimitTotal:       dto.LimitTotal,
		MaxCoverageValue: dto.MaxCoverageValue,
		Unlimited:        dto.Unlimited,
		Notes:            dto.Notes,
		Status:           dto.Status,
		OriginalPrice:    dto.OriginalPrice,
		DiscountType:     dto.DiscountType,
		DiscountValue:    dto.DiscountValue,
		BenefitID:        dto.BenefitID,
	}
}
