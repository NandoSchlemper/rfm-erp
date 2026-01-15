package darwin

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

func GetTrechosPercorridos(initial_date, final_date string) (*string, error) {
	apiConfig := &DarwinAPIScheme{}
	apiConfig.Load()

	layout := "2026-01-15 00:00"

	initial_date_edited, _ := time.Parse(layout, initial_date)
	final_date_edited, _ := time.Parse(layout, final_date)

	url := apiConfig.url + "trechos"
	payload, _ := json.Marshal(&DarwinAPITrechosPayload{
		FirstDate: initial_date_edited.String(),
		LastDate:  final_date_edited.String(),
	})

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiConfig.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao realizar a requisição: %v", err.Error())
	}
	defer resp.Body.Close()

	// var response []DarwinAPITrechosResponse
	// err = json.NewDecoder(resp.Body).Decode(&response)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao descompactar o JSON: %v", err.Error())
	}

	stringfyData := string(data)

	return &stringfyData, nil
}
