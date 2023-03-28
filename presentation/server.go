package presentation

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func Serve(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "20")
}
