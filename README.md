# api-banking
Desafio stone em Go para criação de api baking.

# Desafio Técnico - Go(lang)

**O desafio é criar uma API de transferencia entre contas Internas de um banco digital.**

=====================
#### API

#### Regras gerais

* Usar formato JSON para leitura e escrita. (ex: `GET /accounts/` retorna json, `POST /accounts/ {name: 'james bond'}`)

#### Rotas esperadas

##### `/accounts`

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

* * *
##### `/login`

A entidade `Login` possui os seguintes atributos:

* `cpf`
* `secret`

Espera-se as seguintes ações:

- `POST /login` - autentica a usuaria

*Regras para esta rota*

- Deve retornar token para ser usado nas rotas autenticadas

* * * 

##### `/transfers`

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

_Indicação de material de estudo no final_

## Critérios de Avaliação
O desafio será avaliado através de quatro critérios.

### Entrega
- O resultado final está completo para ser executado?
- O resultado final atende ao que se propõe fazer?
- O resultado final atende totalmente aos requisitos propostos?
- O código possui algum controle de dependências?

### Boas Práticas
- O código está de acordo com o guia de estilo do Go?
- O código está bem estruturado?
- O código está fluente na linguagem?
- O código faz o uso correto de Design Patterns?


### Documentação
- O código foi entregue com um arquivo de README claro de como se guiar?
- O código possui comentários pertinentes?
- O código está em algum controle de versão?
- Os commits são pequenos e consistentes?
- As mensagens de commit são claras?


### Código Limpo
- O código possibilita expansão para novas funcionalidades?
- O código é Don't Repeat Yourself?
- O código é fácil de compreender?
- O código possui testes?


### Material de Estudo

#### Go
- [Lets'Go - WWG Curitiba + Stone para Iniciantes](https://womenwhogocwb.gitbook.io/letsgo/)
- [Effective Go](https://go.dev/doc/effective_go) 
- [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)
- [Curso Aprenda Go (@ellenkorbes) ](https://www.youtube.com/channel/UCxD5EE0H7qOhRr0tIVsOZPQ)
- [Learn Go](https://learn.go.dev)
- [Gophercise](https://gophercises.com/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang)
- [Error Handling](https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html)
- [DDD Lite in Go](https://threedots.tech/post/ddd-lite-in-go-introduction/)
- [Repository Pattern in Go](https://threedots.tech/post/repository-pattern-in-go/)

#### Rest
- [Desing RESTful API's](https://hackernoon.com/restful-api-design-step-by-step-guide-2f2c9f9fcdbf)
- [HTTP Status Code](https://kinsta.com/pt/blog/lista-codigos-status-http/)

#### Boas praticas
- [Boas Práticas na Stone](https://github.com/stone-payments/stoneco-best-practices/blob/master/README_pt.md)
- [Uber-go Guide](https://github.com/uber-go/guide/blob/master/style.md)

#### Outros
- [SOLID](https://www.youtube.com/watch?v=rtmFCcjEgEw)
- [SOLID in GO](https://www.youtube.com/watch?v=AKdvlr-RzEA)
- [Grupo de Estudos de Go (pt-br)](https://www.youtube.com/channel/UCxRoRvJi7NbC2boKAV70t_g)
- [Web Development with Go (samples) -
Jon Calhoun](https://www.youtube.com/playlist?list=PLVEltXlEeWglOJ42pCxf22YVyxkzan3Xg)
- [Go Bootcamp from Gopherguides.tv](https://www.youtube.com/watch?v=22R1PqXvtws)
- [Just for Func](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)
- [Go WEB Examples](https://gowebexamples.com/)
- [Dave Cheney Blog](https://dave.cheney.net/practical-go)
- [Ardan Labs Blog](https://www.ardanlabs.com/blog)

#### Comunidade Go

- [Slack Gopher](https://invite.slack.golangbridge.org/)
- [Telegram](https://t.me/go_br)

### Sugestões

- [Gorilla Mux](https://github.com/gorilla/mux)
- [Negroni](https://github.com/urfave/negroni)
- [Chi](https://github.com/go-chi/chi)
- [sirupsen/log](https://github.com/sirupsen/logrus)
