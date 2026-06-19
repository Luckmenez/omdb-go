# go-imdb

API REST simples em Go para busca de filmes e séries usando a [OMDb API](https://www.omdbapi.com/).

## Tecnologias

- **Go** — linguagem principal
- **chi** — roteador HTTP leve
- **OMDb API** — fonte dos dados de filmes/séries

## Estrutura

```
.
├── main.go         # entrypoint, configuração do servidor HTTP
├── api/
│   └── handler.go  # roteamento e handlers HTTP
└── omdb/
    └── omdb.go     # client para a OMDb API
```

## Pré-requisitos

- Go 1.22+
- Chave de API gratuita do [OMDb](https://www.omdbapi.com/apikey.aspx)

## Configuração

Exporte sua chave de API como variável de ambiente:

```bash
export OMDB_KEY=sua_chave_aqui
```

## Rodando

```bash
go run main.go
```

O servidor sobe na porta **3333**.

## Endpoints

### `GET /`

Busca filmes ou séries pelo título.

**Query params**

| Parâmetro | Tipo   | Descrição                   |
|-----------|--------|-----------------------------|
| `s`       | string | Título (ou parte do título) |

**Exemplo**

```bash
curl "http://localhost:3333/?s=Inception"
```

**Resposta**

```json
{
  "data": {
    "Search": [
      {
        "Title": "Inception",
        "Year": "2010",
        "imdbID": "tt1375666",
        "Type": "movie",
        "Poster": "https://..."
      }
    ],
    "totalResults": "1",
    "Response": "True"
  }
}
```

Em caso de erro:

```json
{
  "error": "somenthing wrong with omdb"
}
```
