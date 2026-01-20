package darwin

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

type DarwinAPIScheme struct {
	url      string
	user     string
	password string
	code     int
}

type DarwinAPITrechosPayload struct {
	FirstDate string `json:"data_inicial"`
	LastDate  string `json:"data_final"`
}

func (d DarwinAPITrechosPayload) VerifyData() error {
	timeFormatter := "2006-01-02 15:04"

	first_date_formatted, err := time.ParseInLocation(timeFormatter, d.FirstDate, time.UTC)
	if err != nil {
		return fmt.Errorf("Data inicial deve estar no formato: YYYY-MM-DD HH:MM")
	}

	final_date_formatted, err := time.ParseInLocation(timeFormatter, d.LastDate, time.UTC)
	if err != nil {
		return fmt.Errorf("Data final deve estar no formato: YYYY-MM-DD HH:MM")
	}

	if first_date_formatted.After(final_date_formatted) {
		return fmt.Errorf("Data inicial não pode ser depois da data final.")
	}

	timeDeadline := time.Date(2026, time.December, 31, 23, 59, 0, 0, time.UTC)

	if first_date_formatted.After(timeDeadline) {
		return fmt.Errorf("Data inicial deve ser antes de 2026-12-31 23:59.")
	}

	if final_date_formatted.After(timeDeadline) {
		return fmt.Errorf("Data final deve ser antes de 2026-12-31 23:59")
	}

	return nil
}

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
	TempoTotal          int             `json:"tempo_total"` // em minutos
}

type LocalizacaoInfo struct {
	Data        string `json:"data"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Endereco    string `json:"endereco"`
	TempoViagem int    `json:"tempo_viagem"` // tempo dessa viagem específica
}

type DarwinAPILoginResponse struct {
	Token string `json:"token"`
}

type DarwinAPILoginPayload struct {
	CodClient int    `json:"cod_cliente"`
	Login     string `json:"login"`
	Password  string `json:"senha"`
}

func (d *DarwinAPIScheme) Load() (*string, error) {
	url_login := os.Getenv("DARWIN_API_LOGIN_URL")
	d.url = os.Getenv("DARWIN_API_URL")
	d.user = os.Getenv("DARWIN_API_LOGIN")
	d.code, _ = strconv.Atoi(os.Getenv("DARWIN_API_CODE"))
	d.password = os.Getenv("DARWIN_API_PASSWORD")

	// Log para debug
	fmt.Printf("=== DEBUG LOGIN ===\n")
	fmt.Printf("URL Login: %s\n", url_login)
	fmt.Printf("User: %s\n", d.user)
	fmt.Printf("Code: %d\n", d.code)
	fmt.Printf("Password: %s\n", "****")

	payload, _ := json.Marshal(DarwinAPILoginPayload{
		CodClient: d.code,
		Login:     d.user,
		Password:  d.password,
	})

	urlParams := url.Values{}
	urlParams.Add("cod_cliente", strconv.Itoa(d.code))
	urlParams.Add("login", d.user)
	urlParams.Add("senha", d.password)

	fullURL := url_login + "?" + urlParams.Encode()

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("Erro receber o token para a aplicação.\n%v", err)
	}

	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro ao solicitar request pelo client\n%v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("body completo: %v\n", string(body))

	bodyStr := string(body)
	token := strings.TrimSpace(bodyStr)
	token = strings.Trim(token, `"`)
	fmt.Printf("token: %s\n", token)

	// var Token DarwinAPILoginResponse
	// json.NewDecoder(resp.Body).Decode(&Token.Token)
	if token == "" {
		return nil, fmt.Errorf("Token não foi recebido.")
	}
	return &token, nil
}
