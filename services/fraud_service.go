package services

import (
    "encoding/json"
    "fmt"
    "fraud-analysis/models"
    "fraud-analysis/repositories"
    "fraud-analysis/config"
    "os"
)

func VerifyFraud(analisePostBody models.AnalisePostBody) (string, error) {
    conn, err:= config.ConnectToDB()
    if err != nil {
		return "", fmt.Errorf("erro ao conectar ao banco: %v", err)
    }

	repo := repositories.AnalysisRepository{DB: conn}
	historico, err := repo.BuscarHistoricoPorCPF(analisePostBody.CPF)
    if err != nil {
		return "",  fmt.Errorf("erro ao buscar histórico: %v", err)
    }

   	jsonBytes, err := json.Marshal(analisePostBody)
	if err != nil {
		return "",  fmt.Errorf("erro ao converter para JSON: %v", err)
	}

	analisePostBodyString := string(jsonBytes)

    prompt := fmt.Sprintf(`
	Com base no histórico de transações abaixo, analise se a nova transação informada é potencialmente fraudulenta.

	Retorne APENAS um JSON com os seguintes campos:
	{
	"probabilidade": decimal entre 0 e 1,
	"status": "aprovado" ou "bloqueado",
	"motivo": string com uma justificativa clara da decisão
	}

	Não inclua nenhuma explicação fora do JSON.

	Histórico de transações analisadas:
	%s

	Nova transação:
	%s
	`, historico, analisePostBodyString)


	requestBody := models.RequestBodyOA{
        Model: os.Getenv("MODEL"),
        Messages: []models.MessageOA{
            {Role: "user", Content: prompt},
        },
    }

    responseOA , err := repositories.CheckPossibilityFraud(requestBody)
    if err != nil {
		return "",  fmt.Errorf("erro ao chamar OpenAI: %v", err)
    }

	err = repo.CriarAnalise(analisePostBody.CPF, analisePostBody, responseOA)
    if err != nil {
		return "",  fmt.Errorf("erro ao salvar análise: %v", err)
    }

	return responseOA, nil
}
