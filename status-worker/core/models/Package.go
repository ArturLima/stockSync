package models

import "github.com/uptrace/bun"

type Package struct {
	bun.BaseModel `bun:"table:PACKAGE,alias:CB"`
}
