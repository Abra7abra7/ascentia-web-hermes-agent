package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainRoutes(t *testing.T) {
	mux := http.NewServeMux()

	// Registrujeme dummy handlery pre otestovanie routovania
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Domov"))
	})
	mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Služby"))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %v", res.Status)
	}

	res, err = http.Get(ts.URL + "/services")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %v", res.Status)
	}
}
