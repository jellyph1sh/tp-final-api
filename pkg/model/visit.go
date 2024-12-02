package model

import (
	"errors"
	"net/http"
)

type VisitRequest struct {
	CatID  int    `json:"cat_id"`
	Date   string `json:"date"`
	Reason string `json:"reason"`
	Doctor string `json:"doctor"`
}

type VisitResponse struct {
	ID     int    `json:"id"`
	CatID  int    `json:"cat_id"`
	Date   string `json:"date"`
	Reason string `json:"reason"`
	Doctor string `json:"doctor"`
}

type VisitHistoryResponse struct {
	ID         int                  `json:"id"`
	CatID      int                  `json:"cat_id"`
	Date       string               `json:"date"`
	Reason     string               `json:"reason"`
	Doctor     string               `json:"doctor"`
	Treatments []*TreatmentResponse `json:"treatments"`
}

func (visit *VisitRequest) Bind(r *http.Request) error {
	if visit.CatID == 0 {
		return errors.New("cat_id can't be null")
	}
	if visit.Date == "" {
		return errors.New("date can't be null")
	}
	if visit.Reason == "" {
		return errors.New("reason can't be null")
	}
	if visit.Doctor == "" {
		return errors.New("doctor can't be null")
	}
	return nil
}
