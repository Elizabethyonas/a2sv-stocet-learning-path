package usecase

import (
	"strings"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

type EquityUsecaseImpl struct {
	mockData []entities.Equity
}

func NewEquityUsecase() *EquityUsecaseImpl {
	data := []entities.Equity{
		{Symbol: "ETHB", Name: "EthioBank", Price: 120.5, ChangePercent: 2.5, Volume: 5000},
		{Symbol: "ETHC", Name: "EthioCorp", Price: 85.3, ChangePercent: -1.2, Volume: 12000},
		{Symbol: "ETHG", Name: "EthioGold", Price: 200.1, ChangePercent: 3.1, Volume: 3000},
		{Symbol: "ETHF", Name: "EthioFoods", Price: 45.0, ChangePercent: -0.8, Volume: 8000},
		{Symbol: "ETHP", Name: "EthioPower", Price: 150.0, ChangePercent: 1.0, Volume: 10000},
	}
	return &EquityUsecaseImpl{mockData: data}
}

func (u *EquityUsecaseImpl) GetEquities(filterType string) []entities.Equity {
	filterType = strings.ToLower(filterType)
	var result []entities.Equity

	for _, equity := range u.mockData {
		if filterType == "gainer" && equity.ChangePercent <= 0 {
			continue
		}
		if filterType == "loser" && equity.ChangePercent >= 0 {
			continue
		}
		result = append(result, equity)
	}
	return result
}
