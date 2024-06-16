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

## Executando a Aplicação com Docker Compose

Se prefere utilizar o Docker Compose para rodar tanto a aplicação quanto o Redis, siga os seguintes passos:

1. Clone o repositório:

```bash
$ git clone <repositório-url>
$ cd project
```

2. Crie um Dockerfile para a aplicação Go:

```Dockerfile
# Usando a imagem base oficial do Go
FROM golang:1.22.4-alpine

# Criando diretório de trabalho
WORKDIR /app

# Copiando arquivos go e go.mod para diretório de trabalho
COPY . .

# Instalando dependências
RUN go mod tidy

# Build da aplicação
RUN go build -o main .

# Comando para rodar a aplicação
CMD ["./main"]
```

3. Adicione o serviço da aplicação ao arquivo `docker-compose.yaml`:

```yaml
version: '3.8'

services:
  redis:
    image: redis:alpine
    container_name: redis
    ports:
      - "6379:6379"
    networks:
      - app-network

  app:
    build: .
    container_name: go_app
    environment:
      - REDIS_ADDR=redis:6379
    ports:
      - "8080:8080"
    depends_on:
      - redis
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
```

4. Suba e execute os containers usando o Docker Compose:

```bash
$ docker-compose up --build
```

5. A aplicação estará disponível em `http://localhost:8080/objects`.

6. Quando quiser parar os containers, execute:

```bash
$ docker-compose down
```

## Observação

Este projeto foi totalmente gerado pelo ChatGPT, uma IA desenvolvida pela OpenAI, com base em suas especificações. Se tiver algum feedback ou melhorias a serem feitas, por favor, abra uma issue ou contribua diretamente com o projeto.
