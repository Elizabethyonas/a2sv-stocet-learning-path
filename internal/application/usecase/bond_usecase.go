package usecase

import (
	"time"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

type BondUsecaseImpl struct {
	mockBonds []entities.Bond
}

func NewBondUsecase() *BondUsecaseImpl {
	// Hardcoded bond data
	mockData := []entities.Bond{
		{ID: "1", Name: "Global Bond A", Issuer: "Gov A", CouponRate: 5.2, MaturityDate: time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC), Price: 1000},
		{ID: "2", Name: "Corporate Bond B", Issuer: "Corp B", CouponRate: 6.5, MaturityDate: time.Date(2027, 10, 15, 0, 0, 0, 0, time.UTC), Price: 980},
		{ID: "3", Name: "Municipal Bond C", Issuer: "City C", CouponRate: 4.0, MaturityDate: time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC), Price: 950},
		// Add 5–7 more...
	}
	return &BondUsecaseImpl{mockBonds: mockData}
}

func (u *BondUsecaseImpl) SearchBonds(minCoupon, maxCoupon float64, maturityAfter, maturityBefore time.Time) []entities.Bond {
	var result []entities.Bond
	for _, bond := range u.mockBonds {
		if (minCoupon == 0 || bond.CouponRate >= minCoupon) &&
			(maxCoupon == 0 || bond.CouponRate <= maxCoupon) &&
			(maturityAfter.IsZero() || bond.MaturityDate.After(maturityAfter)) &&
			(maturityBefore.IsZero() || bond.MaturityDate.Before(maturityBefore)) {
			result = append(result, bond)
		}
	}
	return result
}
