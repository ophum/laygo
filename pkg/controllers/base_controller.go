package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(router gin.IRouter)
}

type BaseController struct {
	router *Router
}

func NewBaseController(router *Router) *BaseController {
	return &BaseController{router}
}

func (c *BaseController) RegisterRoutes(router gin.IRouter) {
	r := router.Group(c.router.basePath)
	for _, route := range c.router.routes {
		switch route.Method {
		case GET:
			r.GET(route.Path, handler(route.Handler))
		case POST:
			r.POST(route.Path, handler(route.Handler))
		case PUT:
			r.PUT(route.Path, handler(route.Handler))
		case PATCH:
			r.PATCH(route.Path, handler(route.Handler))
		case DELETE:
			r.DELETE(route.Path, handler(route.Handler))
		}
	}
}

func handler(f func(ctx *gin.Context) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if err := f(ctx); err != nil {
			ctx.Error(err)
		}
	}
}
