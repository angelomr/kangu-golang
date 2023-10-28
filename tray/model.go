package tray

import "wsbrasil.com/simulate/kangu/kangu"

type Request struct {
	token               string
	zipcode_source      string
	zipcode_destination string
	products            []Products
}

type Products struct {
	Length      float64
	Width       float64
	Height      float64
	Cubic       string
	Quantity    float64
	Weight      float64
	ProductCode string
	Price       float64
}

type Result struct {
	Quotes string
}

type FreightItem struct {
	Code            string           `json:"referencia"`
	Carrier         string           `json:"descricao"`
	Price           float64          `json:"vlrFrete"`
	MaximumForecast int              `json:"prazoEnt"`
	CarrierImage    string           `json:"url_logo"`
	Error           FreightItemError `json:"error"`
}

type FreightItemError struct {
	Code    int    `json:"codigo"`
	Message string `json:"mensagem"`
}

type ParseResult struct {
	Token     string
	Params    kangu.Api
	Validator []string
}
