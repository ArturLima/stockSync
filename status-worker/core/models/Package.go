package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Package struct {
	bun.BaseModel `bun:"table:status,alias:st"`

	Id     uuid.UUID `json:"id" bun:"id,pk"`
	Status string    `json:"status" bun:"status,pk"`
	Motivo string    `json:"motivo" bun:"motivo,pk"`
}
