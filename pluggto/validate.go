package pluggto

import "strconv"

var notNull string = "não pode ser vazio"
var defaultText string = "O campo"

func Validate(paramsPlatform Request, CepOrigem string, TokenKangu string) []string {
	var errors []string

	if len(TokenKangu) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Token", notNull))
	}

	if len(CepOrigem) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "origin_postalcode", notNull))
	}

	if len(paramsPlatform.DestinationPostalCode) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "destination_postalcode", notNull))
	}

	if len(paramsPlatform.Products) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Products", notNull))
	} else {
		for _, v := range paramsPlatform.Products {
			quantity, _ := strconv.ParseFloat(v.Quantity, 64)
			if quantity < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - Quantity", notNull))
			}

			if v.Height < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - Height", notNull))
			}

			if v.Width < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - Width", notNull))
			}

			if v.Length < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - Length", notNull))
			}

			if v.Weight < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - Weight", notNull))
			}

			if v.UnitPrice < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Products - UnitPrice", notNull))
			}
		}
	}

	return errors
}

func SetErrorMessage(text string, field string, errorType string) string {
	return text + " " + field + " " + errorType
}
