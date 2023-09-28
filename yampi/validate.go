package yampi

var notNull string = "não pode ser vazio"
var defaultText string = "O campo"

func Validate(paramsPlatform Request, CepOrigem string, TokenKangu string) []string {
	var errors []string

	if len(TokenKangu) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Token", notNull))
	}

	if len(CepOrigem) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "CepOrigem", notNull))
	}

	if len(paramsPlatform.Zipcode) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Zipcode", notNull))
	}

	if paramsPlatform.Amount < 0.01 {
		errors = append(errors, SetErrorMessage(defaultText, "Amount", notNull))
	}

	if len(paramsPlatform.Skus) == 0 {
		errors = append(errors, SetErrorMessage(defaultText, "Skus", notNull))
	} else {
		for _, v := range paramsPlatform.Skus {
			if v.Quantity < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Quantity", notNull))
			}

			if v.Height < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Height", notNull))
			}

			if v.Width < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Width", notNull))
			}

			if v.Length < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Length", notNull))
			}

			if v.Weight < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Weight", notNull))
			}

			if v.Price < 0.01 {
				errors = append(errors, SetErrorMessage(defaultText, "Skus - Price", notNull))
			}
		}
	}

	return errors
}

func SetErrorMessage(text string, field string, errorType string) string {
	return text + " " + field + " " + errorType
}
