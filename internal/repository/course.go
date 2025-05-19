package repository

import (
	"devSystem/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) CreateCourse(course models.Course) (int, error) {
	query := `
		INSERT INTO "Course" (title, description, created_by, create_date)
		VALUES ($1, $2, $3, $4)
		RETURNING course_id
	`
	var courseID int
	err := r.db.QueryRow(query, course.Title, course.Description, course.CreatedBy, course.CreateDate).Scan(&courseID)
	if err != nil {
		return 0, fmt.Errorf("error creating course: %w", err)
	}

	return courseID, nil
}

func (r *CourseRepository) LinkCourseWithMaterials(courseID int, materialIDs []int) error {
	query := `INSERT INTO "CourseMaterial" (course_id, material_id) VALUES ($1, $2)`
	for _, materialID := range materialIDs {
		_, err := r.db.Exec(query, courseID, materialID)
		if err != nil {
			return fmt.Errorf("error linking course with material ID %d: %w", materialID, err)
		}
	}
	return nil
}

func (r *CourseRepository) LinkCourseWithCompetencies(courseID int, competencyIDs []int) error {
	query := `INSERT INTO "CourseCompetency" (course_id, competency_id) VALUES ($1, $2)`
	for _, competencyID := range competencyIDs {
		_, err := r.db.Exec(query, courseID, competencyID)
		if err != nil {
			return fmt.Errorf("error linking course with competency ID %d: %w", competencyID, err)
		}
	}
	return nil
}

func (r *CourseRepository) GetCourseByID(id int) (models.CourseResponse, error) {
	var course models.CourseResponse
	query := `
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies
		FROM "Course" c
		LEFT JOIN "CourseMaterial" cm ON c.course_id = cm.course_id
		LEFT JOIN "Material" m ON cm.material_id = m.material_id
		LEFT JOIN "CourseCompetency" cc ON c.course_id = cc.course_id
		LEFT JOIN "Competency" comp ON cc.competency_id = comp.competency_id
		WHERE c.course_id = $1
		GROUP BY c.course_id
	`
	err := r.db.Get(&course, query, id)
	if err != nil {
		return models.CourseResponse{}, fmt.Errorf("error fetching course by ID: %w", err)
	}
	return course, nil
}

func (r *CourseRepository) UpdateCourse(course models.Course) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := `
		UPDATE "Course"
		SET title = $1, description = $2, created_by = $3
		WHERE course_id = $4
	`
	_, err = tx.Exec(query, course.Title, course.Description, course.CreatedBy, course.CourseID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update course: %w", err)
	}

	_, err = tx.Exec(`DELETE FROM "CourseMaterial" WHERE course_id = $1`, course.CourseID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error clearing old materials: %w", err)
	}
	_, err = tx.Exec(`DELETE FROM "CourseCompetency" WHERE course_id = $1`, course.CourseID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error clearing old competencies: %w", err)
	}

	for _, mid := range course.Materials {
		_, err := tx.Exec(`INSERT INTO "CourseMaterial" (course_id, material_id) VALUES ($1, $2)`, course.CourseID, mid)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting new material link: %w", err)
		}
	}
	for _, cid := range course.Competencies {
		_, err := tx.Exec(`INSERT INTO "CourseCompetency" (course_id, competency_id) VALUES ($1, $2)`, course.CourseID, cid)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting new competency link: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit course update: %w", err)
	}

	return nil
}

func (r *CourseRepository) GetAllCourses() ([]models.CourseResponse, error) {
	var courses []models.CourseResponse
	query := `
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies
		FROM "Course" c
		LEFT JOIN "CourseMaterial" cm ON c.course_id = cm.course_id
		LEFT JOIN "Material" m ON cm.material_id = m.material_id
		LEFT JOIN "CourseCompetency" cc ON c.course_id = cc.course_id
		LEFT JOIN "Competency" comp ON cc.competency_id = comp.competency_id
		GROUP BY c.course_id
	`
	err := r.db.Select(&courses, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all courses: %w", err)
	}
	return courses, nil
}

func (r *CourseRepository) DeleteCourse(id int) error {
	query := `DELETE FROM "Course" WHERE course_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}
