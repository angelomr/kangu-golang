package tray

var notNull string = "não pode ser vazio"
var defaultText string = "O campo"

func Validate(paramsPlatform Request) []string {
	var errors []string

	if len(paramsPlatform.token) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Token", notNull))
	}

	if len(paramsPlatform.zipcode_source) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Cep", notNull))
	}

	if len(paramsPlatform.zipcode_destination) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "CepDestino", notNull))
	}

	if len(paramsPlatform.products) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Products", notNull))
	} else {
		for _, v := range paramsPlatform.products {
			if v.Quantity < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Quantidade", notNull))
			}

			if v.Height < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Altura", notNull))
			}

			if v.Width < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Largura", notNull))
			}

			if v.Length < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Comprimento", notNull))
			}

			if v.Weight < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Peso", notNull))
			}

			if v.Price < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Prods - Preço", notNull))
			}
		}
	}

	return errors
}

func SetErrorMessage(text string, field string, errorType string) string {
	return text + " " + field + " " + errorType
}
