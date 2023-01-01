package entity

import "github.com/kleber-yokota/go-cep/internal/entity/dto"

type Cdn struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
	Url        string `json:"url"`
}

func NewCdn(dto dto.Cdn, url string) Cdn {
	cdn := Cdn{
		Code:       dto.Code,
		State:      dto.State,
		City:       dto.City,
		District:   dto.District,
		Address:    dto.Address,
		Status:     dto.Status,
		Ok:         dto.Ok,
		StatusText: dto.StatusText,
		Url:        url,
	}
	return cdn
}
