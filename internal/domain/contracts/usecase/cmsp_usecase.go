package usecase

import (
	"time"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

type CMSPUsecase interface {
	GetAll(typeFilter string, before, after time.Time) []entities.CMSP
	GetByID(id string) (*entities.CMSP, error)
}
