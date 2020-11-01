package server

import (
	"fmt"
	"strings"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const defaultPort = "8090"

type handlerFunc func(ctx *fasthttp.RequestCtx)
type middlewareFunc func(next fasthttp.RequestHandler) fasthttp.RequestHandler

// Server ...
type Server struct {
	Handlers []*Handlers
}

// Handlers ...
type Handlers struct {
	Handler       func(ctx *fasthttp.RequestCtx)
	Middleware    func(next fasthttp.RequestHandler) fasthttp.RequestHandler
	Pattern       string
	HasMiddleware bool
	Method        string
}

// AddHandler ....
func (s *Server) AddHandler(pattern string, method string, handler handlerFunc) {
	var exist bool

	for _, h := range s.Handlers {
		if h.Pattern == pattern && h.Method == method {
			fmt.Println("The route", method, ": ", pattern, "already exist, change the name")
			exist = true
			break
		}
	}

	if !exist {
		s.Handlers = append(s.Handlers, &Handlers{Handler: handler, Pattern: pattern, Method: method, HasMiddleware: false})
	}

}

// Start ...
func (s *Server) Start(port string) {

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	router := router.New()

	for _, handler := range s.Handlers {
		if handler.HasMiddleware {
			router.Handle(handler.Method, handler.Pattern, handler.Middleware(handler.Handler))
		} else {
			router.Handle(handler.Method, handler.Pattern, handler.Handler)
		}
	}

	fmt.Println("Http server running at ", port)

	server := fasthttp.Server{
		MaxRequestBodySize: 5 * 1024 * 1024,
		Name:               "mw-http-server",
		ReadBufferSize:     4096 * 3,
		Handler:            router.Handler,
	}

	if err := server.ListenAndServe(":" + port); err != nil {
		panic("error starting up the http server")
	}
}
