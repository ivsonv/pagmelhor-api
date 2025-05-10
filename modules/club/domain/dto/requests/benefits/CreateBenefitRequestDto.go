package benefits

import (
	"app/modules/club/domain/entities"
	"app/modules/club/domain/enums"
)

type CreateBenefitRequestDto struct {
	Name          string             `json:"name" validate:"required"`
	Description   string             `json:"description"`
	OriginalPrice float64            `json:"original_price"`
	DiscountType  enums.DiscountType `json:"discount_type" validate:"required"`
	DiscountValue float64            `json:"discount_value"`
	Status        int16              `json:"status" validate:"required"`
	Notes         string             `json:"notes"`
	ContractorID  *int               `json:"contractor_id"`
	PartnerID     *int               `json:"partner_id"`
}

func (dto CreateBenefitRequestDto) ToMapEntity() *entities.BenefitEntity {
	return &entities.BenefitEntity{
		Name:          dto.Name,
		Description:   dto.Description,
		OriginalPrice: dto.OriginalPrice,
		DiscountType:  dto.DiscountType,
		DiscountValue: dto.DiscountValue,
		Status:        dto.Status,
		Notes:         dto.Notes,
		ContractorID:  dto.ContractorID,
		PartnerID:     dto.PartnerID,
	}
}
