package models

import "time"

type AnalisePostBody struct {
    CPF                string    `json:"cpf"`
    Valor              float64   `json:"valor"`
    Produto            string    `json:"produto"`
    Cidade             string    `json:"cidade"`
    HorarioSolicitacao time.Time `json:"horarioSolicitacao"`
    QuantidadeLocacoesAnteriores int     `json:"quantidadeLocacoesAnteriores"`
}
