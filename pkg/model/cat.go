package model

import (
	"errors"
	"net/http"
)

type CatRequest struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Race   string  `json:"race"`
	Gender string  `json:"gender"`
	Weight float32 `json:"weight"`
}

type CatResponse struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Race   string  `json:"race"`
	Gender string  `json:"gender"`
	Weight float32 `json:"weight"`
}

func (cat *CatRequest) Bind(r *http.Request) error {
	if cat.Name == "" {
		return errors.New("name can't be null")
	}
	if cat.Age < 0 {
		return errors.New("age can't be negative")
	}
	if cat.Race == "" {
		return errors.New("name can't be null")
	}
	if cat.Gender == "" {
		return errors.New("gender can't be null")
	}
	if cat.Weight < 0 {
		return errors.New("weight can't be negative")
	}
	return nil
}
