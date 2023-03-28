package presentation

import (
	"food-app/presentation"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	t.Run("Add route", func(t *testing.T) {
		routes := presentation.Routes{}
		routes.AddRoute("GET", "/test", func(response http.ResponseWriter, request *http.Request) {
		})
		if len(routes.RoutesList) != 1 {
			t.Errorf("Expected 1 route, got %d", len(routes.RoutesList))
		}
	})

	t.Run("Run route", func(t *testing.T) {
		routes := presentation.Routes{}
		ran := false
		routes.AddRoute("GET", "/test", func(response http.ResponseWriter, request *http.Request) {
			ran = true
		})

		request, _ := http.NewRequest("GET", "/test", nil)
		response := http.ResponseWriter(nil)
		routes.Route(request, response)

		if ran != true {
			t.Errorf("Expected true route, got %v", ran)
		}
	})
}
