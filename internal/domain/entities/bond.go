package entities

import "time"

type Bond struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Issuer      string    `json:"issuer"`
	CouponRate  float64   `json:"couponRate"`
	MaturityDate time.Time `json:"maturityDate"`
	Price       float64   `json:"price"`
}
