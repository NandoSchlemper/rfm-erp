package wrsat

import "os"

type WrsatAPIScheme struct {
	url      string
	user     string
	password string
	code     string
}

type WrsatAPIPayload struct {
	User        string `json:"usuario"`
	Password    string `json:"senha"`
	Order       string `json:"ordem"`
	Limit       string `json:"limit"`
	Page        string `json:"pagina"`
	Description string `json:"descricao"`
}

func (w *WrsatAPIScheme) Load() {
	w.url = os.Getenv("WRSAT_API_URL")
	w.user = os.Getenv("WRSAT_API_USER")
	w.password = os.Getenv("WRSAT_API_PASSWORD")
	w.code = os.Getenv("WRSAT_API_CODE")
}
