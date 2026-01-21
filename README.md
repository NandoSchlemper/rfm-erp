# RFM Transportes ERP

## Sobre

Aplicação com o objetivo de centralizar os dados recebidos de diversas APIs da RFM Transportes para uma análise mais concisa da frota. 

## Estrutura do projeto

```sh
<root>
|-- cmd/ # Backend
|-- |-- api/main.go # Instancia da API
|-- |-- external/ # consumo de APIs externas
|-- |-- |-- darwin/ # consumo da API do Darwin (Sonda de Combustivel)
|-- |-- |-- wrsat/ # consumo da API de rastreamento (localização atual)
|-- |-- internal/ # lógica principal do dominio
|-- |-- |-- features/ # componentes essenciais para a aplicação
|-- web/ # Frontend
|-- |-- external/ # APIs para consumo do Frontennd
|-- |-- |-- rfm/ # consumo do backend
|-- |-- views/ # visualizações e handler
|-- |-- main.go # instancia do frontend
```

## Ferramentas

- Golang v1.25.5
- Fiber v3.x