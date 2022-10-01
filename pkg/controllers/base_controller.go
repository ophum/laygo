package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(router gin.IRouter)
}

type BaseController struct {
	basePath string
	routes   []Route
}

func NewBaseController(basePath string, routes []Route) *BaseController {
	return &BaseController{basePath, routes}
}

func (c *BaseController) RegisterRoutes(router gin.IRouter) {
	r := router.Group(c.basePath)
	for _, route := range c.routes {
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
