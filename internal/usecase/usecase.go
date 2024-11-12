package usecase

import (
	"devSystem/internal/service"
	"devSystem/models"
	"errors"
)

type Usecase struct {
	services *service.Service
}

func NewUsecase(services *service.Service) *Usecase {
	return &Usecase{services: services}
}

// CreateMaterial - создает новый материал
func (u *Usecase) CreateMaterial(material models.Material) error {
	if material.Title == "" {
		return errors.New("material name is required")
	}
	if err := u.services.CreateMaterial(material); err != nil {
		return err
	}
	return nil
}

// GetAllMaterials - возвращает все материалы
func (u *Usecase) GetAllMaterials() ([]models.Material, error) {
	materials, err := u.services.GetAllMaterials()
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (u *Usecase) GetMaterial(id int) (*models.Material, error) {
	material, err := u.services.GetMaterialByID(id)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (u *Usecase) UpdateMaterial(material models.Material) error {
	if material.Title == "" {
		return errors.New("material name is required")
	}
	if err := u.services.UpdateMaterial(material); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) DeleteMaterial(id int) error {
	if err := u.services.DeleteMaterial(id); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) CreateCompetency(comp models.Competency) error {
	if comp.Name == "" {
		return errors.New("competency name is required")
	}
	if err := u.services.CreateCompetency(comp); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) GetAllCompetencies() ([]models.Competency, error) {
	competencies, err := u.services.GetAllCompetencies()
	if err != nil {
		return nil, err
	}
	return competencies, nil
}

func (u *Usecase) UpdateCompetency(comp models.Competency) error {
	if comp.Name == "" {
		return errors.New("competency name is required")
	}
	if err := u.services.UpdateCompetency(comp); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) DeleteCompetency(id int) error {
	if err := u.services.DeleteCompetency(id); err != nil {
		return err
	}
	return nil
}
