# go-rest-api

A simple REST API implementation built with [Go](https://go.dev/) and [Gin framework](https://gin-gonic.com/). It's built while I'm learning Go and its conventions.

## Running the server

```sh
$ go run server.go
```

Visit http://localhost:8080

## Implementation

### REST API

- `GET /tasks`
- `GET /tasks/:id`
- `POST /tasks`
- `PATCH /tasks/:id`
- `DELETE /tasks/:id`

### Model

```json
{
  "id": "uuid",
  "title": "Buy milk"
}
```

### Storage

Tasks are currently stored in in-memory map.
