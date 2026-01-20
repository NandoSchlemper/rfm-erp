package models

type DarwinTrechosResponse struct {
	Placa         string  `json:"placa"`
	KmPercorridos float64 `json:"km"`
}

type DarwinTrechosRequest struct {
	Initial_date string `json:"data_inicial"`
	Final_date   string `json:"data_final"`
}
