package http

type Router struct {
	handlers map[route]HandleFunc
}

func newRouter() Router {
	return Router{
		handlers: make(map[route]HandleFunc),
	}
}

func (r *Router) getHandler(method string, path string) HandleFunc {
	route := newRoute(method, path)
	return r.handlers[route]
}

func (r *Router) addHandler(method string, path string, handler HandleFunc) {
	route := newRoute(method, path)
	r.handlers[route] = handler
}

type route struct {
	method string
	path   string
}

func newRoute(method string, path string) route {
	return route{
		method: method,
		path:   path,
	}
}

type HandleFunc = func(r *Request) (Response, error)
