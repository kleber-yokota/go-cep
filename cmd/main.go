package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kleber-yokota/go-cep/internal/entity"
	"github.com/kleber-yokota/go-cep/internal/entity/dto"
)

func main() {
	cep := os.Args[1]

	chViaCep := make(chan entity.ViaCep)
	chCdn := make(chan entity.Cdn)

	go ViaCepColletion(cep, chViaCep)
	go CndColletion(cep, chCdn)

	select {
	case result := <-chViaCep:
		fmt.Printf("%+v\n", result)
	case result := <-chCdn:
		fmt.Printf("%+v\n", result)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}

}

func ViaCepColletion(cep string, ch chan entity.ViaCep) {
	var viaCep dto.ViaCep

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	payload := r.Body
	json.NewDecoder(payload).Decode(&viaCep)
	ch <- entity.NewViaCep(viaCep, url)

}

func CndColletion(cep string, ch chan entity.Cdn) {
	var cdn dto.Cdn

	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	payload := r.Body
	json.NewDecoder(payload).Decode(&cdn)
	ch <- entity.NewCdn(cdn, url)

}
