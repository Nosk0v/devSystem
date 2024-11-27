package models

import "time"

type Competency struct {
	CompetencyID int       `db:"competency_id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	ParentID     *int      `json:"parent_id"  db:"parent_id"`
	CreateDate   time.Time `db:"create_date"`
}
