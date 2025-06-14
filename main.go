package main

import (
    "fmt"
    "log"
    "net/http"
	"os"
    "fraud-analysis/handlers"
	"fraud-analysis/config"
    "fraud-analysis/db"
)

func main() {
	config.LoadEnv()

    conn, err:= config.ConnectToDB()
    if err != nil {
        log.Fatal("Erro para conectar ao banco:", err)
    }

    err = db.RunMigrations(conn)
    if err != nil {
        log.Fatal("Erro nas migrations:", err)
    }

    http.HandleFunc("/analise", handlers.AnalisePost)
    port := os.Getenv("PORT")
	if port == "" {
    	port = "8080"
    }

    fmt.Println("Servidor rodando na porta " + port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
