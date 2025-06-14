package repositories

import (
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "os"
    "fraud-analysis/models"
)

var apiKey string

func init() {
    apiKey = os.Getenv("OPENAI_API_KEY")
}

func CheckPossibilityFraud(requestBody models.RequestBodyOA) (string, error) {
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return "", errors.New("Erro ao serializar JSON: " + err.Error())
    }

    req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
    if err != nil {
        return "", errors.New("Erro ao criar requisição: " + err.Error())
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + apiKey)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", errors.New("Erro ao fazer requisição HTTP: " + err.Error())
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", errors.New("Erro ao ler resposta: " + err.Error())
    }

    var result models.ResponseBodyOA
    err = json.Unmarshal(body, &result)
    if err != nil {
        return "", errors.New("Erro ao decodificar JSON de resposta: " + err.Error())
    }

    return result.Choices[0].Message.Content,  nil
}