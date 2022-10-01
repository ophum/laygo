package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ophum/laygo/pkg/controllers"
)

func Run(addr string, controllers []controllers.Controller) error {
	r := gin.Default()
	RegisterRoutes(r, controllers)
	return r.Run(addr)
}

func RunTLS(addr, certFile, keyFile string, controllers []controllers.Controller) error {
	r := gin.Default()
	RegisterRoutes(r, controllers)
	return r.RunTLS(addr, certFile, keyFile)
}

func RegisterRoutes(router gin.IRouter, controllers []controllers.Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
}
