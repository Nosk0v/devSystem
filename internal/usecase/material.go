package usecase

import (
	"devSystem/models"
	"errors"
	"log"
)

func (u *Usecase) CreateMaterial(material models.Material) error {
	log.Println("Attempting to create material:", material.Title)
	if material.Title == "" {
		log.Println("Failed to create material: title is required")
		return errors.New("material name is required")
	}
	log.Printf("Material to be passed to service: %+v\n", material)
	if err := u.services.CreateMaterial(material); err != nil {
		log.Println("Error creating material:", err)
		return err
	}
	log.Println("Successfully created material:", material.Title)
	return nil
}

func (u *Usecase) GetAllMaterials() ([]models.Material, error) {
	log.Println("Fetching all materials")
	materials, err := u.services.GetAllMaterials()
	if err != nil {
		log.Println("Error fetching all materials:", err)
		return nil, err
	}
	log.Println("Successfully fetched all materials, count:", len(materials))
	return materials, nil
}

func (u *Usecase) GetMaterial(id int) (*models.Material, error) {
	log.Printf("Fetching material with ID: %d\n", id)
	material, err := u.services.GetMaterialByID(id)
	if err != nil {
		log.Println("Error fetching material:", err)
		return nil, err
	}
	log.Printf("Successfully fetched material: %+v\n", material)
	return &material, nil
}

func (u *Usecase) UpdateMaterial(material models.Material) error {
	log.Printf("Attempting to update material with ID %d\n", material.MaterialID)
	if material.Title == "" {
		log.Println("Failed to update material: title is required")
		return errors.New("material name is required")
	}
	if err := u.services.UpdateMaterial(material); err != nil {
		log.Println("Error updating material:", err)
		return err
	}
	log.Printf("Successfully updated material with ID %d\n", material.MaterialID)
	return nil
}

func (u *Usecase) DeleteMaterial(id int) error {
	log.Printf("Attempting to delete material with ID %d\n", id)
	if err := u.services.DeleteMaterial(id); err != nil {
		log.Println("Error deleting material:", err)
		return err
	}
	log.Printf("Successfully deleted material with ID %d\n", id)
	return nil
}
