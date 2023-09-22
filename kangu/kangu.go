package kangu

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

var api = "https://simulacao.kangu.com.br/tms/transporte/simular"

type Api struct {
	CepOrigem  string     `json:"cepOrigem"`
	CepDestino string     `json:"cepDestino"`
	VlrMerc    float64    `json:"vlrMerc"`
	PesoMerc   float64    `json:"pesoMerc"`
	QtdVol     int64      `json:"qtdVol"`
	Origem     string     `json:"origem"`
	Servicos   []string   `json:"servicos"`
	Produtos   []Products `json:"produtos"`
}

type Products struct {
	Tipo        string  `json:"tipo"`
	Quantidade  string  `json:"quantidade"`
	Altura      float64 `json:"altura"`
	Largura     float64 `json:"largura"`
	Comprimento float64 `json:"comprimento"`
	Peso        float64 `json:"peso"`
	Valor       float64 `json:"valor"`
	Produto     string  `json:"produto"`
}

func Request(body []byte, tokenKangu string) []byte {
	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL
	customTransport := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	r, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	r.Header.Add("token", tokenKangu)

	client := &http.Client{Transport: customTransport}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	responseBody, _ := ioutil.ReadAll(res.Body)

	return responseBody

}
