package usecase

import "a2sv_stocet_learning_path/internal/domain/entities"

type EquityUsecase interface {
	GetEquities(filterType string) []entities.Equity
}
