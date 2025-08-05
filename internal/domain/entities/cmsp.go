package entities

import "time"

type CMSP struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	LicensedDate time.Time `json:"licensedDate"`
	Description  string    `json:"description"`
	Plan         string    `json:"plan"`
	Source       string    `json:"source"`
}
