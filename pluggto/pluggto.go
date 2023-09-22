package pluggto

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wsbrasil.com/simulate/kangu/kangu"
)

type Request struct {
	DestinationPostalCode string     `json:"destination_postalcode"`
	AmountTotal           float64    `json:"amount"`
	Products              []Products `json:"products"`
	OriginPostalCode      string     `json:"origin_postalcode"`
}

type Products struct {
	Sku       string  `json:"sku"`
	UnitPrice float64 `json:"unit_price"`
	Quantity  string  `json:"quantity"`
	Length    float64 `json:"lenght"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
}

type QuoteItem struct {
	Method       string `json:"method"`
	Company      string `json:"company"`
	DeliveryType string `json:"delivery_type"`
	Message      string `json:"message"`
	Estimative   int64  `json:"estimative"`
	Price        string `json:"price"`
}

type FreightItem struct {
	Code            int64   `json:"idSimulacao"`
	Company         string  `json:"transp_nome"`
	Carrier         string  `json:"descricao"`
	Price           float64 `json:"vlrFrete"`
	MaximumForecast int64   `json:"prazoEnt"`
	CarrierImage    string  `json:"url_logo"`
	Method          string  `json:"servico"`
}

type ParseResult struct {
	Token  string
	Params kangu.Api
}

func Parse(decoder *json.Decoder, r *http.Request) ParseResult {
	var paramsPlatform Request
	err3 := decoder.Decode(&paramsPlatform)
	if err3 != nil {
		panic(err3)
	}

	cepOrigem := paramsPlatform.OriginPostalCode
	tokenKangu := r.Header.Get("token")
	paramsKangu := MakeKangu(paramsPlatform, cepOrigem)

	return ParseResult{
		Token:  tokenKangu,
		Params: paramsKangu,
	}
}

func MakeKangu(paramsPlatform Request, cepOrigem string) kangu.Api {
	var products []kangu.Products
	var productsTotalWeight float64 = 0.00
	var productsTotalPrice float64 = 0.00
	for _, v := range paramsPlatform.Products {
		products = append(products, kangu.Products{
			Tipo:        "C",
			Quantidade:  v.Quantity,
			Altura:      v.Height,
			Largura:     v.Width,
			Comprimento: v.Length,
			Peso:        v.Weight,
			Valor:       v.UnitPrice,
			Produto:     "C",
		})
		productsTotalWeight += v.Weight
		productsTotalPrice += v.UnitPrice
	}

	params := kangu.Api{
		CepOrigem:  cepOrigem,
		CepDestino: paramsPlatform.DestinationPostalCode,
		VlrMerc:    productsTotalPrice,
		PesoMerc:   productsTotalWeight,
		QtdVol:     1,
		Origem:     "plugg",
		Servicos:   []string{"E", "M", "X"},
		Produtos:   products,
	}
	// fmt.Println(paramsKangu)
	return params
}

func MakeResultItems(freightItems []FreightItem) []QuoteItem {
	services := make(map[string]string)
	services["P"] = "Postagem"
	services["C"] = "Coleta"
	services["R"] = "Retira"
	services["E"] = "Entrega Normal"
	services["X"] = "Entrega Expressa"
	var resultItems []QuoteItem
	for _, item := range freightItems {
		// fmt.Println(item)
		resultItems = append(resultItems, QuoteItem{
			item.Carrier,
			item.Company,
			services[item.Method],
			"",
			item.MaximumForecast,
			fmt.Sprintf("%.2f", item.Price),
		})
	}
	return resultItems
}

func MakeResult(responseBody []byte) []QuoteItem {
	var freightItems []FreightItem
	json.Unmarshal(responseBody, &freightItems)

	return MakeResultItems(freightItems)
}
