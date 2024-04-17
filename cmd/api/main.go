package main

import (
	"context"
	"fmt"
	cars "helloWRLDs/bookings/internal/cars/delivery/http"
	config "helloWRLDs/bookings/pkg/configs"
	pg "helloWRLDs/bookings/pkg/datastore/postgresql"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		DisableColors:    false,
		TimestampFormat:  "2006-01-02 15:04:05",
		FieldMap: logger.FieldMap{
			logrus.FieldKeyTime:  "@time",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyLevel: "lvl",
		},
	})
	logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// logger.SetOutput(os.Stdout)
	logger.SetOutput(logFile)
}

func main() {
	cfg := config.LoadAppConfig()
	w := logger.New().WriterLevel(logger.ErrorLevel)
	defer w.Close()

	db, err := pg.Open(cfg.Db)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err": err.Error(),
			"dsn": fmt.Sprintf("%s@%s:%s/%s", cfg.Db.Type, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name),
		}).Fatal("db conn error")
		os.Exit(1)
	}
	logger.WithField("dsn", fmt.Sprintf("%s@%s:%s/%s", cfg.Db.Type, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)).Info("db conn ok")

	carsDlvr := cars.NewCarsDelivery(db)

	router := initRouter()
	router.Route("/api/v1", func(router chi.Router) {
		router.Mount("/cars", carsDlvr.Routes())
	})

	srv := &http.Server{
		Addr:     cfg.Addr,
		Handler:  router,
		ErrorLog: log.New(w, "", 0),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.WithField("addr", cfg.Addr).Info("server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}
	logger.Info("Server exiting")
}
