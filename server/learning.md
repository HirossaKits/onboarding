# Learging go api with go-chi and ent.

## Crate Project

```shell
go mod init go-chi-api
go mod tidy
```

## Install go-chi & ent.

```shell
go get -u github.com/go-chi/chi/v5
go get -d entgo.io/ent/cmd/ent
```

## Initialize schema

```shell
go run entgo.io/ent/cmd/ent init Todo
```

```mermaid
erDiagram

user ||--o{ todo : user_id

user {
  number id
  string email
  string password
  string name
  datetime updated_at
  datetime created_at
}

todo {
  number id
  string user_id
  string title
  string content
  datetime updated_at
  datetime created_at
}

```

## Generate Code

Write schema and run below command

```shell
go generate ./ent
```

## Install mysql driver
