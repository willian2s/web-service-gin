# Album Service

The Album Service is a web service that allows users to create and search for albums.

## Prerequisites

- [Go](https://golang.org/doc/install)

### Optional

- [Air](https://github.com/cosmtrek/air)
- [Thunder Client](https://www.thunderclient.io/)

## Installation

1. Clone the repository:

```bash
  git clone https://github.com/willian2s/web-service-gin.git
```

2. Install the dependencies:

```bash
  go mod download
```

3. Start the development server:

```bash
  air
```

or

```bash
  go run .
```

## Usage

### List all albums

To list all albums, send a GET request to the `/albums` endpoint:

```bash
  curl http://localhost:8080/albums
```

```json
  {
    "id": "b3997ac1-c367-40e3-8327-223ef777355c",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "60831261-8bd1-4fb3-9df7-99f75a003caf",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  },
```

#### Create an Album

To create an album, send a POST request to the `/albums` endpoint:

```bash
  curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## Contributing

Contributions are welcome! Please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).
