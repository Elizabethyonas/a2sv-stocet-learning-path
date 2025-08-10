package entities

// Raw yearly macro data for a single year
type RawYearData struct {
	Year         int     `json:"year"`
	DepositRate  float64 `json:"depositRate"`  // percent, e.g., 8.5 means 8.5%
	Inflation    float64 `json:"inflation"`    // percent (CPI)
	GDPGrowth    float64 `json:"gdpGrowth"`    // percent
	M2Growth     float64 `json:"m2Growth"`     // percent
	Unemployment float64 `json:"unemployment"` // percent
	ExchangeRate float64 `json:"exchangeRate"` // ETB per USD (annual average)
}

// Metadata for an indicator
type IndicatorMeta struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}

// A single year/value pair for an indicator
type YearValue struct {
	Year  int     `json:"year"`
	Value float64 `json:"value"`
}

// Derived indicator structure
type DerivedIndicator struct {
	Meta   IndicatorMeta `json:"meta"`
	Values []YearValue   `json:"values"`
}

// Full response structure
type EconomicResponse struct {
	RawData            []RawYearData      `json:"rawData"`
	Derived            []DerivedIndicator `json:"derived"`
	Explanation        string             `json:"explanation"`
	AvailableYearsFrom int                `json:"availableYearsFrom"`
	AvailableYearsTo   int                `json:"availableYearsTo"`
}
