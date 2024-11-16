package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"log"
)

type CompetencyService struct {
	repo *repository.Repository
}

func NewCompetencyService(repo *repository.Repository) *CompetencyService {
	return &CompetencyService{repo: repo}
}

func (c *CompetencyService) CreateCompetency(comp models.Competency) error {
	log.Printf("Usecase received input: %+v\n", comp)
	return c.repo.CreateCompetency(comp)
}

func (c *CompetencyService) GetAllCompetencies() ([]models.Competency, error) {
	return c.repo.GetAllCompetencies()
}

func (c *CompetencyService) UpdateCompetency(comp models.Competency) error {
	return c.repo.UpdateCompetency(comp)
}

func (c *CompetencyService) DeleteCompetency(id int) error {
	return c.repo.DeleteCompetency(id)
}
