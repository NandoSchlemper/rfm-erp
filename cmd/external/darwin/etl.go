package darwin

func ProcessData(data []DarwinAPITrechosResponse) ([]DarwinTrechosResponse, error) {
	response := make(map[string]DarwinTrechosResponse)

	for _, record := range data {
		if r, ok := response[record.Placa]; ok {
			r.KmPercorridos += record.KmPercorrindos
			response[record.Placa] = r
		} else {
			response[record.Placa] = DarwinTrechosResponse{
				Placa:         record.Placa,
				KmPercorridos: record.KmPercorrindos,
			}
		}
	}
	result := make([]DarwinTrechosResponse, 0, len(response))
	for _, v := range response {
		result = append(result, v)
	}
	return result, nil
}
