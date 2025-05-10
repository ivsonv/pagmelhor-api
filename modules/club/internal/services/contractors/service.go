package contractors

import (
	"net/http"

	"app/modules/club/domain/interfaces/repository"
	"app/modules/club/domain/interfaces/services"
	"app/modules/club/domain/results"
)

type ContractorService struct {
	contractorRepository repository.IContractorRepository
}

func NewContractorService(contractorRepository repository.IContractorRepository) services.IContractorServices {
	return &ContractorService{contractorRepository: contractorRepository}
}

var (
	ErrCreateContractor  = results.NewError("CREATE_CONTRACTOR_ERROR", "error creating contractor", http.StatusUnprocessableEntity)
	ErrInvalidEntity     = results.NewError("INVALID_ENTITY", "invalid to mapentity", http.StatusBadRequest)
	ErrTimeoutOrCanceled = results.NewError("TIMEOUT_OR_CANCELED_CONTRACTOR_ERROR", "operation contractor timeout or canceled", http.StatusRequestTimeout)
	ErrCpfCnpjDuplicated = results.NewError("CPF_CNPJ_DUPLICATED", "cpf/cnpj already exists", http.StatusConflict)
	ErrEmailDuplicated   = results.NewError("EMAIL_DUPLICATED", "email already exists", http.StatusConflict)
	ErrSlugDuplicated    = results.NewError("SLUG_DUPLICATED", "slug already exists", http.StatusConflict)
)
