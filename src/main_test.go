package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nonsenseguy/sd-exam2/handlers"
	"github.com/nonsenseguy/sd-exam2/models"
	"github.com/nonsenseguy/sd-exam2/store"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	router    *mux.Router
	flushAll  func(t *testing.T)
	createOne func(t *testing.T, name string) *models.FraudataItem
	getOne    func(t *testing.T, id string, wantErr bool) *models.FraudataItem
)

func TestMain(t *testing.M) {
	log.Println("Registering")

	conn := "postgres://user:password@localhost:5432/db?sslmode=disable"
	if c := os.Getenv("DB_CONN"); c != "" {
		conn = c
	}

	router = mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	st := store.NewPostgresConnection(conn)
	hand := handlers.NewFraudataHandler(st)
	RegisterAllRoutes(router, hand)

	flushAll = func(t *testing.T) {
		db, err := gorm.Open(postgres.Open(conn), nil)
		if err != nil {
			t.Fatal(err)
		}

		db.Delete(&models.IDRequest{}, "1=1")
	}

	createOne = func(t *testing.T, name string) *models.FraudataItem {
		item := &models.FraudataItem{
			Name:          name,
			IsReported:    false,
			ReportReasons: "",
		}

		err := st.Report(context.TODO(), &models.ReportRequest{Item: item})
		if err != nil {
			t.Fatal(err)
		}

		return item
	}

	getOne = func(t *testing.T, id string, wantErr bool) *models.FraudataItem {
		item, err := st.Get(context.TODO(), &models.IDRequest{ID: id})
		if err != nil && wantErr {
			t.Fatal(err)
		}

		return item
	}

	log.Println("Starting")
	os.Exit(t.Run())
}

func Do(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestUnknownEndpoints(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T) *http.Request
	}{
		{
			name: "root",
			setup: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/", nil)
				if err != nil {
					t.Fatal(err)
				}

				return req
			},
		},
		{
			name: "random",
			setup: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/random", nil)
				if err != nil {
					t.Fatal(err)
				}

				return req
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Do(tt.setup(t))
			_ = assert.Equal(t, http.StatusNotFound, w.Code) && assert.Equal(t, "404 page not found\n", w.Body.String())
		})
	}
}

func TestGetEndpoint(t *testing.T) {
	flushAll(t)
	tests := []struct {
		name  string
		code  int
		setup func(t *testing.T) *http.Request
	}{
		{
			name: "OK",
			setup: func(t *testing.T) *http.Request {
				item := createOne(t, "Ok")
				req, err := http.NewRequest(http.MethodGet, "/api/v1/fraudata/item?id="+item.ID, nil)
				if err != nil {
					t.Fatal(err)
				}

				return req
			},
			code: http.StatusOK,
		},
		{
			name: "NotFound",
			setup: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/api/v1/fraudata/item?id=3123122", nil)
				if err != nil {
					t.Fatal(err)
				}

				return req
			},
			code: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Do(tt.setup(t))
			assert.Equal(t, tt.code, w.Code)
			got := &models.FraudataResponseWrapper{}
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), got))
		})
	}
}
