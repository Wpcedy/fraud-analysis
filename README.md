# Analise de Fraude

Desafio Analise de Fraude com OpenApi

Este projeto implementa um serviço de análise de probabilidade de fraude utilizando a API da OpenAI. Ele recebe dados de transações e retorna uma probabilidade de fraude, além de um status de recomendação (liberar/bloquear) e um motivo.

# Requisitos
- Go 1.24
- Postgres
- Conta na OpenAI com API Key válida

## Configurando

Duplicar o ".env.example" e preencher os valores.

```env
# GENERAL
PORT=8080

# DATABASE
DB_HOST=localhost
DB_USER=seu_usuario
DB_PASS=sua_senha
DB_NAME=fraud_db

# OPENAI
OPENAI_API_KEY=sua_api_key
MODEL=gpt-4
```

## Executar

```bash
go run main.go
```

## Testes

Usar sua plataforma de API de sua preferência.

- POST /analise

BODY EXAMPLE
```json
{
"cpf": "12345678901",
"valor": 150.0,
"produto": "furadeira",
"cidade": "São Paulo",
"horarioSolicitacao": "2025-06-03T15:00:00Z",
"quantidadeLocacoesAnteriores": 0
}
```

RESPONSE EXAMPLE
```json
{
"probabilidade": 0.82,
"status": "bloquear",
"motivo": "..."
}
```

## Estrutura do projeto

```bash
fraud-analysis/
├── config/         # Arquivos de Configurações
├── handlers/       # Rotas e controladores HTTP
├── models/         # Estruturas de dados (request/response)
├── repositories/   # Integração com OpenAI e banco
├── services/       # Lógica de negócio (análise de fraude)
├── main.go         # Arquivo principal
├── .env.example    # Exemplo de variáveis de ambiente
└── README.md
```
