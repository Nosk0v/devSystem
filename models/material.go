package models

import (
	"github.com/lib/pq"
	"time"
)

type Material struct {
	MaterialID   int       `json:"material_id" example:"1"`
	Title        string    `json:"title" example:"Введение в Go"`
	Description  string    `json:"description" example:"Руководство для начинающих по программированию на языке Go"`
	Type         int       `json:"type_id" example:"2"`
	Content      string    `json:"content" example:"Go - статически типизированный..."`
	Competencies []int     `json:"competencies" example:"[1, 3, 5]"`
	CreateDate   time.Time `json:"create_date" example:"2024-11-28T15:04:05Z"`
}

type CreateMaterialRequest struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	TypeID       int    `json:"type_id"`
	Content      string `json:"content"`
	Competencies []int  `json:"competencies"`
}

type MaterialResponse struct {
	MaterialID   int            `db:"material_id" json:"material_id" example:"1"`
	Title        string         `db:"title" json:"title" example:"Введение в Go"`
	Description  string         `db:"description" json:"description" example:"Руководство для начинающих по программированию на языке Go"`
	TypeName     string         `db:"type_name" json:"type_name" example:"Статья"`
	Content      string         `db:"content" json:"content" example:"Go - статически типизированный..."`
	Competencies pq.StringArray `db:"competencies" json:"competencies" swaggertype:"array,string" example:"[\"Основы ООП\", \"Знание синтаксиса\"]"`
	CreateDate   string         `db:"create_date" json:"create_date" example:"2024-11-28T15:04:05Z"`
}

type MaterialType struct {
	TypeID int    `db:"type_id"`
	Type   string `db:"type"`
}
