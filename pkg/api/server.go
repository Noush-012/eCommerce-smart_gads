package http

import (
	_ "github.com/Noush-012/Project-eCommerce-smart_gads/cmd/api/docs"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/handler"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/api/routes"

	"github.com/gin-gonic/gin"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler, orderHandler *handler.OrderHandler) *ServerHTTP {

	engine := gin.New()

	// to load views
	// engine.LoadHTMLGlob("views/static/*.html")
	engine.Static("/static", "./views/static")
	// engine.StaticFS("/static", http.Dir("./views/static"))

	// engine.LoadHTMLGlob("views/*.html")
	engine.Use(gin.Logger())

	// Get swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	// Calling routes
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, paymentHandler, orderHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, orderHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run() {
	s.engine.Run(":3000")
}
