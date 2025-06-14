package handlers

import (
    "encoding/json"
    "net/http"
    "fraud-analysis/models"
    "fraud-analysis/services"
)

func AnalisePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var analisePostBody models.AnalisePostBody
    err := json.NewDecoder(r.Body).Decode(&analisePostBody)
    if err != nil {
        http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
        return
    }

    responseOA , err := services.VerifyFraud(analisePostBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusContinue)
    w.Write([]byte(responseOA))
}
