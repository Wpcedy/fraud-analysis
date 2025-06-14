package repositories

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "fraud-analysis/models"
)

type AnalysisRepository struct {
    DB *sql.DB
}

func (r *AnalysisRepository) CriarAnalise(cpf string, dados models.AnalisePostBody, resposta string) error {
    dadosJSON, err := json.Marshal(dados)
    if err != nil {
        return fmt.Errorf("erro ao serializar dados enviados: %w", err)
    }

    query := `INSERT INTO analysis (cpf, dados_enviados, resposta_ia) VALUES ($1, $2, $3)`
    _, err = r.DB.Exec(query, cpf, dadosJSON, resposta)
    if err != nil {
        return fmt.Errorf("erro ao inserir análise: %w", err)
    }

    return nil
}

func (r *AnalysisRepository) BuscarHistoricoPorCPF(cpf string) ([]models.HistoricoAnalise, error) {
    query := `SELECT dados_enviados, resposta_ia FROM analysis WHERE cpf = $1 ORDER BY created_at ASC`
    rows, err := r.DB.Query(query, cpf)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar histórico: %w", err)
    }
    defer rows.Close()

    var historico []models.HistoricoAnalise

    for rows.Next() {
        var h models.HistoricoAnalise
        if err := rows.Scan(&h.DadosEnviados, &h.RespostaIA); err != nil {
            return nil, fmt.Errorf("erro ao ler linha: %w", err)
        }
        historico = append(historico, h)
    }

    return historico, nil
}
