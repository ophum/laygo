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
