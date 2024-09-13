package db

import (
	"context"

	"github.com/Arturlima/status-worker/core/models"
	"github.com/uptrace/bun"
)

type IStatusRepository interface {
	Insert(product *models.Package) (err error)
}

type StatusRepository struct {
	context *bun.DB
}

func NewRepository() IStatusRepository {
	return &StatusRepository{
		context: NewDB(),
	}
}

func (s *StatusRepository) Insert(product *models.Package) (err error) {

	err = s.context.
		NewInsert().
		Model(product).
		Scan(context.Background())

	return
}
