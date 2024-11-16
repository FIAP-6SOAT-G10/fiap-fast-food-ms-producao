package main

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/infra/ctx"
	"fiap-fast-food-ms-producao/infra/db"
	"fiap-fast-food-ms-producao/main/producer"
	"fiap-fast-food-ms-producao/main/router"
	"fiap-fast-food-ms-producao/main/worker"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartRouter(ctx context_manager.ContextManager, dbManager database.DatabaseManger) {
	productionUpdateChannel := make(chan []byte)
	defer func(dbManager database.DatabaseManger) {
		err := dbManager.Disconnect()
		if err != nil {
			log.Fatalf("Err Disconnect from MongoDB")
		}
	}(dbManager)
	router := router.InitRouter(ctx, dbManager, productionUpdateChannel)
	go producer.ProductionOrderUpdateProducer(ctx, productionUpdateChannel)
	port := ctx.Get("port")
	router.Run(fmt.Sprintf(":%v", port))
}

func StartWorker(ctx context_manager.ContextManager, dbManager database.DatabaseManger) {
	productionOrderChannel := make(chan map[string]interface{})
	go worker.InitWorker(ctx, productionOrderChannel)
	go worker.ProductionOrderConsumer(ctx, dbManager, productionOrderChannel)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	ctx := ctx.NewContextManager()
	mongoClient, err := db.NewDatabaseManager(ctx)
	if err != nil {
		return
	}

	go StartRouter(ctx, mongoClient)
	go StartWorker(ctx, mongoClient)
	<-quit
	fmt.Println("Shutting down")
}
