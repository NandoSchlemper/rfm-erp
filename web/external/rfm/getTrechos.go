package darwin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/NandoSchlemper/rfm-erp/frontend/models"
)

func GetTrechosData() ([]models.DarwinTrechosResponse, error) {
	// layout := "2000-12-31"
	client := http.Client{}
	url := "http://host.docker.internal:3000/darwin/trechos"

	// now_timer := time.Now()

	initial_date := "2026-01-19 00:00"
	final_date := "2026-01-19 23:59"

	requestBody, err := json.Marshal(models.DarwinTrechosRequest{
		Initial_date: initial_date,
		Final_date:   final_date,
	})
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar o body da requisição. %v\n", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Erro ao montar a request através do cliente %v\n", err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Erro ao receber a resposta dos dados. %v\n", err)
	}
	defer response.Body.Close()

	var returnValue []models.DarwinTrechosResponse
	err = json.NewDecoder(response.Body).Decode(&returnValue)
	if err != nil {
		return nil, fmt.Errorf("Erro ao descompactar a resposta para um JSON. %v\n", err)
	}

	return returnValue, nil
}

func SortPlacasByPriorityAndKM(data []models.DarwinTrechosResponse, priorityPlacas []string) []models.DarwinTrechosResponse {
	priorityMap := make(map[string]bool)
	for _, placa := range priorityPlacas {
		placaSemHifen := strings.ReplaceAll(placa, "-", "")
		priorityMap[placaSemHifen] = true
	}

	var priority, nonPriority []models.DarwinTrechosResponse
	for _, item := range data {
		if priorityMap[item.Placa] {
			priority = append(priority, item)
		} else {
			nonPriority = append(nonPriority, item)
		}
	}

	sort.Slice(priority, func(i, j int) bool {
		return priority[i].KmPercorridos > priority[j].KmPercorridos
	})

	sort.Slice(nonPriority, func(i, j int) bool {
		return nonPriority[i].KmPercorridos > nonPriority[j].KmPercorridos
	})

	return append(priority, nonPriority...)
}
