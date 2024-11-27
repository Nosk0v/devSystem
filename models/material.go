package models

import "time"

type Material struct {
	MaterialID    int       `db:"material_id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	Type          int       `json:"type" db:"type"`
	Content       string    `json:"content" db:"content"`
	CompetencyIDs []int     `json:"competencyIDs" db:"competency_ids"`
	CreateDate    time.Time `json:"createDate" db:"create_date"`
}

type MaterialType struct {
	TypeID int    `db:"type_id"`
	Type   string `db:"type"`
}
