package usecase

import "a2sv_stocet_learning_path/internal/domain/entities"

type EconUsecase interface {
	// Return indicators for all years (2019-2024).
	GetAllIndicators() (entities.EconomicResponse, error)

	// Optionally filter by year range
	GetIndicatorsForRange(fromYear, toYear int) (entities.EconomicResponse, error)
}
