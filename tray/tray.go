package tray

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"wsbrasil.com/simulate/kangu/kangu"
)

func MakeKangu(paramsPlatform Request) kangu.Api {
	var products []kangu.Products
	var productsTotalWeight float64 = 0.00
	var productsTotalPrice float64 = 0.00
	for _, v := range paramsPlatform.products {
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
		CepOrigem:  paramsPlatform.zipcode_source,
		CepDestino: paramsPlatform.zipcode_destination,
		VlrMerc:    productsTotalPrice,
		PesoMerc:   productsTotalWeight,
		QtdVol:     1,
		Origem:     "tray",
		Servicos:   []string{"E", "M", "X"},
		Produtos:   products,
	}
	return paramsKangu
}

func parseRequest(u *url.URL) Request {
	var params Request
	m, _ := url.ParseQuery(u.RawQuery)
	for k, v := range m {
		if k == "cep" {
			params.zipcode_source = v[0]
		} else if k == "cep_destino" {
			params.zipcode_destination = v[0]
		} else if k == "token" {
			params.token = v[0]
		} else if k == "prods" {
			product := strings.Split(v[0], "/")
			for _, vp := range product {
				item := strings.Split(vp, "|")
				length, _ := strconv.ParseFloat(item[0], 64)
				width, _ := strconv.ParseFloat(item[1], 64)
				height, _ := strconv.ParseFloat(item[2], 64)
				quantity, _ := strconv.ParseFloat(item[4], 64)
				weight, _ := strconv.ParseFloat(item[5], 64)
				price, _ := strconv.ParseFloat(item[7], 64)
				productItem := Products{
					Length:      length,
					Width:       width,
					Height:      height,
					Cubic:       item[3],
					Quantity:    quantity,
					Weight:      weight,
					ProductCode: item[6],
					Price:       price,
				}
				params.products = append(params.products, productItem)
			}
		}
	}
	return params
}

func MakeResultItems(freightItems []FreightItem) string {
	trayResponse := `<?xml version="1.0"?><cotacao>`
	resultItem := ""
	for _, item := range freightItems {
		resultItem = `			<resultado>
		    	<codigo>` + item.Code + `</codigo>
				<transportadora>` + item.Carrier + `</transportadora><servico></servico><transporte></transporte>
				<valor>` + fmt.Sprintf("%v", item.Price) + `</valor><peso></peso>
				<prazo_min>` + fmt.Sprintf("%v", item.MaximumForecast) + `</prazo_min>
				<prazo_max>` + fmt.Sprintf("%v", item.MaximumForecast) + `</prazo_max>
				<imagem_frete>` + item.CarrierImage + `</imagem_frete><aviso_envio></aviso_envio><entrega_domiciliar>1</entrega_domiciliar>
			</resultado>
		`
		trayResponse = trayResponse + resultItem
	}
	trayResponse = trayResponse + "</cotacao>"
	return trayResponse
}

func MakeResult(responseBody []byte) string {
	var freightItems []FreightItem
	json.Unmarshal(responseBody, &freightItems)
	return MakeResultItems(freightItems)
}

func Parse(decoder *json.Decoder, r *http.Request) ParseResult {
	getParams := strings.ReplaceAll(r.URL.RawQuery, ";", "|")
	u, err2 := url.Parse("/url?" + getParams)
	if err2 != nil {
		panic(err2)
	}

	paramsPlatform := parseRequest(u)

	validator := Validate(paramsPlatform)

	if len(validator) > 0 {
		return ParseResult{
			Validator: validator,
		}
	} else {
		paramsKangu := MakeKangu(paramsPlatform)

		return ParseResult{
			Token:  paramsPlatform.token,
			Params: paramsKangu,
		}
	}
}
