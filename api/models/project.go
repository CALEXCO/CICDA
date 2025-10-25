package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Project struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Leader         string      `json:"leader"`
	Description    string      `json:"description"`
	DateOfCreation pgtype.Date `json:"dateOfCreation"`
}

func (p Project) GetName() string {
	return p.Name
}

func (p Project) GetLeader() string {
	return p.Leader
}

func (p Project) GetDescription() string {
	return p.Description
}

func (p Project) GetDateOfCreation() pgtype.Date {
	return p.DateOfCreation
}
