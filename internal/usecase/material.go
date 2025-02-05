package usecase

import (
	"devSystem/models"
	"fmt"
	"log"
	"strings"
)

func (u *Usecase) CreateMaterial(material models.Material) (int, error) {
	if material.Title == "" {
		return 0, fmt.Errorf("material title is required")
	}

	materialID, err := u.services.Material.CreateMaterial(material)
	if err != nil {
		return 0, fmt.Errorf("error creating material: %w", err)
	}
	if len(material.Competencies) > 0 {
		err = u.services.Material.LinkMaterialWithCompetencies(materialID, material.Competencies)
		if err != nil {
			return 0, fmt.Errorf("error linking material with competencies: %w", err)
		}
	}

	return materialID, nil
}

func (u *Usecase) CreateMaterialType(materialType models.MaterialType) (int, error) {
	if materialType.Type == "" {
		return 0, fmt.Errorf("material type name is required")
	}
	materialTypeID, err := u.services.Material.CreateMaterialType(materialType)
	if err != nil {
		return 0, fmt.Errorf("error creating material type: %w", err)
	}

	return materialTypeID, nil
}

func (u *Usecase) GetAllMaterials() ([]models.MaterialResponse, error) {
	// Логируем запрос к данным
	log.Printf("Fetching all materials from the service layer.")

	materials, err := u.services.Material.GetAllMaterials()
	if err != nil {
		log.Printf("Error fetching all materials in usecase: %v", err) // Логируем ошибку
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}

	var response []models.MaterialResponse
	for _, material := range materials {
		// Логируем каждое возвращаемое значение материала
		log.Printf("Fetched material: %+v", material)

		response = append(response, models.MaterialResponse{
			MaterialID:   material.MaterialID,
			Title:        material.Title,
			Description:  material.Description,
			TypeName:     material.TypeName,
			Content:      material.Content,
			Competencies: material.Competencies,
			CreateDate:   material.CreateDate,
		})
	}

	// Логируем итоговый список материалов
	log.Printf("Total materials fetched: %d", len(response))

	return response, nil
}

func (u *Usecase) GetAllMaterialTypes() ([]models.MaterialType, error) {
	materialTypes, err := u.services.Material.GetAllMaterialTypes()
	if err != nil {
		return nil, fmt.Errorf("error fetching all materialTypes: %w", err)
	}
	var response []models.MaterialType
	for _, materialTypes := range materialTypes {
		response = append(response, models.MaterialType{
			TypeID: materialTypes.TypeID,
			Type:   materialTypes.Type,
		})

	}
	return response, nil
}

func (u *Usecase) GetMaterial(id int) (*models.MaterialResponse, error) {
	if id <= 0 {
		log.Printf("Invalid material ID: %d", id) // Логируем некорректный ID
		return nil, fmt.Errorf("invalid material ID: %d", id)
	}

	log.Printf("Fetching material with ID: %d", id)

	material, err := u.services.Material.GetMaterialByID(id)
	if err != nil {
		log.Printf("Error fetching material with ID %d: %v", id, err) // Логируем ошибку
		return nil, fmt.Errorf("error fetching material with ID %d: %w", id, err)
	}
	log.Printf("Fetched material: %+v", material)

	response := &models.MaterialResponse{
		MaterialID:   material.MaterialID,
		Title:        material.Title,
		Description:  material.Description,
		TypeName:     material.TypeName,
		Content:      material.Content,
		Competencies: material.Competencies,
		CreateDate:   material.CreateDate,
	}

	return response, nil
}

func (u *Usecase) GetMaterialType(id int) (*models.MaterialType, error) {
	if id <= 0 {
		log.Printf("Invalid material type ID: %d", id) // Логируем некорректный ID
		return nil, fmt.Errorf("invalid material type ID: %d", id)
	}

	log.Printf("Fetching material type with ID: %d", id)

	materialType, err := u.services.Material.GetMaterialTypeByID(id)

	if err != nil {
		log.Printf("Error fetching material type with ID %d: %v", id, err)
		return nil, fmt.Errorf("error fetching material type with ID %d: %w", id, err)
	}

	log.Printf("Fetched material type: %+v", materialType)

	response := &models.MaterialType{
		TypeID: materialType.TypeID,
		Type:   materialType.Type,
	}

	return response, nil

}

func (u *Usecase) UpdateMaterial(material models.Material) error {
	err := u.services.Material.UpdateMaterial(material)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil
		}
		return fmt.Errorf("error updating material: %w", err)
	}
	return nil
}

func (u *Usecase) DeleteMaterial(id int) error {
	if err := u.services.Material.DeleteMaterial(id); err != nil {
		return fmt.Errorf("error deleting material: %w", err)
	}
	return nil
}
