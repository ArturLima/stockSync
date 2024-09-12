package handlers

import (
	"github.com/Arturlima/status-worker/infra/db"
	"github.com/uptrace/bun"
)

type IPackageHandler interface {
	AddOrUpdateStatus(b []byte) (done bool)
}

type PackageHandler struct {
	db *bun.DB
}

func NewPackageHandler() IPackageHandler {
	return &PackageHandler{
		db: db.NewDB(),
	}
}

func (p *PackageHandler) AddOrUpdateStatus(b []byte) (done bool) {
	return true
}
