# Project ReadMe

## Overview

Este é um projeto de API REST construído em Go (versão 1.22.4), utilizando a arquitetura Hexagonal para manter o código limpo e aderente aos princípios SOLID. O projeto faz uso do Redis como cache e integra-se com uma API externa para buscar dados. Se os dados não forem encontrados no cache, a API faz uma requisição à API externa "https://api.restful-api.dev/objects".

## Funcionalidades

- **Busca de Objetos**: A API possui um endpoint `/objects` que retorna uma lista de objetos.
- **Cache com Redis**: Antes de fazer uma chamada à API externa, a aplicação verifica se os dados solicitados estão no cache Redis. Se estiverem, os dados são retornados do cache; caso contrário, são buscados na API externa e armazenados no cache para futuras requisições.
- **Arquitetura Hexagonal**: O projeto é estruturado em diferentes camadas para facilitar a manutenção e o teste do código, aderindo ao Clean Code e princípios SOLID.

## Estrutura do Projeto

```
project/
├── docker-compose.yaml
├── Makefile
├── main.go
├── internal/
│   ├── handler/
│   │   ├── object_handler.go
│   │   └── object_handler_test.go
│   └── service/
│       ├── object_service.go
│       ├── object_service_impl.go
│       └── object_service_test.go
```

## Configuração do Ambiente

### Pré-requisitos

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/)

### Como Rodar Localmente

1. Clone o repositório:

```bash
$ git clone <repositório-url>
$ cd project
```

2. Suba os containers Docker:

```bash
$ make run
```

3. Rode a aplicação:

```bash
$ make dev
```

A aplicação estará disponível em `http://localhost:8080/objects`.

### Como Executar Testes

Para rodar os testes, use o comando:

```bash
$ make test
```

### Como Parar os Containers

Para parar os containers Docker, use o comando:

```bash
$ make stop
```

## Observação

Este projeto foi totalmente gerado pelo ChatGPT, uma IA desenvolvida pela OpenAI, com base em suas especificações. Se tiver algum feedback ou melhorias a serem feitas, por favor, abra uma issue ou contribua diretamente com o projeto.
