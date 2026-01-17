package darwin

func SumByPlates(data []DarwinAPITrechosResponse) (map[string]float64, error) {
	response := make(map[string]float64)

	for _, record := range data {
		response[record.Placa] += record.KmPercorrindos
	}

	return response, nil
}
