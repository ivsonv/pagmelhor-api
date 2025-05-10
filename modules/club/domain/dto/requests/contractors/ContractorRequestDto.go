package contractors

import "app/modules/club/domain/entities"

type ContractorRequestDto struct {
	Name            string `json:"name" validate:"required"`
	CpfOrCnpj       string `json:"cpf_cnpj" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Slug            string `json:"slug" validate:"required"`
	Phone           string `json:"phone"`
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

func (dto ContractorRequestDto) ToMapEntity() *entities.ContractorEntity {
	return &entities.ContractorEntity{
		Name:            dto.Name,
		CpfOrCnpj:       dto.CpfOrCnpj,
		Email:           dto.Email,
		Slug:            dto.Slug,
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
