package entity

import "github.com/kleber-yokota/go-cep/internal/entity/dto"

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Url         string `json:"url"`
}


func NewViaCep(dto dto.ViaCep, url string) ViaCep{
	viaCep := ViaCep{
		Cep: dto.Cep,
		Logradouro: dto.Logradouro,
		Complemento: dto.Complemento,
		Bairro: dto.Bairro,
		Localidade: dto.Localidade,
		Uf: dto.Uf,
		Ibge: dto.Ibge,
		Gia: dto.Gia,
		Ddd: dto.Ddd,
		Siafi: dto.Siafi,
		Url: url,
	}
	return viaCep	
}