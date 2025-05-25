package models

import (
	"github.com/lib/pq"
	"time"
)

type Course struct {
	CourseID       int       `json:"course_id" example:"1"`
	Title          string    `json:"title" example:"Основы Go"`
	Description    string    `json:"description" example:"Курс по основам языка Go"`
	CreatedBy      string    `json:"created_by" example:"admin@test.ru"`
	Materials      []int     `json:"materials" example:"[1, 2, 3]"`
	Competencies   []int     `json:"competencies" example:"[1, 5]"`
	CreateDate     time.Time `json:"create_date" example:"2024-11-28T15:04:05Z"`
	OrganizationID int       `json:"organization_id"`
}

type CreateCourseRequest struct {
	Title          string `json:"title" example:"Основы Go"`
	Description    string `json:"description" example:"Курс по Go для начинающих"`
	CreatedBy      string `json:"created_by" example:"admin@test.ru"`
	Materials      []int  `json:"materials" example:"[1,2]"`
	Competencies   []int  `json:"competencies" example:"[3,5]"`
	OrganizationID int    `json:"organization_id"`
}

type CourseResponse struct {
	CourseID       int            `db:"course_id" json:"course_id" example:"1"`
	Title          string         `db:"title" json:"title" example:"Основы Go"`
	Description    string         `db:"description" json:"description" example:"Курс по Go для начинающих"`
	CreatedBy      string         `db:"created_by" json:"created_by" example:"admin@test.ru"`
	Materials      pq.StringArray `db:"materials" json:"materials" swaggertype:"array,string" example:"[\"Материал 1\", \"Материал 2\"]"`
	Competencies   pq.StringArray `db:"competencies" json:"competencies" swaggertype:"array,string" example:"[\"Go Basics\", \"Concurrency\"]"`
	CreateDate     time.Time      `db:"create_date" json:"create_date" example:"2024-11-28T15:04:05Z"`
	OrganizationID int            `db:"organization_id" json:"organization_id"`
}

type CourseMaterial struct {
	CourseID   int `db:"course_id"`
	MaterialID int `db:"material_id"`
}

type CourseCompetency struct {
	CourseID     int `db:"course_id"`
	CompetencyID int `db:"competency_id"`
}
