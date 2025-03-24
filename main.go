package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"indication/internal/config"
	"indication/internal/repository"

	"indication/internal/transport/rest"
	"indication/pkg/database"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"

	log "indication/pkg/logging"

	_ "indication/docs"

	"github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR      = "./configs"
	CONFIG_FILE     = "main"
	CONFIG_LOG_FILE = "log_config.json"
)

func init() {
	// Add security definition using Swagger annotations in your code
	// @securityDefinitions.basic BasicAuth
	err := log.InitLogger(CONFIG_LOG_FILE)
	if err != nil {
		log.LogFatal("Faild to initialize logger", logrus.Fields{"Fatal:": err.Error()})
	}

}

// Swagger
//
//	@title                       Pet Indication Management API
//	@version                     1.0
//	@description                 A comprehensive API for managing pet project indications
func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.LogFatal("Failed to read config", logrus.Fields{"Fatal:": err.Error()})
	}

	dbx, err := database.SqlxPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.User,
		Password: cfg.DB.Pass,
		DBName:   cfg.DB.Db,
		SSLMode:  cfg.DB.SSLMode,
	})

	if err != nil {
		log.LogFatal("Error connecting to database", logrus.Fields{"Error:": err.Error()})
	}
	defer dbx.Close()

	userRepo := repository.NewUsers(dbx)
	lsService := repository.NewNumLSRepository(dbx)
	handler := rest.NewHandler(lsService, userRepo)

	//int & run server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	//log.WithField("port", server.Addr).Info("SERVER STARTED")
	log.LogInfo("SERVER STARTED ON", logrus.Fields{"port": server.Addr, "server": cfg.Server.Name})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.LogFatal("Error starting server", logrus.Fields{"error": err.Error()})
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.LogInfo("Shutting down server...", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.LogFatal("Server forced to shutdown", logrus.Fields{"error": err.Error()})
	}
	log.LogInfo("Server stopped", nil)
}
