package transport

import (
	"CarFix/internal/database"
	"CarFix/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"unicode"
)

type ClientHandler struct {
	http.Handler
	DB *db.Database
}

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidLastName = errors.New("invalid last name")
)

func (h ClientHandler) validateReq(client *models.Client) error {
	if client.Name == "" {
		return ErrInvalidName
	}
	for _, v := range client.Name {
		if !unicode.Is(unicode.Latin, v) {
			return ErrInvalidName
		}
	}

	if client.LastName == "" {
		return ErrInvalidLastName
	}
	for _, v := range client.LastName {
		if !unicode.Is(unicode.Latin, v) {
			return ErrInvalidLastName
		}
	}

	return nil
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Printf("can't parse body of request: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.validateReq(client)
	if err != nil {
		log.Printf("bad request: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DB.CreateNewClient(h.DB.Ctx, client)
	if err != nil {
		log.Printf("can't create client: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Printf("can't parse body of request: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.validateReq(client)
	if err != nil {
		log.Printf("bad request: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DB.EditClient(h.DB.Ctx, client)
	if err != nil {
		log.Printf("can't edit client: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("can't parse id: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DB.DeleteClient(h.DB.Ctx, id)
	if err != nil {
		log.Printf("can't delete client: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ClientHandler) ListClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.DB.GetAllClients(h.DB.Ctx)
	if err != nil {
		log.Printf("can't get clients: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalled, err := json.MarshalIndent(clients, "", "	")
	if err != nil {
		log.Printf("marshalling clients: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}

func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("can't  parse id: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clients, err := h.DB.GetClientById(h.DB.Ctx, id)
	if err != nil {
		log.Printf("can't get client: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(clients) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	marshalled, err := json.MarshalIndent(clients[0], "", "    ")
	if err != nil {
		log.Printf("marshalling client: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}
