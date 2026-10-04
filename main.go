package main

import (
	"projeto-teste/config"
	"projeto-teste/handler"
	"projeto-teste/repository"
	"projeto-teste/router"
	"projeto-teste/service"
)

func main() {
	logger := config.InitLogger("main")

	db, err := config.InitDatabase()

	if err != nil {
		logger.Error(err)
	}

	// contextKey := "ctx_key"
	// tm := config.InitTransactionManager(db, &contextKey)

	engine := config.InitGin()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	auditRepo := repository.NewAuditRepository(db)
	auditService := service.NewAuditService(auditRepo)
	auditHandler := handler.NewAuditHandler(auditService)


	
	router.InitProductRouter(engine, productHandler)
	router.InitAuditRouter(engine, auditHandler)
	
	engine.Run(":8080")
}
