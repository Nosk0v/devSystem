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
		INSERT INTO "Course" (title, description, created_by, create_date, organization_id, department_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING course_id
	`
	var courseID int
	err := r.db.QueryRow(
		query,
		course.Title,
		course.Description,
		course.CreatedBy,
		course.CreateDate,
		course.OrganizationID,
		course.DepartmentID, // 👈 добавлено
	).Scan(&courseID)
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
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date, c.organization_id, c.department_id,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies,
		       array_agg(DISTINCT m.material_id) AS material_ids
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

func (r *CourseRepository) GetCoursesByDepartment(orgID int, departmentID int) ([]models.CourseResponse, error) {
	var courses []models.CourseResponse
	query := `
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date, c.organization_id, c.department_id,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies,
		       array_agg(DISTINCT m.material_id) AS material_ids
		FROM "Course" c
		LEFT JOIN "CourseMaterial" cm ON c.course_id = cm.course_id
		LEFT JOIN "Material" m ON cm.material_id = m.material_id
		LEFT JOIN "CourseCompetency" cc ON c.course_id = cc.course_id
		LEFT JOIN "Competency" comp ON cc.competency_id = comp.competency_id
		WHERE c.department_id = $1 AND c.organization_id = $2
		GROUP BY c.course_id
	`
	err := r.db.Select(&courses, query, departmentID, orgID)
	if err != nil {
		return nil, fmt.Errorf("error fetching courses for department: %w", err)
	}
	return courses, nil
}

func (r *CourseRepository) GetAllCourses() ([]models.CourseResponse, error) {
	var courses []models.CourseResponse
	query := `
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date, c.organization_id,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies,
		       array_agg(DISTINCT m.material_id) AS material_ids
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

func (r *CourseRepository) GetCoursesByOrganization(orgID int) ([]models.CourseResponse, error) {
	var courses []models.CourseResponse
	query := `
		SELECT c.course_id, c.title, c.description, c.created_by, c.create_date, c.organization_id,
		       array_agg(DISTINCT m.title) AS materials,
		       array_agg(DISTINCT comp.name) AS competencies,
		       array_agg(DISTINCT m.material_id) AS material_ids
		FROM "Course" c
		LEFT JOIN "CourseMaterial" cm ON c.course_id = cm.course_id
		LEFT JOIN "Material" m ON cm.material_id = m.material_id
		LEFT JOIN "CourseCompetency" cc ON c.course_id = cc.course_id
		LEFT JOIN "Competency" comp ON cc.competency_id = comp.competency_id
		WHERE c.organization_id = $1
		GROUP BY c.course_id
	`
	err := r.db.Select(&courses, query, orgID)
	if err != nil {
		return nil, fmt.Errorf("error fetching courses for organization: %w", err)
	}
	return courses, nil
}

func (r *CourseRepository) UpdateCourse(course models.Course) error {
	tx, err := r.db.Beginx() // ← важно: BeginX, чтобы использовать tx.Select
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	var existingMaterialIDs []int
	err = tx.Select(&existingMaterialIDs, `
		SELECT material_id FROM "CourseMaterial" WHERE course_id = $1
	`, course.CourseID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch existing materials: %w", err)
	}

	materialsChanged := !equalIntSlices(existingMaterialIDs, course.Materials)

	query := `
	UPDATE "Course"
	SET title = $1, description = $2, created_by = $3, department_id = $4
	WHERE course_id = $5
`
	_, err = tx.Exec(query, course.Title, course.Description, course.CreatedBy, course.DepartmentID, course.CourseID)
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

	if materialsChanged {
		_, err := tx.Exec(`
			DELETE FROM "CourseProgress"
			WHERE course_id = $1 AND is_completed = TRUE
		`, course.CourseID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete completed progress: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit course update: %w", err)
	}

	return nil
}

func (r *CourseRepository) DeleteCourse(id int) error {
	query := `DELETE FROM "Course" WHERE course_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}

// GetUserCourseProgress returns the IDs of materials that the user has completed (viewed) in the given course.
func (r *CourseRepository) GetUserCourseProgress(userEmail string, courseID int) ([]int, error) {
	var materialIDs []int
	query := `
        SELECT material_id FROM "MaterialProgress"
        WHERE user_email = $1 AND course_id = $2 AND is_viewed = TRUE
    `
	err := r.db.Select(&materialIDs, query, userEmail, courseID)
	if err != nil {
		return nil, fmt.Errorf("error getting course progress: %w", err)
	}
	return materialIDs, nil
}

// MarkMaterialAsCompleted marks a material as viewed for a user in a course.
func (r *CourseRepository) MarkMaterialAsCompleted(userEmail string, courseID int, materialID int) error {
	query := `
        INSERT INTO "MaterialProgress" (user_email, course_id, material_id, is_viewed, viewed_at)
        VALUES ($1, $2, $3, TRUE, now())
        ON CONFLICT (user_email, course_id, material_id) DO NOTHING
    `
	_, err := r.db.Exec(query, userEmail, courseID, materialID)
	if err != nil {
		return fmt.Errorf("error marking material as viewed: %w", err)
	}
	return nil
}

// CompleteCourse marks the course as completed for a user.
func (r *CourseRepository) CompleteCourse(userEmail string, courseID int) error {
	query := `
        INSERT INTO "CourseProgress" (user_email, course_id, is_completed, completed_at)
        VALUES ($1, $2, TRUE, now())
        ON CONFLICT (user_email, course_id)
        DO UPDATE SET is_completed = TRUE, completed_at = now()
    `
	_, err := r.db.Exec(query, userEmail, courseID)
	if err != nil {
		return fmt.Errorf("error completing course: %w", err)
	}
	return nil
}
func (r *CourseRepository) IsCourseCompleted(userEmail string, courseID int) (bool, error) {
	var isCompleted bool
	query := `
	  SELECT is_completed FROM "CourseProgress"
	  WHERE user_email = $1 AND course_id = $2
	`
	err := r.db.Get(&isCompleted, query, userEmail, courseID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, fmt.Errorf("error checking course completion: %w", err)
	}
	return isCompleted, nil
}

// GetCompletedCourses retrieves the list of completed courses for a user.
func (r *CourseRepository) GetCompletedCourses(userEmail string) ([]models.CourseResponse, error) {
	var courses []models.CourseResponse
	query := `
        SELECT c.course_id, c.title, c.description, c.created_by, c.create_date, c.organization_id,
               array_agg(DISTINCT m.title) AS materials,
               array_agg(DISTINCT comp.name) AS competencies,
               array_agg(DISTINCT m.material_id) AS material_ids
        FROM "CourseProgress" cp
        JOIN "Course" c ON cp.course_id = c.course_id
        LEFT JOIN "CourseMaterial" cm ON c.course_id = cm.course_id
        LEFT JOIN "Material" m ON cm.material_id = m.material_id
        LEFT JOIN "CourseCompetency" cc ON c.course_id = cc.course_id
        LEFT JOIN "Competency" comp ON cc.competency_id = comp.competency_id
        WHERE cp.user_email = $1 AND cp.is_completed = TRUE
        GROUP BY c.course_id
    `
	err := r.db.Select(&courses, query, userEmail)
	if err != nil {
		return nil, fmt.Errorf("error fetching completed courses: %w", err)
	}
	return courses, nil
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	counter := make(map[int]int)
	for _, v := range a {
		counter[v]++
	}
	for _, v := range b {
		if counter[v] == 0 {
			return false
		}
		counter[v]--
	}
	for _, count := range counter {
		if count != 0 {
			return false
		}
	}
	return true
}

func (r *CourseRepository) GetCourseProgressByOrganization(orgID int) ([]models.UserCourseProgress, error) {
	var progress []models.UserCourseProgress
	query := `
		SELECT a.email AS user_email,
		       c.course_id,
		       c.title AS course_title,
		       COALESCE(cp.is_completed, false) AS is_completed,
		       cp.completed_at
		FROM "Account" a
		JOIN "Course" c ON c.organization_id = a.organization_id
		LEFT JOIN "CourseProgress" cp ON cp.course_id = c.course_id AND cp.user_email = a.email
		WHERE a.organization_id = $1
		ORDER BY a.email, c.course_id
	`
	err := r.db.Select(&progress, query, orgID)
	if err != nil {
		return nil, fmt.Errorf("error fetching course progress by organization: %w", err)
	}
	return progress, nil
}
