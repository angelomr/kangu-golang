# Gateway de Simulação de Frete na Kangu

Gateway de simulação de frete, onde faz a abstração das chamadas da Yampi e PluggTo para a API da Kangu.

## Kangu
- Documentação da API de Simulação de Frete: (https://portal.kangu.com.br/docs/api/transporte/#/M%C3%A9todos%20do%20Servi%C3%A7o/post_simular)

### Requisitos de negócio
É necessário ter o Token da Kangu, acessivel através do portal Kangu (https://portal.kangu.com.br), em "Meu Acesso".

## Yampi
- Documentação da API de Frete: (https://docs.yampi.com.br/referencia-da-api/logistica#api-de-frete)

### Requisitos de negócio
Na configuração de integração da Yampi, você deve configurar:
- Nome: Kangu 
- URL: http://simulador.wserp.com.br - servidor de testes da Web Studio Brasil (https://wsbrasil.com)
- Headers:
    - Token - Aqui você deve colocar o Token da Kangu
    - cepOrigem - Aqui você deve cadastrar o CEP da sua loja, da onde irão ser postados os pacotes.

## PluggTo
- Documentação da API - Freight (Source): (https://developers.plugg.to/reference/freight-source)

### Requisitos de negócio
Eu não tenho um acesso de Testes na PluggTo, então somente segui a documentação para criação da simulação de frete. Na configuração de integração da PluggTo, você deve configurar o Token no Header:
- Token - Aqui você deve colocar o Token da Kangu

### Importante
Caso a PluggTo não suporte a criação de variáveis no Header, podemos alterar o código para que os dados sejam enviados através de parâmetros de URL, porém, conforme dito acima, não temos acesso a PluggTo para poder fazer um teste e verificar como isso funcionaria.

## Tray
- Documentação da API - FRETE X API: (https://developers.tray.com.br/#frete-x-api)

### Requisitos de negócio
Para configurar o Gateway de Frete na loja da Tray, basta acessar o menu Configurações > Frete e envio na área administrativa da loja.

Terão que ser preenchidos 2 campos:
- URL do WebService: http://simulador.wserp.com.br
- Token de identificação do cliente (Token da Kangu)
