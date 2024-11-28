package models

import "time"

type Competency struct {
	CompetencyID int       `db:"competency_id" json:"competency_id" example:"3"`
	Name         string    `db:"name" json:"name" example:"Основы ООП"`
	Description  string    `db:"description" json:"description" example:"Понимание принципов объектно-ориентированного программирования"`
	ParentID     *int      `json:"parent_id" db:"parent_id" example:"1"`
	CreateDate   time.Time `db:"create_date" json:"create_date" example:"2024-11-28T15:04:05Z"`
}
