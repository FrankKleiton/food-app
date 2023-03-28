package http

import "food-app/presentation/adapters"

type Router struct {
	RoutesList map[string]func(adapters.IResponse, adapters.IRequest)
}

func (r *Router) AddRoute(method string, path string, handler func(response adapters.IResponse, request adapters.IRequest)) {
	if r.RoutesList == nil {
		r.RoutesList = make(map[string]func(adapters.IResponse, adapters.IRequest))
	}
	r.RoutesList[method+" "+path] = handler
}

func (r *Router) Route(response adapters.IResponse, request adapters.IRequest) {
	method := request.GetMethod()
	path := request.GetUrl()
	r.RoutesList[method+" "+path](response, request)
}
