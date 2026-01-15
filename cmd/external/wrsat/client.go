package wrsat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

func GetActualPositions() (*WrsatListVehiclesResponse, error) {
	apiConfig := WrsatAPIScheme{}
	apiConfig.Load()

	apiPayload, _ := json.Marshal(WrsatAPIPayload{
		User:        apiConfig.user,
		Password:    apiConfig.password,
		Order:       "ASC",
		Limit:       "100",
		Page:        "1",
		Description: "",
	})

	req, err := http.NewRequest("POST", apiConfig.url+"/lista_veiculos", bytes.NewBuffer(apiPayload))
	if err != nil {
		return nil, fmt.Errorf("Erro ao montar a requisição: %v\n", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiConfig.user+":"+apiConfig.password)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao requisitar a API: %v\n", err.Error())
	}
	defer resp.Body.Close()

	var actualPositions WrsatListVehiclesResponse
	err = json.NewDecoder(resp.Body).Decode(&actualPositions)
	if err != nil {
		return nil, fmt.Errorf("Erro ao descompactar o JSON: %v\n", err.Error())
	}
	return &actualPositions, nil
}
