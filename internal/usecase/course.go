package usecase

import (
	"devSystem/models"
	"fmt"
	"log"
)

const (
	RoleAdmin      = 0
	RoleUser       = 1
	RoleSuperAdmin = 2
)

func (u *Usecase) CreateCourse(course models.Course) (int, error) {
	if course.Title == "" {
		return 0, fmt.Errorf("course title is required")
	}
	if course.OrganizationID <= 0 {
		return 0, fmt.Errorf("organization is required")
	}

	courseID, err := u.services.Course.CreateCourse(course)
	if err != nil {
		return 0, fmt.Errorf("error creating course: %w", err)
	}

	if len(course.Competencies) > 0 {
		if err := u.services.Course.LinkCourseWithCompetencies(courseID, course.Competencies); err != nil {
			return 0, fmt.Errorf("error linking course with competencies: %w", err)
		}
	}

	if len(course.Materials) > 0 {
		if err := u.services.Course.LinkCourseWithMaterials(courseID, course.Materials); err != nil {
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

	return &models.CourseResponse{
		CourseID:       course.CourseID,
		Title:          course.Title,
		Description:    course.Description,
		CreatedBy:      course.CreatedBy,
		Competencies:   course.Competencies,
		Materials:      course.Materials,
		CreateDate:     course.CreateDate,
		OrganizationID: course.OrganizationID,
	}, nil
}

func (u *Usecase) GetCoursesByClaims(claims *models.JWTClaims) ([]models.CourseResponse, error) {
	log.Printf("Fetching courses by role: %d and org: %v", claims.Role, claims.OrganizationID)

	switch claims.Role {
	case RoleSuperAdmin:
		return u.getAllCourses()
	case RoleAdmin, RoleUser:
		if claims.OrganizationID == nil || *claims.OrganizationID <= 0 {
			return nil, fmt.Errorf("organization is required for this role")
		}
		return u.getCoursesByOrganization(*claims.OrganizationID)
	default:
		return nil, fmt.Errorf("unauthorized role: %d", claims.Role)
	}
}

func (u *Usecase) UpdateCourse(course models.Course) error {
	if course.OrganizationID <= 0 {
		return fmt.Errorf("organization is required for update")
	}
	if err := u.services.Course.UpdateCourse(course); err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	return nil
}

func (u *Usecase) DeleteCourse(id int) error {
	if err := u.services.Course.DeleteCourse(id); err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}

func (u *Usecase) getAllCourses() ([]models.CourseResponse, error) {
	log.Printf("Fetching all courses (superadmin only)")
	courses, err := u.services.Course.GetAllCourses()
	if err != nil {
		return nil, fmt.Errorf("error fetching all courses: %w", err)
	}
	return courses, nil
}

func (u *Usecase) getCoursesByOrganization(orgID int) ([]models.CourseResponse, error) {
	log.Printf("Fetching courses for organization: %d", orgID)
	courses, err := u.services.Course.GetCoursesByOrganization(orgID)
	if err != nil {
		return nil, fmt.Errorf("error fetching courses by organization: %w", err)
	}
	return courses, nil
}

// Progress-related methods
func (u *Usecase) IsCourseCompleted(userEmail string, courseID int) (bool, error) {
	completed, err := u.services.Course.IsCourseCompleted(userEmail, courseID)
	if err != nil {
		return false, fmt.Errorf("error checking if course is completed: %w", err)
	}
	return completed, nil
}
func (u *Usecase) GetUserCourseProgress(userEmail string, courseID int) ([]int, error) {
	progress, err := u.services.Course.GetUserCourseProgress(userEmail, courseID)
	if err != nil {
		return nil, fmt.Errorf("error getting user course progress: %w", err)
	}
	return progress, nil
}

func (u *Usecase) MarkMaterialAsCompleted(userEmail string, courseID int, materialID int) error {
	if err := u.services.Course.MarkMaterialAsCompleted(userEmail, courseID, materialID); err != nil {
		return fmt.Errorf("error marking material as completed: %w", err)
	}
	return nil
}

func (u *Usecase) CompleteCourse(userEmail string, courseID int) error {
	if err := u.services.Course.CompleteCourse(userEmail, courseID); err != nil {
		return fmt.Errorf("error marking course as completed: %w", err)
	}
	return nil
}

func (u *Usecase) GetCompletedCourses(userEmail string) ([]models.CourseResponse, error) {
	courses, err := u.services.Course.GetCompletedCourses(userEmail)
	if err != nil {
		return nil, fmt.Errorf("error fetching completed courses: %w", err)
	}
	return courses, nil
}
