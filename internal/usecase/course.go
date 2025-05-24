package usecase

import (
	"devSystem/models"
	"fmt"
	"log"
)

func (u *Usecase) CreateCourse(course models.Course) (int, error) {
	if course.Title == "" {
		return 0, fmt.Errorf("course title is required")
	}

	courseID, err := u.services.Course.CreateCourse(course)
	if err != nil {
		return 0, fmt.Errorf("error creating course: %w", err)
	}

	if len(course.Competencies) > 0 {
		err = u.services.Course.LinkCourseWithCompetencies(courseID, course.Competencies)
		if err != nil {
			return 0, fmt.Errorf("error linking course with competencies: %w", err)
		}
	}

	if len(course.Materials) > 0 {
		err = u.services.Course.LinkCourseWithMaterials(courseID, course.Materials)
		if err != nil {
			return 0, fmt.Errorf("error linking course with materials: %w", err)
		}
	}

	return courseID, nil
}

func (u *Usecase) GetCourse(id int) (*models.CourseResponse, error) {
	if id <= 0 {
		log.Printf("Invalid course ID: %d", id)
		return nil, fmt.Errorf("invalid course ID: %d", id)
	}

	log.Printf("Fetching course with ID: %d", id)

	course, err := u.services.Course.GetCourseByID(id)
	if err != nil {
		log.Printf("Error fetching course with ID %d: %v", id, err)
		return nil, fmt.Errorf("error fetching course with ID %d: %w", id, err)
	}

	log.Printf("Fetched course: %+v", course)

	response := &models.CourseResponse{
		CourseID:     course.CourseID,
		Title:        course.Title,
		Description:  course.Description,
		CreatedBy:    course.CreatedBy,
		Competencies: course.Competencies,
		Materials:    course.Materials,
		CreateDate:   course.CreateDate,
	}

	return response, nil
}

func (u *Usecase) GetAllCourses() ([]models.CourseResponse, error) {
	log.Printf("Fetching all courses from the service layer.")

	courses, err := u.services.Course.GetAllCourses()
	if err != nil {
		log.Printf("Error fetching all courses in usecase: %v", err)
		return nil, fmt.Errorf("error fetching all courses: %w", err)
	}

	var response []models.CourseResponse
	for _, course := range courses {
		log.Printf("Fetched course: %+v", course)

		response = append(response, models.CourseResponse{
			CourseID:     course.CourseID,
			Title:        course.Title,
			Description:  course.Description,
			CreatedBy:    course.CreatedBy,
			Competencies: course.Competencies,
			Materials:    course.Materials,
			CreateDate:   course.CreateDate,
		})
	}

	log.Printf("Total courses fetched: %d", len(response))

	return response, nil
}

func (u *Usecase) UpdateCourse(course models.Course) error {
	err := u.services.Course.UpdateCourse(course)
	if err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	return nil
}

func (u *Usecase) DeleteCourse(id int) error {
	err := u.services.Course.DeleteCourse(id)
	if err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}
