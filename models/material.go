package models

import (
	"github.com/lib/pq"
	"time"
)

type Material struct {
	MaterialID   int       `json:"material_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Type         int       `json:"type_id"`
	Content      string    `json:"content"`
	Competencies []int     `json:"competencies"`
	CreateDate   time.Time `json:"create_date"`
}

type CreateMaterialRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	TypeID       int    `json:"type_id"`
	Content      string `json:"content"`
	Competencies []int  `json:"competencies"`
}

type MaterialResponse struct {
	MaterialID   int            `db:"material_id" json:"material_id"`
	Title        string         `db:"title" json:"title"`
	Description  string         `db:"description" json:"description"`
	TypeName     string         `db:"type_name" json:"type_name"`
	Content      string         `db:"content" json:"content"`
	Competencies pq.StringArray `db:"competencies" json:"competencies"`
	CreateDate   string         `db:"create_date" json:"create_date"`
}

type MaterialType struct {
	TypeID int    `db:"type_id"`
	Type   string `db:"type"`
}
