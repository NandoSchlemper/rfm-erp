package darwin

import (
	"bytes"
	"net/http"
	"os"
	"strconv"

	"github.com/goccy/go-json"
)

type DarwinAPIScheme struct {
	url      string
	token    string
	user     string
	password string
	code     int
}

type DarwinAPITrechosPayload struct {
	FirstDate string `json:"data_inicial"`
	LastDate  string `json:"data_final"`
}

type DarwinAPITrechosResponse struct {
	Trecho         string  `json:"trecho"`
	TotalParadas   int     `json:"total_paradas"`
	KmPercorrindos float64 `json:"km_percorridos"`
}

type DarwinAPILoginPayload struct {
	CodClient int    `json:"cod_cliente"`
	Login     string `json:"login"`
	Password  string `json:"senha"`
}

func (d *DarwinAPIScheme) Load() {
	url_login := os.Getenv("DARWIN_API_LOGIN_URL")
	d.url = os.Getenv("DARWIN_API_URL")
	d.user = os.Getenv("DARWIN_API_LOGIN")
	code, _ := strconv.Atoi(os.Getenv("DARWIN_API_CODE"))
	d.code = code
	d.password = os.Getenv("DARWIN_API_PASSWORD")

	payload, _ := json.Marshal(DarwinAPILoginPayload{
		CodClient: code,
		Login:     d.user,
		Password:  d.password,
	})

	resp, _ := http.Post(url_login, "application/json", bytes.NewBuffer(payload))
	resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&d.token)
}
