package transport

import (
	"CarFix/internal/config"
	"CarFix/internal/database"
	"net/http"
)

func StartServer(db *db.Database, cfg config.Config) {
	mux := http.NewServeMux()

	clientHandler := &ClientHandler{DB: db}
	mux.HandleFunc("GET /clients/", clientHandler.ListClients)
	mux.HandleFunc("GET /clients/{id}/", clientHandler.GetClient)
	mux.HandleFunc("POST /clients/", clientHandler.CreateClient)
	mux.HandleFunc("PUT /clients/", clientHandler.UpdateClient)
	mux.HandleFunc("DELETE /clients/{id}/", clientHandler.DeleteClient)

	http.ListenAndServe(cfg.Address, mux)
}