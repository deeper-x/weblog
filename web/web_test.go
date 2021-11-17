package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSave(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/save?signature=SYSTEMXYZ&message=test", nil)

	save(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestLoad(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/load?signature=SYSTEMXYZ", nil)

	load(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
