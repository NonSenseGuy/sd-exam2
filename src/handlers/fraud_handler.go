package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nonsenseguy/sd-exam2/errors"
	"github.com/nonsenseguy/sd-exam2/models"
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
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrValidEventIDIsRequired)
		return
	}

	item, err := h.store.Get(r.Context(), &models.IDRequest{ID: id})
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &models.FraudataResponseWriter{Item: item})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	limit, err := IntFromString(w, values.Get("limit"))
	if err != nil {
		return
	}

	list, err := h.store.List(r.Context(), &models.ListRequest{Limit: limit})
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &models.FraudataResponseWrapper{Items: list})
}

func (h *handler) Report(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, errors.ErrUnprocessableEntity)
		return
	}

	item := &models.FraudataItem{}
	if err := Unmarshal(w, data, item); err != nil {
		return
	}

	err = h.store.Report(r.Context(), &models.ReportRequest{Item: item})
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &models.FraudataResponseWrapper{Item: item})
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, errors.ErrUnprocessableEntity)
	}

	req := &models.ReportRequest{}
	if err := Unmarshal(w, data, req); err != nil {
		return
	}

	if _, err := h.store.Get(r.Context(), &models.IDRequest{ID: req.Item.ID}); err != nil {
		WriteError(w, err)
		return
	}

	if err = h.store.Update(r.Context(), req); err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &models.FraudataResponseWrapper{})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrValidEventIDIsRequired)
		return
	}

	if _, err := h.store.Get(r.Context(), &models.IDRequest{ID: id}); err != nil {
		WriteError(w, err)
		return
	}

	if err := h.store.Delete(r.Context(), &models.IDRequest{ID: id}); err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &models.FraudataResponseWrapper{})
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
