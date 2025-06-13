package models

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type Course struct {
	CourseID       int       `json:"course_id" example:"1"`
	Title          string    `json:"title" example:"Основы Go"`
	Description    string    `json:"description" example:"Курс по основам языка Go"`
	CreatedBy      *string   `json:"created_by" example:"admin@test.ru"`
	Materials      []int     `json:"materials" example:"[1, 2, 3]"`
	Competencies   []int     `json:"competencies" example:"[1, 5]"`
	CreateDate     time.Time `json:"create_date" example:"2024-11-28T15:04:05Z"`
	OrganizationID int       `json:"organization_id"`
	DepartmentID   int       `json:"department_id"` // 👈 добавлено
}

type CreateCourseRequest struct {
	Title          string `json:"title" example:"Основы Go"`
	Description    string `json:"description" example:"Курс по Go для начинающих"`
	CreatedBy      string `json:"created_by" example:"admin@test.ru"`
	Materials      []int  `json:"materials" example:"[1,2]"`
	Competencies   []int  `json:"competencies" example:"[3,5]"`
	OrganizationID int    `json:"organization_id"`
	DepartmentID   int    `json:"department_id"` // 👈 добавлено
}

type CourseResponse struct {
	CourseID       int            `db:"course_id" json:"course_id" example:"1"`
	Title          string         `db:"title" json:"title" example:"Основы Go"`
	Description    string         `db:"description" json:"description" example:"Курс по Go для начинающих"`
	CreatedBy      *string        `db:"created_by" json:"created_by" example:"admin@test.ru"`
	Materials      pq.StringArray `db:"materials" json:"materials" swaggertype:"array,string" example:"[\"Материал 1\", \"Материал 2\"]"`
	MaterialIDs    pq.Int64Array  `db:"material_ids" json:"material_ids" swaggertype:"array,integer"`
	Competencies   pq.StringArray `db:"competencies" json:"competencies" swaggertype:"array,string" example:"[\"Go Basics\", \"Concurrency\"]"`
	CreateDate     time.Time      `db:"create_date" json:"create_date" example:"2024-11-28T15:04:05Z"`
	OrganizationID int            `db:"organization_id" json:"organization_id"`
	DepartmentID   int            `db:"department_id" json:"department_id"` // 👈 добавлено
}

type CourseMaterial struct {
	CourseID   int `db:"course_id"`
	MaterialID int `db:"material_id"`
}

type CourseCompetency struct {
	CourseID     int `db:"course_id"`
	CompetencyID int `db:"competency_id"`
}

type MaterialProgress struct {
	UserEmail  string    `db:"user_email" json:"user_email"`
	CourseID   int       `db:"course_id" json:"course_id"`
	MaterialID int       `db:"material_id" json:"material_id"`
	ViewedAt   time.Time `db:"viewed_at" json:"viewed_at"`
}

type CourseProgress struct {
	ProgressID  int       `db:"progress_id" json:"progress_id"`
	UserEmail   string    `db:"user_email" json:"user_email"`
	CourseID    int       `db:"course_id" json:"course_id"`
	IsCompleted bool      `db:"is_completed" json:"is_completed"`
	CompletedAt time.Time `db:"completed_at" json:"completed_at"`
}

type CourseProgressResponse struct {
	CompletedMaterials []int `json:"completed_materials" swaggertype:"array,integer"`
}

type UserCourseProgress struct {
	UserEmail       string       `db:"user_email" json:"user_email"`
	UserName        string       `db:"user_name" json:"user_name"`
	CourseID        int          `db:"course_id" json:"course_id"`
	CourseTitle     string       `db:"course_title" json:"course_title"`
	IsCompleted     sql.NullBool `db:"is_completed" json:"-"`
	CompletedAt     *time.Time   `db:"completed_at" json:"completed_at,omitempty"`
	TotalMaterials  int          `db:"total_materials" json:"total_materials"`
	ViewedMaterials int          `db:"viewed_materials" json:"viewed_materials"`
	LastViewedAt    *time.Time   `db:"last_viewed_at" json:"last_viewed_at,omitempty"`
}
