package models

import "time"

type Material struct {
	MaterialID  int       `db:"material_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Type        int       `db:"type"`
	Content     string    `db:"content"`
	CreateDate  time.Time `db:"create_date"`
}

type MaterialType struct {
	TypeID int    `db:"type_id"`
	Type   string `db:"type"`
}
