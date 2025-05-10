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
	ErrCreateContractor = results.NewError("CREATE_CONTRACTOR_ERROR", "error creating contractor", http.StatusInternalServerError)
)
