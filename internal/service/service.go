package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMaterial(material models.Material) error {
	return s.repo.CreateMaterial(material)
}

func (s *Service) GetMaterialByID(id int) (models.Material, error) {
	return s.repo.GetMaterialByID(id)
}

func (s *Service) UpdateMaterial(material models.Material) error {
	return s.repo.UpdateMaterial(material)
}

func (s *Service) DeleteMaterial(id int) error {
	return s.repo.DeleteMaterial(id)
}
func (s *Service) GetAllMaterials() ([]models.Material, error) {
	return s.repo.GetAllMaterials()
}

func (s *Service) GetAllCompetencies() ([]models.Competency, error) {
	return s.repo.GetAllCompetencies()
}

func (s *Service) CreateCompetency(comp models.Competency) error {
	return s.repo.CreateCompetency(comp)
}

func (s *Service) UpdateCompetency(comp models.Competency) error {
	return s.repo.UpdateCompetency(comp)
}

func (s *Service) DeleteCompetency(id int) error {
	return s.repo.DeleteCompetency(id)
}
