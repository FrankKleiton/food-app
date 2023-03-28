package adapters

type IRouter interface {
	AddRoute(method string, path string, handler func(IResponse, IRequest))
	Route(response IResponse, request IRequest)
}
