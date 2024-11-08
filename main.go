package main

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/infra/ctx"
	"fiap-fast-food-ms-producao/main/router"
	"fiap-fast-food-ms-producao/main/worker"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func StartRouter(ctx context_manager.ContextManager) {
	router := router.InitRouter()
	port := ctx.Get("port")
	router.Run(fmt.Sprintf(":%v", port))
}

func StartWorker(ctx context_manager.ContextManager) {
	worker.InitWorker(ctx)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	ctx := ctx.NewContextManager()

	go StartRouter(ctx)
	go StartWorker(ctx)
	<-quit
	fmt.Println("Shutting down")
}
