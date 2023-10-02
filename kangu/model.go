package kangu

type Api struct {
	CepOrigem  string     `json:"cepOrigem"`
	CepDestino string     `json:"cepDestino"`
	VlrMerc    float64    `json:"vlrMerc"`
	PesoMerc   float64    `json:"pesoMerc"`
	QtdVol     int64      `json:"qtdVol"`
	Origem     string     `json:"origem"`
	Servicos   []string   `json:"servicos"`
	Produtos   []Products `json:"produtos"`
}

type Products struct {
	Tipo        string  `json:"tipo"`
	Quantidade  string  `json:"quantidade"`
	Altura      float64 `json:"altura"`
	Largura     float64 `json:"largura"`
	Comprimento float64 `json:"comprimento"`
	Peso        float64 `json:"peso"`
	Valor       float64 `json:"valor"`
	Produto     string  `json:"produto"`
}
