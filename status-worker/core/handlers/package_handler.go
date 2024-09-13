package handlers

import (
	"encoding/json"
	"log"

	"github.com/Arturlima/status-worker/core/models"
	"github.com/Arturlima/status-worker/infra/db"
)

type IPackageHandler interface {
	AddOrUpdateStatus(b []byte) (done bool)
}

type PackageHandler struct {
	repo db.IStatusRepository
}

func NewPackageHandler() IPackageHandler {
	return &PackageHandler{
		repo: db.NewRepository(),
	}
}

func (p *PackageHandler) AddOrUpdateStatus(b []byte) (done bool) {

	var model models.Package

	err := json.Unmarshal(b, &model)
	if err != nil {
		log.Println("Error > ", err)
	}

	err = p.repo.Insert(&model)
	return err != nil
}
