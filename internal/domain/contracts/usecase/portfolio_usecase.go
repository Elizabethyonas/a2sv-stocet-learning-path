package usecase

import "a2sv_stocet_learning_path/internal/domain/entities"

// PortfolioUsecase defines the interface for portfolio recommendations
type PortfolioUsecase interface {
	RecommendPortfolio(req entities.RecommendationRequest) (entities.RecommendationResponse, error)
}
