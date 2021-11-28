info: essa implementação ficou um pouco supeficial, e reflete muito o que está escrito nos livros. 

Uma implementação mais fluida, sem se preocupar tanto em seguir as regras, mas seguir a ideia da arquitetura foi aplicada nesse [fork](https://github.com/Jacquin-Home/hexagonal-arch-bank-example).

--

# Hexagonal architecture bank account example

Esse repositório contém um pouco dos meus estudos sobre arquitetura hexagonal.

Go Lang é a linguagem utilizada.

## Estrutura

```
├── cmd
│   ├── cli
│   │   └── main.go
│   └── httpserver
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── core
│   │   ├── domain
│   │   │   ├── bank
│   │   │   │   ├── account.go
│   │   │   │   └── account_test.go
│   │   │   └── health
│   │   │       └── health.go
│   │   ├── ports
│   │   │   ├── databaseRepository.go
│   │   │   └── services.go
│   │   └── services
│   │       ├── bank-service.go
│   │       ├── bank-service_test.go
│   │       └── health-service.go
│   ├── handlers
│   │   ├── bank-handler.go
│   │   └── health-handler.go
│   └── repositories
│       ├── memory-db.go
│       └── memory-db_test.go
```
