package entities

type InvestmentProfile string

const (
	Active  InvestmentProfile = "active"
	Passive InvestmentProfile = "passive"
)

type RecommendationRequest struct {
	Profile       InvestmentProfile `json:"profile"`
	InitialCapital float64          `json:"initialCapital"`
	TargetGoal     float64          `json:"targetGoal"`
}

type Asset struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Yield  float64 `json:"yield"`
}

type RecommendationResponse struct {
	Stock      Asset  `json:"stock"`
	Bond       Asset  `json:"bond"`
	Cash       Asset  `json:"cash"`
	TotalReturn float64 `json:"totalReturn"`
	GoalMet     bool    `json:"goalMet"`
	GapToGoal   float64 `json:"gapToGoal"`
}
