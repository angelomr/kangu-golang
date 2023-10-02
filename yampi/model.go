package yampi

import "wsbrasil.com/simulate/kangu/kangu"

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
	Length           float64 `json:"length"`
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
	Token     string
	Params    kangu.Api
	Validator []string
}
