# Listagem de Ordens 

Esse projeto é uma aplicação que consiste em cadastrar uma Ordem de serviço e listar as Ordens de seriços já cadastradas.

## Descrição

O principal objetivo desse projeto é cadastrar e consultar Orders via REST, GraphQL e gRPC. O projeto foi estruturado seguindo os padões de arquitetura  `Clean Architecture` desenvolvidos pelo [Uncle Bob](http://cleancoder.com/).

O programa está estruturado utilizando clean arquiteture com as seguintes bibliotecas para desenvolvimento  
 - [go-chi - REST](https://github.com/go-chi/chi)
 - [gRPC](https://grpc.io/docs/languages/go/quickstart/)
 - [gqlgen - GraphQL](https://github.com/99designs/gqlgen)
 - [amqp - RabbitMQ](github.com/streadway/amqp)
 - [golang-migrate](https://github.com/golang-migrate/migrate)
 - [testify](https://github.com/stretchr/testify)
 - [wire](https://github.com/google/wire)
 - [DB Mysql](https://github.com/go-sql-driver/mysql)


## Funcionamento

1. **Requisições REST:**
   
   Para requisições REST temos 2 endpoints um `POST /order` para criar novas ordens de serviço e `GET /order?page=1&limit=20` para obter os resultados paginados.

   Exemplo `POST /order`

  Request 
  ```json
     {
      "id":"41c854b0-32c6-4cc9-9f7b-1f833fff7f17",
      "price": 100,
      "tax": 1
     }
  ```
   Response
   ```json
    {
      "id": "41c854b0-32c6-4cc9-9f7b-1f833fff7f17",
      "price": 100,
      "tax": 1,
      "final_price": 101
    }
  ```
   Exemplo `GET /order?page=1&limit=1`

   Response
   ```json
    {
    "orders": [
        {
        "id": "41c854b0-32c6-4cc9-9f7b-1f833fff7f17",
        "price": 100,
        "tax": 1,
        "final_price": 101
        }
    ],
    "currentPage": 1,
    "totalPages": 1
    }
  ```
   OBS: Dentro da pasta  `api` existem arquivos .http para serem usados na extensão do vscode [REST_CLIENT](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

 

 
2. **Requisições GraphQL:**
   Para requisições GraphQL podemos utilizar a interface do [gqlgen](https://github.com/99designs/gqlgen) acessando-a na porta `:8080`:

   Exemplo `mutation CreateOrder`

  Request 
  ```graphql
      mutation CreateOrder {
        createOrder(input: {
          id: "41c854b0-32c6-4cc9-9f7b-1f833fff7f17"
          Price: 150.0
          Tax: 15.0
        }) {
          id
          Price
          Tax
          FinalPrice
        }
      }
  ```
   Response
   ```json
      {
          "data": {
          "createOrder": {
              "id": "41c854b0-32c6-4cc9-9f7b-1f833fff7f17",
              "Price": 150,
              "Tax": 15,
              "FinalPrice": 165
          }
          }
      }
  ```
   Exemplo `query GetOrders`

   Response
  ```graphql
      query GetOrders {
          orders(page: 1, limit: 1) {
              orders {
                  id
                  Price
                  Tax
                  FinalPrice
              }
              currentPage
              totalPages
          }
      }
  ```
   Response
   ```json
      {
          "data": {
              "orders": {
              "orders": [
                  {
                  "id": "41c854b0-32c6-4cc9-9f7b-1f833fff7f17",
                  "Price": 100,
                  "Tax": 1,
                  "FinalPrice": 101
                  }
              ],
              "currentPage": 1,
              "totalPages": 1
              }
          }
      }
  ```

3. **Requisições gRPC:**
   Para requisições gRPC recomendo utilizar a [evans](https://github.com/ktr0731/evans) acessando-a na porta `:50051`:


## Pré-requisitos

 - Para executar a aplicação, é necessário ter o Go instalado na máquina. A instalação pode ser feita conforme as instruções no [site oficial do Go](https://golang.org/dl/);
 - Para gerar novas migrações você deve ter instalado em sua maquina o [golang-migrate](https://github.com/golang-migrate/migrate);
 - (opcional) Para executar comandos no gRPC recomendo a utilização do [evans](https://github.com/ktr0731/evans).


## Instruções de Execução

1. Clone o repositório e navegue até o diretório do projeto.
   
   ```bash
    git clone https://github.com/brunoofgod/goexpert-lesson-3.git
    cd goexpert-lesson-3
   ```

2. Suba os containers docker:
   
   ```bash
    docker-compose up -d
   ```
