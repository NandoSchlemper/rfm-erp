package darwin

import (
	"sort"
	"time"
)

func ProcessData(data []DarwinAPITrechosResponse) ([]DarwinTrechosResponse, error) {
	placasMap := make(map[string]*DarwinTrechosResponse)

	viagensPorPlaca := make(map[string][]DarwinAPITrechosResponse)

	for _, record := range data {
		viagensPorPlaca[record.Placa] = append(viagensPorPlaca[record.Placa], record)
	}

	for placa, viagens := range viagensPorPlaca {
		sort.Slice(viagens, func(i, j int) bool {
			dateI, _ := time.Parse("2006-01-02 15:04:05.000000", viagens[i].DataInicial["date"].(string))
			dateJ, _ := time.Parse("2006-01-02 15:04:05.000000", viagens[j].DataInicial["date"].(string))
			return dateI.Before(dateJ)
		})

		var totalKm float64
		var totalTempo int

		for _, v := range viagens {
			totalKm += v.KmPercorrido
			totalTempo += v.Tempo
		}

		primeiraViagem := viagens[0]

		ultimaViagem := viagens[len(viagens)-1]

		primeiraDataStr := primeiraViagem.DataInicial["date"].(string)
		primeiraData, _ := time.Parse("2006-01-02 15:04:05.000000", primeiraDataStr)

		ultimaDataStr := ultimaViagem.DataFinal["date"].(string)
		ultimaData, _ := time.Parse("2006-01-02 15:04:05.000000", ultimaDataStr)

		placasMap[placa] = &DarwinTrechosResponse{
			Placa:         placa,
			KmPercorridos: totalKm,
			TotalViagens:  len(viagens),
			TempoTotal:    totalTempo,
			PrimeiraLocalizacao: LocalizacaoInfo{
				Data:        primeiraData.Format("02/01/2006 15:04"),
				Latitude:    primeiraViagem.LatitudeInicial,
				Longitude:   primeiraViagem.LongitudeInicial,
				Endereco:    primeiraViagem.AddressInicial,
				TempoViagem: primeiraViagem.Tempo,
			},
			UltimaLocalizacao: LocalizacaoInfo{
				Data:        ultimaData.Format("02/01/2006 15:04"),
				Latitude:    ultimaViagem.LatitudeFinal,
				Longitude:   ultimaViagem.LongitudeFinal,
				Endereco:    ultimaViagem.AddressFinal,
				TempoViagem: ultimaViagem.Tempo,
			},
		}
	}

	result := make([]DarwinTrechosResponse, 0, len(placasMap))
	for _, v := range placasMap {
		result = append(result, *v)
	}

	return result, nil
}
