package main

import (
	"fiap-fast-food-ms-producao/main/router"
	"fiap-fast-food-ms-producao/main/worker"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func StartRouter() {
	router := router.InitRouter()
	router.Run(":8080")
}

func StartWorker() {
	worker.InitWorker()
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go StartRouter()
	go StartWorker()
	<-quit
	fmt.Println("Shutting down")
}
