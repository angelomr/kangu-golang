package yampi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wsbrasil.com/simulate/kangu/kangu"
)

type Request struct {
	Zipcode string     `json:"zipcode"`
	Amount  float64    `json:"amount"`
	Skus    []Products `json:"skus"`
}

type Products struct {
	Id               int     `json:"id"`
	ProductId        int     `json:"product_id"`
	Sku              string  `json:"sku"`
	Price            float64 `json:"price"`
	Quantity         float64 `json:"quantity"`
	Length           float64 `json:"lenght"`
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	Weight           float64 `json:"weight"`
	AvailabilityDays int     `json:"availability_days"`
}

type Result struct {
	Quotes []QuoteItem `json:"quotes"`
}

type QuoteItem struct {
	Name    string  `json:"name"`
	Service string  `json:"service"`
	Price   float64 `json:"price"`
	Days    int64   `json:"days"`
	QuoteId int64   `json:"quote_id"`
}

type FreightItem struct {
	Code            int64   `json:"idSimulacao"`
	Carrier         string  `json:"descricao"`
	Price           float64 `json:"vlrFrete"`
	MaximumForecast int64   `json:"prazoEnt"`
	CarrierImage    string  `json:"url_logo"`
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

	cepOrigem := r.Header.Get("Ceporigem")
	tokenKangu := r.Header.Get("Token")
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
	for _, v := range paramsPlatform.Skus {

		products = append(products, kangu.Products{
			Tipo:        "C",
			Quantidade:  fmt.Sprintf("%.2f", v.Quantity),
			Altura:      v.Height,
			Largura:     v.Width,
			Comprimento: v.Length,
			Peso:        v.Weight,
			Valor:       v.Price,
			Produto:     "C",
		})
		productsTotalWeight += v.Weight
		productsTotalPrice += v.Price
	}

	paramsKangu := kangu.Api{
		CepOrigem:  cepOrigem,
		CepDestino: paramsPlatform.Zipcode,
		VlrMerc:    productsTotalPrice,
		PesoMerc:   productsTotalWeight,
		QtdVol:     1,
		Origem:     "yampi",
		Servicos:   []string{"E", "M", "X"},
		Produtos:   products,
	}
	return paramsKangu
}

func MakeResultItems(freightItems []FreightItem) []QuoteItem {
	var resultItems []QuoteItem
	for _, item := range freightItems {
		resultItems = append(resultItems, QuoteItem{
			item.Carrier,
			"1",
			item.Price,
			item.MaximumForecast,
			item.Code,
		})
	}
	return resultItems
}

func MakeResult(responseBody []byte) Result {
	var freightItems []FreightItem
	json.Unmarshal(responseBody, &freightItems)

	return Result{
		Quotes: MakeResultItems(freightItems),
	}
}
