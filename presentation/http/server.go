package http

import (
	"food-app/presentation/adapters"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Router adapters.IRouter
}

func (s *Server) Serve(port int) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			HttpResponse: w,
		}
		request := Request{
			HttpRequest: r,
		}
		s.Router.Route(response, &request)
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}
