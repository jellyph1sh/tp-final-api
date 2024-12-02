package model

import (
	"errors"
	"net/http"
)

type TreatmentRequest struct {
	VisitID   int    `json:"visit_id"`
	Medicine  string `json:"medicine"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	DoctorTip string `json:"doctor_tip"`
}

type TreatmentResponse struct {
	ID        int    `json:"id"`
	VisitID   int    `json:"visit_id"`
	Medicine  string `json:"medicine"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	DoctorTip string `json:"doctor_tip"`
}

func (treatment *TreatmentRequest) Bind(r *http.Request) error {
	if treatment.Medicine == "" {
		return errors.New("medicine can't be null")
	}
	if treatment.StartDate == "" {
		return errors.New("start_date can't be null")
	}
	if treatment.EndDate == "" {
		return errors.New("end_date can't be null")
	}
	if treatment.DoctorTip == "" {
		return errors.New("doctor_tip can't be null")
	}
	return nil
}
