package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"fmt"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) CreateCourse(course models.Course) (int, error) {
	courseID, err := s.repo.CreateCourse(course)
	if err != nil {
		return 0, fmt.Errorf("error creating course: %w", err)
	}
	return courseID, nil
}

func (s *CourseService) LinkCourseWithMaterials(courseID int, materialIDs []int) error {
	if err := s.repo.LinkCourseWithMaterials(courseID, materialIDs); err != nil {
		return fmt.Errorf("error linking course with materials: %w", err)
	}
	return nil
}

func (s *CourseService) LinkCourseWithCompetencies(courseID int, competencyIDs []int) error {
	if err := s.repo.LinkCourseWithCompetencies(courseID, competencyIDs); err != nil {
		return fmt.Errorf("error linking course with competencies: %w", err)
	}
	return nil
}

func (s *CourseService) GetCourseByID(id int) (models.CourseResponse, error) {
	course, err := s.repo.GetCourseByID(id)
	if err != nil {
		return models.CourseResponse{}, fmt.Errorf("error getting course by ID: %w", err)
	}
	return course, nil
}

func (s *CourseService) GetAllCourses() ([]models.CourseResponse, error) {
	courses, err := s.repo.GetAllCourses()
	if err != nil {
		return nil, fmt.Errorf("error getting all courses: %w", err)
	}
	return courses, nil
}

func (s *CourseService) UpdateCourse(course models.Course) error {
	if err := s.repo.UpdateCourse(course); err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	return nil
}

func (s *CourseService) DeleteCourse(id int) error {
	if err := s.repo.DeleteCourse(id); err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}
