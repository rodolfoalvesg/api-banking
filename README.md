# Desafio Técnico - Go(lang) - api-banking 
Desafio stone em Go para criação de api baking.

<a name="ancora"></a>
# Desafio e Instrunções
- [Desafio Concluído](#ancora1)
- [Instruções e Regras](#ancora2)


<a id="ancora1"></a>
## Documentação do Desafio
**API de transferencia entre contas Internas de um banco digital.**
#### Rotas e exemplos
##### Criação de Contas: 
- `POST \accounts`
- body:
```json
{
    "name": "My Name",
    "cpf": "1234567800",
    "secret": "12345678",
    "balance": 2000
}
```
- Especificação: campos `name`, `cpf` e `secret` obrigatórios. O campo `balance` poderá ser omitido, porém receberá como valor inicial, 0 (zero).
- Casos de respostas
  * Sucesso:
     + status code: `201` 
     + content-type: `application/json`
     + body:
    ```json
      "cc5c3af1-e1ba-47ee-94b5-bd7fda44999b"
    ```
  * Falhas:
     1. Id
  

[Topo](#ancora)

* * * *
* * * *
<a id="ancora2"></a>


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
