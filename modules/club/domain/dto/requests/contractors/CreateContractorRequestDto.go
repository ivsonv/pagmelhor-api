package contractors

import "app/modules/club/domain/entities"

type CreateContractorRequestDto struct {
	Name            string `json:"name" validate:"required,min=3,max=100"`
	CpfCnpj         string `json:"cpf_cnpj" validate:"required,cpf_cnpj"`
	Email           string `json:"email" validate:"required,email"`
	Slug            string `json:"slug" validate:"required"`
	Phone           string `json:"phone" validate:"omitempty,phone"`
	Password        string `json:"password"`
	Image           string `json:"image"`
	Address         string `json:"address" validate:"omitempty,max=200"`
	City            string `json:"city" validate:"omitempty,max=100"`
	State           string `json:"state" validate:"omitempty,max=2"`
	ZipCode         string `json:"zip_code" validate:"omitempty,max=10"`
	Description     string `json:"description" validate:"omitempty,max=1000"`
	MetaTitle       string `json:"meta_title" validate:"omitempty,max=100"`
	MetaDescription string `json:"meta_description" validate:"omitempty,max=200"`
	MetaKeywords    string `json:"meta_keywords" validate:"omitempty,max=200"`
}

func (dto CreateContractorRequestDto) ToMapEntity() *entities.ContractorEntity {
	return &entities.ContractorEntity{
		Name:            dto.Name,
		CpfCnpj:         dto.CpfCnpj,
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
