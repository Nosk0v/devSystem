package models

import "time"

type Material struct {
	MaterialID  int    `db:"material_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Type        int    `db:"type"`
	Content     string `db:"content"`
	CreateDate  string `db:"create_date"`
}

type MaterialType struct {
	TypeID int    `db:"type_id"`
	Type   string `db:"type"`
}

type Competency struct {
	CompetencyID int       `db:"competency_id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	ParentID     *int      `db:"parent_id"` // ParentID может быть NULL
	CreateDate   time.Time `db:"create_date"`
}

type User struct {
	Username string `db:"username"`
}

type MaterialCompetency struct {
	MaterialID   int `db:"material_id"`
	CompetencyID int `db:"competency_id"`
}
