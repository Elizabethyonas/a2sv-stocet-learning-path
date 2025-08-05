package usecase

import (
	"errors"
	"strings"
	"time"

	"a2sv_stocet_learning_path/internal/domain/entities"
)

type CMSPUsecaseImpl struct {
	data []entities.CMSP
}

func NewCMSPUsecase() *CMSPUsecaseImpl {
	// Sample mock CMSP data
	cmsps := []entities.CMSP{
		{
			ID:           "1",
			Name:         "EthioCMSP",
			Type:         "Telecom",
			LicensedDate: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			Description:  "Licensed for telecom infrastructure services.",
			Plan:         "5G rollout by 2026",
			Source:       "https://ethiocommunications.gov.et",
		},
		{
			ID:           "2",
			Name:         "Safaricom Ethiopia",
			Type:         "Mobile",
			LicensedDate: time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
			Description:  "Kenyan-based mobile service provider.",
			Plan:         "Nationwide expansion",
			Source:       "https://safaricom.et",
		},
	}
	return &CMSPUsecaseImpl{data: cmsps}
}

func (u *CMSPUsecaseImpl) GetAll(typeFilter string, before, after time.Time) []entities.CMSP {
	var result []entities.CMSP
	for _, c := range u.data {
		if typeFilter != "" && !strings.EqualFold(c.Type, typeFilter) {
			continue
		}
		if !before.IsZero() && c.LicensedDate.After(before) {
			continue
		}
		if !after.IsZero() && c.LicensedDate.Before(after) {
			continue
		}
		result = append(result, c)
	}
	return result
}

func (u *CMSPUsecaseImpl) GetByID(id string) (*entities.CMSP, error) {
	for _, c := range u.data {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("CMSP not found")
}
