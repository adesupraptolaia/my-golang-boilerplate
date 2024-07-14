package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adesupraptolaia/assetfindr/config"
	"github.com/adesupraptolaia/assetfindr/controller"
	"github.com/adesupraptolaia/assetfindr/controller/asset"
	assetService "github.com/adesupraptolaia/assetfindr/service/asset"
	assetRepository "github.com/adesupraptolaia/assetfindr/service/asset/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.GetConfig()

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUsername,
		cfg.PostgresPassword,
		cfg.PostgresDBName,
		cfg.PostgresPort,
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("error when create ", err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get sqlDB from gorm: %s", err.Error())
		}

		if err := sqlDB.Close(); err != nil {
			log.Fatalf("failed to close database connection: %s", err.Error())
		}

		log.Println("DB disconnected")
	}()

	assetRepo := assetRepository.NewRepository(db)
	assetSvc := assetService.NewService(assetRepo)
	assetCtrl := asset.NewAssetController(assetSvc)
	ctrl := controller.NewRouterController(assetCtrl)

	router := gin.Default()
	controller.RegisterRoute(router, ctrl)

	srv := &http.Server{
		Addr:    ":" + cfg.PORT,
		Handler: router,
	}

	go func() {
		log.Printf("http server listening to %s...\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start the server", err.Error())
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
		log.Fatal("Server Shutdown:", err.Error())
	}

	select {
	case <-ctx.Done():
		log.Printf("timeout of %s\n", timeout)
	default:
		log.Println("Server exiting")
	}
}
