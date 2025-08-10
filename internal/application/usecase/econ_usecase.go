package usecase

import (
	"errors"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

// NOTE: Raw static data embedded here as a Go slice for simplicity.
// In real work you may place it in data/*.json and read from file.
var rawStaticData = []entities.RawYearData{
	{Year: 2019, DepositRate: 8.0, Inflation: 19.2, GDPGrowth: 6.1, M2Growth: 17.5, Unemployment: 19.1, ExchangeRate: 28.5},
	{Year: 2020, DepositRate: 8.0, Inflation: 20.5, GDPGrowth: 6.0, M2Growth: 15.0, Unemployment: 20.0, ExchangeRate: 35.0},
	{Year: 2021, DepositRate: 7.5, Inflation: 21.0, GDPGrowth: 7.6, M2Growth: 18.2, Unemployment: 19.5, ExchangeRate: 45.6},
	{Year: 2022, DepositRate: 8.0, Inflation: 30.0, GDPGrowth: 9.0, M2Growth: 23.0, Unemployment: 18.8, ExchangeRate: 50.2},
	{Year: 2023, DepositRate: 10.5, Inflation: 23.2, GDPGrowth: 3.8, M2Growth: 12.0, Unemployment: 17.9, ExchangeRate: 55.4},
	{Year: 2024, DepositRate: 12.0, Inflation: 21.7, GDPGrowth: 5.0, M2Growth: 10.5, Unemployment: 17.0, ExchangeRate: 60.0},
}

type EconUsecaseImpl struct {
	raw []entities.RawYearData
}

func NewEconUsecase() *EconUsecaseImpl {
	return &EconUsecaseImpl{raw: rawStaticData}
}

// helper to check range
func (u *EconUsecaseImpl) filterRange(fromYear, toYear int) ([]entities.RawYearData, error) {
	if fromYear == 0 && toYear == 0 {
		return u.raw, nil
	}
	if fromYear > toYear {
		return nil, errors.New("fromYear cannot be greater than toYear")
	}
	var out []entities.RawYearData
	for _, r := range u.raw {
		if r.Year >= fromYear && r.Year <= toYear {
			out = append(out, r)
		}
	}
	return out, nil
}

func (u *EconUsecaseImpl) GetAllIndicators() (entities.EconomicResponse, error) {
	return u.computeFrom(u.raw)
}

func (u *EconUsecaseImpl) GetIndicatorsForRange(fromYear, toYear int) (entities.EconomicResponse, error) {
	filtered, err := u.filterRange(fromYear, toYear)
	if err != nil {
		return entities.EconomicResponse{}, err
	}
	return u.computeFrom(filtered)
}

func (u *EconUsecaseImpl) computeFrom(raw []entities.RawYearData) (entities.EconomicResponse, error) {
	if len(raw) == 0 {
		return entities.EconomicResponse{}, errors.New("no data for requested range")
	}

	// Derived: Real Interest Rate = depositRate - inflation
	realMeta := entities.IndicatorMeta{
		Name:        "Real Interest Rate",
		Unit:        "percentage points",
		Description: "Calculated as depositRate - inflation (both in percent). Positive means nominal deposit > inflation.",
	}
	realValues := make([]entities.YearValue, 0, len(raw))
	for _, r := range raw {
		real := r.DepositRate - r.Inflation
		realValues = append(realValues, entities.YearValue{Year: r.Year, Value: round(real, 2)})
	}

	// Additional chosen indicators: GDP Growth, M2 Growth, Unemployment, Exchange Rate
	gdpMeta := entities.IndicatorMeta{
		Name:        "GDP Growth Rate",
		Unit:        "percent",
		Description: "Annual real GDP growth rate (percent). Source: NBE annual summary (simulated).",
	}
	m2Meta := entities.IndicatorMeta{
		Name:        "M2 Growth",
		Unit:        "percent",
		Description: "Year-over-year growth rate of M2 money supply (percent).",
	}
	unempMeta := entities.IndicatorMeta{
		Name:        "Unemployment Rate",
		Unit:        "percent",
		Description: "Share of labour force unemployed (percent).",
	}
	exchMeta := entities.IndicatorMeta{
		Name:        "Exchange Rate (ETB per USD)",
		Unit:        "ETB per USD",
		Description: "Annual average exchange rate in ETB per 1 USD.",
	}

	gdpVals := make([]entities.YearValue, 0, len(raw))
	m2Vals := make([]entities.YearValue, 0, len(raw))
	unempVals := make([]entities.YearValue, 0, len(raw))
	exchVals := make([]entities.YearValue, 0, len(raw))

	for _, r := range raw {
		gdpVals = append(gdpVals, entities.YearValue{Year: r.Year, Value: round(r.GDPGrowth, 2)})
		m2Vals = append(m2Vals, entities.YearValue{Year: r.Year, Value: round(r.M2Growth, 2)})
		unempVals = append(unempVals, entities.YearValue{Year: r.Year, Value: round(r.Unemployment, 2)})
		exchVals = append(exchVals, entities.YearValue{Year: r.Year, Value: round(r.ExchangeRate, 2)})
	}

	derived := []entities.DerivedIndicator{
		{Meta: realMeta, Values: realValues},
		{Meta: gdpMeta, Values: gdpVals},
		{Meta: m2Meta, Values: m2Vals},
		{Meta: unempMeta, Values: unempVals},
		{Meta: exchMeta, Values: exchVals},
	}

	explanation := "Real Interest Rate computed as depositRate - inflation (both in %). Other indicators are provided as annual values (2019-2024). Raw data is static and for demonstration; sources should be replaced with actual NBE report numbers."

	resp := entities.EconomicResponse{
		RawData:            raw,
		Derived:            derived,
		Explanation:        explanation,
		AvailableYearsFrom: raw[0].Year,
		AvailableYearsTo:   raw[len(raw)-1].Year,
	}

	return resp, nil
}

// small rounding helper
func round(v float64, decimals int) float64 {
	pow := 1.0
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int(v*pow+0.5)) / pow
}
