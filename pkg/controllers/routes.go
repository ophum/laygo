package controllers

import "github.com/gin-gonic/gin"

type RouteMethod string

const (
	GET    RouteMethod = "GET"
	POST   RouteMethod = "POST"
	PUT    RouteMethod = "PUT"
	PATCH  RouteMethod = "PATCH"
	DELETE RouteMethod = "DELETE"
)

type Route struct {
	Method  RouteMethod
	Path    string
	Handler func(ctx *gin.Context) error
}

type HandlerFunc func(ctx *gin.Context) error
type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) append(method RouteMethod, path string, handler HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.append(GET, path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.append(POST, path, handler)
}

func (r *Router) PUT(path string, handler HandlerFunc) {
	r.append(PUT, path, handler)
}

func (r *Router) PATCH(path string, handler HandlerFunc) {
	r.append(PATCH, path, handler)
}

func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.append(DELETE, path, handler)
}
