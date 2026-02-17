<p align="center">
  <img src="media/smiley-flights-logo.png" width="300" alt="Repository logo" />
</p>
<h3 align="center">Smiley Flights</h3>
<p align="center">A Go-based API to search for the cheapest flights on the Smiles platform using miles.<p>
<p align="center">
    <img src="https://img.shields.io/github/repo-size/lhbelfanti/smiley-flights?label=Repo%20size" alt="Repo size" />
    <img src="https://img.shields.io/github/license/lhbelfanti/smiley-flights?label=License" alt="License" />
</p>

---

# Smiley Flights

Smiley Flights is a specialized tool designed to query the Smiles API and find the best flight deals using miles. It provides a structured API that handles complex date-range searches and processes results to identify the cheapest options for both departure and return legs.

## Getting Started

### Prerequisites

- [Go 1.25](https://go.dev/) or higher.

### Environment Setup

1. Clone the repository.
2. Create a `.env` file in the project root by copying the [.env.example](.env.example):

```bash
cp .env.example .env
```

3. Configure your environment variables in `.env`:

```ini
API_PORT=8080
SMILES_API_KEY=aJqPU7xNHl9qN3NVZnPaJ208aPo2Bh2p2ZV844tw
SMILES_AUTHORIZATION=Bearer <your_bearer_token_here>
```

> **Note**: The `SMILES_AUTHORIZATION` token is a Bearer token that can be obtained from the network logs of a browser session on the Smiles website.

## How to Run

To start the API server locally:

```bash
go run cmd/api/main.go
```

The server will be available at `http://localhost:8080`.

## Postman Collection

You can find the Postman collection for this project in the [postman/smiley-flights-collection.json](postman/smiley-flights-collection.json) file.

---

## License

[MIT](LICENSE)

## Logo License

Generated with AI
