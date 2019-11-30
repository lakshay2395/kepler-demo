package routes

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterRoutes(t *testing.T) {
	router := mux.NewRouter()
	RegisterRoutes(router)
	if router == nil {
		t.Error("Failed to register error")
	}
}
