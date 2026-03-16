package services

import (
	"bloom-wallet/api/request"
	"bloom-wallet/model"
	"bloom-wallet/repository"
)

type ProjectContractService struct {
	ProjectContractRepository *repository.ProjectContractRepository
}

func NewProjectContractService(projectContractRepository *repository.ProjectContractRepository) *ProjectContractService {
	return &ProjectContractService{
		ProjectContractRepository: projectContractRepository,
	}
}

func (s *ProjectContractService) GetById(id uint) (*model.ProjectContract, error) {
	return s.ProjectContractRepository.GetById(id)
}

func (s *ProjectContractService) GetAll() ([]model.ProjectContract, error) {
	return s.ProjectContractRepository.GetAll()
}

func (s *ProjectContractService) Update(request *request.ProjectContractUpdateRequest) error {
	return s.ProjectContractRepository.Update(request)
}
