package usecase

import (
	"time"
	"a2sv_stocet_learning_path/internal/domain/entities"
)

type BondUsecase interface {
	SearchBonds(minCoupon, maxCoupon float64, maturityAfter, maturityBefore time.Time) []entities.Bond
}
