package yampi

import (
	"fmt"
	"strconv"
	"testing"

	"wsbrasil.com/simulate/kangu/kangu"
)

func TestParse(t *testing.T) {
	var productsResult []kangu.Products
	productsResult = append(productsResult, kangu.Products{
		Tipo:        "C",
		Quantidade:  "1.00",
		Altura:      2,
		Largura:     4,
		Comprimento: 16,
		Peso:        0.3,
		Valor:       15.25,
		Produto:     "C",
	})

	var resultKangu = kangu.Api{
		CepOrigem:  "04191270",
		CepDestino: "04126100",
		VlrMerc:    15.25,
		PesoMerc:   0.3,
		QtdVol:     1,
		Origem:     "yampi",
		Servicos:   []string{"E", "M", "X"},
		Produtos:   productsResult,
	}

	var requestProducts []Products
	quantityTemp, _ := strconv.ParseFloat(productsResult[len(productsResult)-1].Quantidade, 64)
	requestProducts = append(requestProducts, Products{
		Id:               1,
		ProductId:        1,
		Sku:              "0011",
		Price:            productsResult[len(productsResult)-1].Valor,
		Quantity:         quantityTemp,
		Length:           productsResult[len(productsResult)-1].Altura,
		Width:            productsResult[len(productsResult)-1].Largura,
		Height:           productsResult[len(productsResult)-1].Comprimento,
		Weight:           productsResult[len(productsResult)-1].Peso,
		AvailabilityDays: 0,
	})

	requestParamsPlatform := Request{
		Zipcode: resultKangu.CepDestino,
		Amount:  resultKangu.VlrMerc,
		Skus:    requestProducts,
	}

	resultado := MakeKangu(
		requestParamsPlatform,
		resultKangu.CepOrigem,
	)

	fmt.Println(requestParamsPlatform)
	fmt.Println(resultado)
	fmt.Println(resultKangu)

	t.Run("Verifica se CepOrigem está Correto", func(t *testing.T) {

		if resultado.CepOrigem != resultKangu.CepOrigem {
			t.Errorf("resultado '%s', esperado '%s'", resultado.CepOrigem, resultKangu.CepOrigem)
		}
	})

	t.Run("Verifica se CepDestino está Correto", func(t *testing.T) {
		if resultado.CepDestino != resultKangu.CepDestino {
			t.Errorf("resultado '%s', esperado '%s'", resultado.CepDestino, resultKangu.CepDestino)
		}
	})

	t.Run("Verifica se VlrMerc está Correto", func(t *testing.T) {
		if resultado.VlrMerc != resultKangu.VlrMerc {
			t.Errorf("resultado '%s', esperado '%s'", fmt.Sprintf("%.2f", resultado.VlrMerc+123), fmt.Sprintf("%.2f", resultKangu.VlrMerc))
		}
	})

	t.Run("Verifica se PesoMerc está Correto", func(t *testing.T) {
		if resultado.PesoMerc != resultKangu.PesoMerc {
			t.Errorf("resultado '%s', esperado '%s'", fmt.Sprintf("%.2f", resultado.PesoMerc), fmt.Sprintf("%.2f", resultKangu.PesoMerc))
		}
	})

	t.Run("Verifica se Quantidade do Produto está Correto", func(t *testing.T) {
		if resultado.Produtos[len(productsResult)-1].Quantidade != resultKangu.Produtos[len(productsResult)-1].Quantidade {
			t.Errorf(
				"resultado '%s', esperado '%s'",
				resultado.Produtos[len(productsResult)-1].Quantidade,
				resultKangu.Produtos[len(productsResult)-1].Quantidade,
			)
		}
	})

}
