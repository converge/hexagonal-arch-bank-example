# Hexagonal architecture bank account example

Esse repositório contém um pouco dos meus estudos sobre arquitetura hexagonal.

Go Lang é a linguagem utilizada.

## Estrutura

```
.
├── cmd
│   └── cli
│       └── main.go
├── go.mod
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── bank
│   │   │       └── account.go
│   │   ├── ports
│   │   │   └── databaseRepository.go
│   │   └── service
│   │       └── service.go
│   └── repositories
│       └── memory-db.go
└── README.md
```
