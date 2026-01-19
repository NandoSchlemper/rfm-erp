package darwin

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
)

type service struct{}

type Service interface {
	GetVehiclesKM(DarwinAPITrechosPayload) ([]DarwinTrechosResponse, error)
}

func (s service) GetVehiclesKM(body DarwinAPITrechosPayload) ([]DarwinTrechosResponse, error) {
	apiConfig := &DarwinAPIScheme{}
	token, err := apiConfig.Load()
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar as variaveis da API.\n%v", err)
	}

	fmt.Printf("Bearer token: %v", token)

	urlParams := url.Values{}
	urlParams.Add("token", *token)
	urlParams.Add("data_inicial", body.FirstDate)
	urlParams.Add("data_final", body.LastDate)

	url := apiConfig.url + "trechos" + "?" + urlParams.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+*token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao realizar a requisição: \n%v\n", err)
	}
	defer resp.Body.Close()

	var response []DarwinAPITrechosResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("Erro ao descompactar o JSON: \n%v\n", err.Error())
	}

	process, err := ProcessData(response)
	if err != nil {
		return nil, fmt.Errorf("Erro ao transformar os dados recebidos. \n%v\n", err)
	}

	fmt.Printf("Dados da API de Trechos (Darwin) formatados: %v", process)

	return process, nil
}

func NewService() Service {
	return &service{}
}
