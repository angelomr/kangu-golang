package pluggto

import "wsbrasil.com/simulate/kangu/kangu"

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
	Length    float64 `json:"length"`
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
	Token     string
	Params    kangu.Api
	Validator []string
}
