package usecase

import (
	"errors"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

type PortfolioUsecaseImpl struct{}

func NewPortfolioUsecase() *PortfolioUsecaseImpl {
	return &PortfolioUsecaseImpl{}
}

func (u *PortfolioUsecaseImpl) RecommendPortfolio(req entities.RecommendationRequest) (entities.RecommendationResponse, error) {
	if req.InitialCapital <= 0 || req.TargetGoal <= 0 {
		return entities.RecommendationResponse{}, errors.New("invalid input")
	}

	var stockWeight, bondWeight, cashWeight float64

	switch req.Profile {
	case entities.Active:
		stockWeight = 0.6
		bondWeight = 0.3
		cashWeight = 0.1
	case entities.Passive:
		stockWeight = 0.3
		bondWeight = 0.4
		cashWeight = 0.3
	default:
		return entities.RecommendationResponse{}, errors.New("invalid profile")
	}

	stockAmount := req.InitialCapital * stockWeight
	bondAmount := req.InitialCapital * bondWeight
	cashAmount := req.InitialCapital * cashWeight

	stockYield := stockAmount * 0.20 // 20%
	bondYield := bondAmount * 0.10   // 10%
	cashYield := cashAmount * 0.07   // 7%

	totalReturn := stockYield + bondYield + cashYield
	gap := totalReturn - (req.TargetGoal - req.InitialCapital)
	goalMet := totalReturn+req.InitialCapital >= req.TargetGoal

	return entities.RecommendationResponse{
		Stock: entities.Asset{
			Name:   "EthioTelecom",
			Amount: stockAmount,
			Yield:  stockYield,
		},
		Bond: entities.Asset{
			Name:   "Ethiopian Gov Bond",
			Amount: bondAmount,
			Yield:  bondYield,
		},
		Cash: entities.Asset{
			Name:   "Birr",
			Amount: cashAmount,
			Yield:  cashYield,
		},
		TotalReturn: totalReturn,
		GoalMet:     goalMet,
		GapToGoal:   gap,
	}, nil
}
