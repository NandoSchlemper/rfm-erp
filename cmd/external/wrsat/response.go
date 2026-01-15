package wrsat

type WrsatVehicle struct {
	ID         string `json:"id"`
	IDVeiculo  string `json:"idveiculo"`
	Principal  string `json:"principal"`
	WebGrupoID string `json:"web_grupo_id"`
	Placa      string `json:"placa"`
	// Descricao      string `json:"descricao"`
	// DataGPS        string `json:"datagps"`
	// Localizacao    string `json:"localizacao"`
	// StatusIgn      int    `json:"statusign"`
	// StatusGPS      int    `json:"statusgps"`
	// Velocidade     int    `json:"velocidade"`
	// Direcao        string `json:"direcao"`
	// Lat            string `json:"lat"`
	// Lon            string `json:"lon"`
	// Satelites      int    `json:"satelites"`
	// NivelCel       int    `json:"nivelcel"`
	// MovingStatus   int    `json:"moving_status"`
	// GprsConnection int    `json:"gprs_connection"`
	// GsmJamming     string `json:"gsm_jamming"`
	// Ancora         int    `json:"ancora"`
	// WebGrupo       string `json:"web_grupo"`
	// OdometerCan    string `json:"odometerCan"`
	// Odometer       string `json:"odometer"`
	// Blocked        string `json:"blocked"`
}

type WrsatListVehiclesResponse struct {
	Erro      bool           `json:"erro"`
	Status    int            `json:"status"`
	Mensagem  string         `json:"mensagem"`
	Ordem     string         `json:"ordem"`
	Limit     string         `json:"limit"`
	Pagina    string         `json:"pagina"`
	QtdResult int            `json:"qtd_result"`
	Dados     []WrsatVehicle `json:"dados"`
}
