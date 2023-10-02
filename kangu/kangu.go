package kangu

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

var api = "https://simulacao.kangu.com.br/tms/transporte/simular"

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
