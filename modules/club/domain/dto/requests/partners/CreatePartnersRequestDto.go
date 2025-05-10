package partners

import "app/modules/club/domain/entities"

type CreatePartnerRequestDto struct {
	Name            string `json:"name" validate:"required,min=3"`
	CpfCnpj         string `json:"cpf_cnpj" validate:"required,cpf_cnpj"`
	Email           string `json:"email" validate:"required,email"`
	Slug            string `json:"slug" validate:"required"`
	Status          int16  `json:"status" validate:"required"`
	Phone           string `json:"phone" validate:"omitempty,phone"`
	Password        string `json:"password"`
	Image           string `json:"image"`
	Address         string `json:"address"`
	City            string `json:"city"`
	State           string `json:"state"`
	ZipCode         string `json:"zip_code"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
}

func (dto CreatePartnerRequestDto) ToMapEntity() *entities.PartnerEntity {
	return &entities.PartnerEntity{
		Name:            dto.Name,
		CpfCnpj:         dto.CpfCnpj,
		Email:           dto.Email,
		Slug:            dto.Slug,
		Status:          dto.Status,
		Phone:           dto.Phone,
		Password:        dto.Password,
		Image:           dto.Image,
		Address:         dto.Address,
		City:            dto.City,
		State:           dto.State,
		ZipCode:         dto.ZipCode,
		Description:     dto.Description,
		MetaTitle:       dto.MetaTitle,
		MetaDescription: dto.MetaDescription,
		MetaKeywords:    dto.MetaKeywords,
	}
}
