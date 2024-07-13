package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adesupraptolaia/assetfindr/controller"
	"github.com/adesupraptolaia/assetfindr/controller/asset"
	assetService "github.com/adesupraptolaia/assetfindr/service/asset"
	"github.com/gin-gonic/gin"
)

func main() {
	assetSvc := assetService.NewService()
	assetCtrl := asset.NewAssetController(assetSvc)
	ctrl := controller.NewRouterController(assetCtrl)

	router := gin.Default()
	controller.RegisterRoute(router, ctrl)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start the server", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Printf("timeout of %s\n", timeout)
	default:
		log.Println("Server exiting")
	}
}
