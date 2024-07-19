package transport

import (
	"CarFix/internal/database"
	"CarFix/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	http.Handler
	DB *db.Database
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read body of request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	client := &models.Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		log.Printf("json unmarshalling: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.DB.CreateNewClient(h.DB.Ctx, client)
	if err != nil {
		log.Printf("Create new client: %s", err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read body of request: %s", err.Error())
	}
	client := &models.Client{}

	err = json.Unmarshal(body, &client)
	if err != nil {
		log.Printf("json unmarshalling: %s", err.Error())
	}

	err = h.DB.EditClient(h.DB.Ctx, client)
	if err != nil {
		log.Printf("Edit client: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
	}

	err = h.DB.DeleteClient(h.DB.Ctx, id)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ClientHandler) ListClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.DB.GetAllClients(h.DB.Ctx)

	if err != nil {
		log.Println(err)
	}

	marshalled, err := json.MarshalIndent(clients, "", "	")
	if err != nil {
		log.Printf("marshall error %s", marshalled)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}

func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
	}

	clients, err := h.DB.GetClientById(h.DB.Ctx, id)
	if err != nil {
		log.Println(err)
	}

	if len(clients) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	marshalled, err := json.MarshalIndent(clients[0], "", "    ")
	if err != nil {
		log.Printf("marshall error %s", marshalled)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}
