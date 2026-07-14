package day5

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateThenGet(t *testing.T) {
	router := NewRouter(NewStore())

	// POST /items
	body := bytes.NewBufferString(`{"id":"1","name":"widget"}`)
	req := httptest.NewRequest(http.MethodPost, "/items", body)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want 201", rec.Code)
	}

	// GET /items/1
	req = httptest.NewRequest(http.MethodGet, "/items/1", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("GET status = %d, want 200", rec.Code)
	}
	var got Item
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Name != "widget" {
		t.Errorf("Name = %q, want widget", got.Name)
	}
}

func TestGetMissingIs404(t *testing.T) {
	router := NewRouter(NewStore())
	req := httptest.NewRequest(http.MethodGet, "/items/nope", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rec.Code)
	}
}

func TestPostBadJSONIs400(t *testing.T) {
	router := NewRouter(NewStore())
	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBufferString(`{not json`))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rec.Code)
	}
}
