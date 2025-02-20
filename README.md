# Golang API

Esta é uma API simples escrita em Go que permite realizar operações CRUD (Create, Read, Update, Delete) em uma lista de pessoas.

## Endpoints

### GET /

Retorna uma mensagem de boas-vindas.

### GET /people

Retorna a lista de pessoas.

### POST /people

Adiciona uma nova pessoa à lista. O corpo da requisição deve conter um JSON com os campos `name`, `age` e `salary`.

Exemplo de corpo da requisição:
```json
{
  "name": "John",
  "age": 30,
  "salary": 1000.0
}
```

### DELETE /people

Remove uma pessoa da lista. O corpo da requisição deve conter um JSON com o campo `id`.

Exemplo de corpo da requisição:
```json
{
  "id": 1
}
```

### PUT /people

Atualiza uma pessoa na lista. O corpo da requisição deve conter um JSON com os campos `id`, `name`, `age`, e `salary`.

Exemplo de corpo da requisição:
```json
{
  "id": 1,
  "name": "John",
  "age": 31,
  "salary": 1100.0
}
```

## Como executar
1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório.
3. Navegue até o diretório do projeto.
4. Execute o comando go run main.go.
5. A API estará disponível em `http://localhost:8080`.