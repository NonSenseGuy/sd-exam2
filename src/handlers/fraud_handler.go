package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nonsenseguy/sd-exam2/store"
)

type IFraudataHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Report(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	store store.IStore
}

func NewFraudataHandler(store store.IStore) IFraudataHandler {
	return &handler{store: store}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Report(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
}
