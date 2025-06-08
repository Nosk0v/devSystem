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

func (s *CourseService) IsCourseCompleted(userEmail string, courseID int) (bool, error) {
	isCompleted, err := s.repo.IsCourseCompleted(userEmail, courseID)
	if err != nil {
		return false, fmt.Errorf("error checking if course is completed: %w", err)
	}
	return isCompleted, nil
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

func (s *CourseService) GetCoursesByOrganization(organizationID int) ([]models.CourseResponse, error) {
	courses, err := s.repo.GetCoursesByOrganization(organizationID)
	if err != nil {
		return nil, fmt.Errorf("error getting courses by organization: %w", err)
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

func (s *CourseService) GetUserCourseProgress(userEmail string, courseID int) ([]int, error) {
	progress, err := s.repo.GetUserCourseProgress(userEmail, courseID)
	if err != nil {
		return nil, fmt.Errorf("error getting user course progress: %w", err)
	}
	return progress, nil
}

func (s *CourseService) MarkMaterialAsCompleted(userEmail string, courseID int, materialID int) error {
	if err := s.repo.MarkMaterialAsCompleted(userEmail, courseID, materialID); err != nil {
		return fmt.Errorf("error marking material as completed: %w", err)
	}
	return nil
}

func (s *CourseService) CompleteCourse(userEmail string, courseID int) error {
	if err := s.repo.CompleteCourse(userEmail, courseID); err != nil {
		return fmt.Errorf("error marking course as completed: %w", err)
	}
	return nil
}

func (s *CourseService) GetCourseProgressByOrganization(orgID int) ([]models.UserCourseProgress, error) {
	progress, err := s.repo.GetCourseProgressByOrganization(orgID)
	if err != nil {
		return nil, fmt.Errorf("error getting course progress for organization: %w", err)
	}
	return progress, nil
}

func (s *CourseService) GetCompletedCourses(userEmail string) ([]models.CourseResponse, error) {
	courses, err := s.repo.GetCompletedCourses(userEmail)
	if err != nil {
		return nil, fmt.Errorf("error getting completed courses: %w", err)
	}
	return courses, nil
}

func (s *CourseService) GetCoursesByDepartment(orgID int, departmentID int) ([]models.CourseResponse, error) {
	courses, err := s.repo.GetCoursesByDepartment(departmentID, orgID)
	if err != nil {
		return nil, fmt.Errorf("error getting courses by department: %w", err)
	}
	return courses, nil
}
