# Desafio Técnico - Go(lang) - api-banking 
Desafio stone em Go para criação de api baking.

<a id="ancora"></a>
# Desafio e Instrunções
- [Desafio - Documentação de Rotas](#ancora1)
- [Dendências e Modo de Uso](#ancora2)
- [Instruções e Regras](#ancora3)


<a id="ancora1"></a>
# Documentação do Desafio
**API de transferencia entre contas Internas de um banco digital.**
## Rotas e exemplos


#### CRIAÇÃO DE CONTAS: `POST /accounts`
- body:
```json
{
    "name": "My Account A",
    "cpf": "12345678900",
    "secret": "12345678",
    "balance": 10000
}
```
- Especificação dos campos obrigatórios: 
    - `cpf`: deve ter exatos 11 carateres.
    - `secret`: deve ter 8 ou mais caracteres.
    - `balance`: poderá ser omitido, porém receberá como valor inicial, 0 (zero).
- Casos de respostas
  * Sucesso, retorna o ID da conta:
     + status code: `201` 
     + content-type: `application/json`
     + body:
    ```json
      "cc5c3af1-e1ba-47ee-94b5-bd7fda44999b"
    ```
  * Falhas:
     1. `cpf` já existe registrado na base dados:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "CPF already exists"
            }
        ```
    1. `secret` possui tamanho menor que o requisitado:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "password must be at least 8 characters long"
            }   
        ```
    1. `cpf` possui tamanho maior ou menor que o requisitado:
        - status code: `409`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "CPF must be exactly 11 characters long"
            }
        ```
    1. `id` ocorreu um erro na geração do id:
        - status code: `409`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "ID account already exists"
            }
        ```
#### LISTANDO TODAS AS CONTAS: `GET /accounts`
- Casos de respostas
  * Sucesso, retorna todas as contas cadastradas:
     + status code: `200` 
     + content-type: `application/json`
     + body:
     ```json
        [
            {
                "id": "dc4e264b-67c4-4d2f-9104-4446eb8be402",
                "name": "My Account A",
                "cpf": "12345678900",
                "secret": "",
                "balance": 10000,
                "created_at": "2022-04-29T11:03:57.681663578Z"
            },
            {
                "id": "b216839b-da10-4e9e-8a6d-e6303c50277c",
                "name": "My Account B",
                "cpf": "12345678911",
                "secret": "",
                "balance": 2000,
                "created_at": "2022-04-29T11:04:37.038091299Z"
            }
        ]
    ```
  * Em caso de falha interna no servidor:
     + status code: `500` 
     + content-type: `application/json`
     + body:
        ```json
            {
                "Error": "internal server error"
            }
        ```
#### EXIBIR SALDO DE UMA CONTA: `GET /accounts/{account_id}/balance`
- account_id: `dc4e264b-67c4-4d2f-9104-4446eb8be402`
- Casos de respostas
  * Sucesso, retorna o saldo da conta informada:
     + status code: `200` 
     + content-type: `application/json`
     + body: 
     ```json
        2000
     ```
   * Falhas:
     1. Para account_id vazio:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "empty uuid"
            }
        ```
     1. Para account_id de formato inválido:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "invalid accountID format"
            }
        ```
     1. Não é encontrada uma conta correspondente:
        - status code: `404`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "account not found"
            }
        ```
#### LOGIN: `POST /login`
- body:
```json
{
    "cpf": "12345678900",
    "secret": "12345678"
}
```
- Especificação dos campos obrigatórios: 
    - `cpf`: deve ser igual a seu cpf de cadastro.
    - `secret`: deve ser igual a sua senha de cadastro.
- Casos de respostas
  * Sucesso, retorna um token jwt de autorização:
     + status code: `200` 
     + content-type: `application/json`
     + body:
    ```json
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTEyMzcxNTEsInVzZXJfaWQiOiJkYzRlMjY0Yi02N2M0LTRkMmYtOTEwNC00NDQ2ZWI4YmU0MDIifQ.nUkxwMwzclFzXe7cHrw3cnImlfkEfMkRjLSKTCGRtlk"
    ```
   * Falhas:
     1. Para erro no processamento do login:
        - status code: `422`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "EOF"
            }
        ```
     1. Conta não encontrada:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "account not found"
            }
        ```
     1. Senha inválida:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "validating password"
            }
        ```
#### CRIAR UMA TRANSFERÊNCIA: `POST /transfers`
- body:
```json
{
    "account_destination_ID": "b216839b-da10-4e9e-8a6d-e6303c50277c",
    "amount":2500
}
```
- Authentication: `Bearer Token` : 
```
 0eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTEyMzcxNTEsInVzZXJfaWQiOiJkYzRlMjY0Yi02N2M0LTRkMmYtOTEwNC00NDQ2ZWI4YmU0MDIifQ.nUkxwMwzclFzXe7cHrw3cnImlfkEfMkRjLSKTCGRtlk
```
- Especificação dos campos obrigatórios: 
    - `account_destination_ID`: não pode ser vazio, nem pode ser igual ao id da conta de origem.
    - `amount`: deverá ser um valor maior que 0.
- Casos de respostas
  * Sucesso, retorna o id da transação:
     + status code: `201` 
     + content-type: `application/json`
     + body:
    ```json
      "38d7f184-8972-4c3f-82f5-609e9d1a1338"
    ```
   * Falhas:
     1. Para erro na decodifição do Json:
        - status code: `400`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "EOF"
            }
        ```
     1. Na extração do ID do token:
        - status code: `401`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "Invalid Token"
            }
        ```
     1. Na validação do tamanho da conta de destino:
        - status code: `412`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "empty ID not allowed"
            }
        ```
     1. Na validação do valor do saldo:
        - status code: `412`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "invalid amount, be must > 0"
            }
        ```
     1. Na validação da conta de destino, quando igual a conta de origem:
        - status code: `412`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "equal source and destination accounts are not allowed."
            }
        ```
     1. Conta de destino não existe:
        - status code: `404`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "account not found"
            }
        ```
     1. Em caso de duplicidade de ID na criação da transação:
        - status code: `409`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "ID already exists"
            }
        ```
     1. Em caso de erro interno do servidor ao atualzar saldo:
        - status code: `500`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "internal server error"
            }
        ```
#### LISTANDO TODAS AS TRANSAÇÕES DE UM USUÁRIO AUTENTICADO: `GET /transfers`
- Authentication: `Bearer Token` : 
```
 0eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTEyMzcxNTEsInVzZXJfaWQiOiJkYzRlMjY0Yi02N2M0LTRkMmYtOTEwNC00NDQ2ZWI4YmU0MDIifQ.nUkxwMwzclFzXe7cHrw3cnImlfkEfMkRjLSKTCGRtlk
```

- Casos de respostas
  * Sucesso, retorna todas as transações cadastradas:
     + status code: `200` 
     + content-type: `application/json`
     + body:
     ```json
        [
            {
                "id": "38d7f184-8972-4c3f-82f5-609e9d1a1338",
                "account_origin_id": "dc4e264b-67c4-4d2f-9104-4446eb8be402",
                "account_destination_id": "b216839b-da10-4e9e-8a6d-e6303c50277c",
                "amount": 2500,
                "created_at": "2022-04-29T13:01:13.191001585Z"
            },
            {
                "id": "c998049a-7b21-4ff2-817b-7f030a2bb0d1",
                "account_origin_id": "dc4e264b-67c4-4d2f-9104-4446eb8be402",
                "account_destination_id": "b216839b-da10-4e9e-8a6d-e6303c50277c",
                "amount": 2500,
                "created_at": "2022-04-29T13:32:18.113337339Z"
            }
        ]
    ```
  * Falhas:
     1. Na extração do ID do token:
        - status code: `401`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "Invalid Token"
            }
        ```
     1. Em caso de erro interno do servidor:
        - status code: `401`
        - content-type: `application/json`
        - body:
        ```json
            {
                "Error": "internal server error"
            }
        ```
[Topo](#ancora)

* * * *
* * * *
<a id="ancora2"></a>
# Como Executar e Dependências

## Modo de usar com Docker

- Faça o pull para sua máquina local: 
```
    docker pull rodolfo1992/app-banking:1.0
```

- Para rodar o container:
```
    docker run -it -p 8080:3000 app-banking
```

## Modo de usar sem o Docker

- API baseada na versão: `go version go1.18.1`

- Faça o Clone deste repositório.

- Baixe as dependências:
```
    go mod tidy
```

- Execute a aplicação com:
```
    go run main.go
```

## Dependências necessárias para o funcionamento da API

  - JWT-GO: github.com/golang-jwt/jwt/v4 v4.4.1
  -	UUID: github.com/google/uuid v1.3.0
  - Gorilla Mux: github.com/gorilla/mux v1.8.0
  - BCrypt: github.com/joho/godotenv v1.4.0
  - GoDoEnv: golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064

[Topo](#ancora)
* * * *
* * * *
<a id="ancora3"></a>


## Instruções e Regras do Desafio

**O desafio é criar uma API de transferencia entre contas Internas de um banco digital.**

#### Regras gerais

* Usar formato JSON para leitura e escrita. (ex: `GET /accounts/` retorna json, `POST /accounts/ {name: 'james bond'}`)

#### Rotas esperadas

##### `/accounts`
##
A entidade `Account` possui os seguintes atributos:

* `id`
* `name` 
* `cpf`
* `secret`
* `balance` 
* `created_at` 

Espera-se as seguintes ações:

- `GET /accounts` - obtém a lista de contas
- `GET /accounts/{account_id}/balance` - obtém o saldo da conta
- `POST /accounts` - cria uma `Account`

*Regras para esta rota*

- `balance` pode iniciar com 0 ou algum valor para simplificar
- `secret` deve ser armazenado como hash


##### `/login`
##

A entidade `Login` possui os seguintes atributos:

* `cpf`
* `secret`

Espera-se as seguintes ações:

- `POST /login` - autentica a usuaria

*Regras para esta rota*

- Deve retornar token para ser usado nas rotas autenticadas


##### `/transfers`
##

A entidade `Transfer` possui os seguintes atributos:

* `id`
* `account_origin_id`
* `account_destination_id`
* `amount`
* `created_at`

Espera-se as seguintes ações:

- `GET /transfers` - obtém a lista de transferencias da usuaria autenticada.
- `POST /transfers` - faz transferencia de uma `Account` para outra.

*Regras para esta rota*

- Quem fizer a transferência precisa estar autenticada.
- O `account_origin_id` deve ser obtido no Token enviado.
- Caso `Account` de origem não tenha saldo, retornar um código de erro apropriado
- Atualizar o `balance` das contas

*Ponto de Atenção*

- [Floating Point Math](https://0.30000000000000004.com/)

=====================

### Requisitos Técnicos

- O código do desafio estar na linguagem [Go](https://golang.org/)
- Pode-se utilizar qualquer package ou framework, mas lembre-se: stdlib > packages > frameworks (usar standard library do Go é melhor que usar packages, que é melhor que usar frameworks...)
- Utilização de Docker é obrigatório

[Topo](#ancora)
* * *
