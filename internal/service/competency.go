package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"fmt"
	"log"
)

type CompetencyService struct {
	repo *repository.CompetencyRepository
}

func NewCompetencyService(repo *repository.CompetencyRepository) *CompetencyService {
	return &CompetencyService{repo: repo}
}

func (c *CompetencyService) CreateCompetency(comp models.Competency) error {
	log.Printf("Received competency data: %+v\n", comp)
	return c.repo.CreateCompetency(comp)
}

func (c *CompetencyService) GetAllCompetencies() ([]models.Competency, error) {
	competencies, err := c.repo.GetAllCompetencies()
	if err != nil {
		return nil, fmt.Errorf("error getting competencies: %w", err)
	}
	return competencies, nil
}

func (c *CompetencyService) UpdateCompetency(comp models.Competency) error {
	return c.repo.UpdateCompetency(comp)
}

func (c *CompetencyService) DeleteCompetency(id int) error {
	return c.repo.DeleteCompetency(id)
}
