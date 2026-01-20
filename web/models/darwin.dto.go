package models

type DarwinAPITrechosResponse struct {
	Placa            string         `json:"placa"`
	DataInicial      map[string]any `json:"data_inicial"`
	DataFinal        map[string]any `json:"data_final"`
	LatitudeInicial  string         `json:"latitude_inicial"`
	LongitudeInicial string         `json:"longitude_inicial"`
	LatitudeFinal    string         `json:"latitude_final"`
	LongitudeFinal   string         `json:"longitude_final"`
	Tempo            int            `json:"tempo"`
	KmPercorrido     float64        `json:"km_percorrido"`
	AddressInicial   string         `json:"address_inicial"`
	AddressFinal     string         `json:"address_final"`
}

type DarwinTrechosResponse struct {
	Placa               string          `json:"placa"`
	KmPercorridos       float64         `json:"km"`
	PrimeiraLocalizacao LocalizacaoInfo `json:"primeira_localizacao"`
	UltimaLocalizacao   LocalizacaoInfo `json:"ultima_localizacao"`
	TotalViagens        int             `json:"total_viagens"`
	TempoTotal          int             `json:"tempo_total"`
}

type LocalizacaoInfo struct {
	Data        string `json:"data"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Endereco    string `json:"endereco"`
	TempoViagem int    `json:"tempo_viagem"`
}

type DarwinTrechosRequest struct {
	Initial_date string `json:"data_inicial"`
	Final_date   string `json:"data_final"`
}
