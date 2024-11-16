package usecase

import (
	"devSystem/models"
	"errors"
	"log"
)

func (u *Usecase) CreateCompetency(comp models.Competency) error {
	log.Printf("Usecase received input: %+v\n", comp)

	if comp.Name == "" {
		log.Println("Failed to create competency: name is required")
		return errors.New("competency name is required")
	}

	if err := u.services.CreateCompetency(comp); err != nil {
		log.Println("Error creating competency:", err)
		return err
	}

	log.Printf("Successfully created competency: %+v\n", comp)
	return nil
}

func (u *Usecase) GetAllCompetencies() ([]models.Competency, error) {
	log.Println("Fetching all competencies")
	competencies, err := u.services.GetAllCompetencies()
	if err != nil {
		log.Println("Error fetching all competencies:", err)
		return nil, err
	}
	log.Println("Successfully fetched all competencies, count:", len(competencies))
	return competencies, nil
}

func (u *Usecase) UpdateCompetency(comp models.Competency) error {
	log.Printf("Attempting to update competency with ID %d\n", comp.CompetencyID)
	if comp.Name == "" {
		log.Println("Failed to update competency: name is required")
		return errors.New("competency name is required")
	}
	if err := u.services.UpdateCompetency(comp); err != nil {
		log.Println("Error updating competency:", err)
		return err
	}
	log.Printf("Successfully updated competency with ID %d\n", comp.CompetencyID)
	return nil
}

func (u *Usecase) DeleteCompetency(id int) error {
	log.Printf("Attempting to delete competency with ID %d\n", id)
	if err := u.services.DeleteCompetency(id); err != nil {
		log.Println("Error deleting competency:", err)
		return err
	}
	log.Printf("Successfully deleted competency with ID %d\n", id)
	return nil
}
