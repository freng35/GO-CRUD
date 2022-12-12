package main

import (
	"app/pkg/book"
	"app/pkg/config"
	database "app/pkg/db"
	"app/pkg/health"
	"app/pkg/user"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, cfg := setupApp()
	if err := database.CreateDBConnection(cfg, logger); err != nil {
		logger.Fatalw("Failed on db connection", "err", err)
	}

	r := setupRoutes()
	srv := setupServer(r, cfg)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalw("Failed on start server", "err", err)
		}
	}()

	shutdown := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals,
			syscall.SIGTERM,
			syscall.SIGINT)
		<-signals
		close(shutdown)
	}()
	<-shutdown

}

func setupApp() (*zap.SugaredLogger, *config.Config) {
	prodLogger, _ := zap.NewProduction()
	defer prodLogger.Sync()
	logger := prodLogger.Sugar()

	cfg, err := config.NewConfig()

	if err != nil {
		logger.Fatalw("Filed on parsing config", "err", err)
	}

	return logger, cfg
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", health.Healthcheck).Methods("GET")
	r.HandleFunc("/api/user", user.CreateUser).Methods("POST")
	r.HandleFunc("/api/books", book.AddBooks).Methods("POST")
	r.HandleFunc("/api/books", book.GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{name}", book.GetBook).Methods("GET")

	return r
}

func setupServer(r *mux.Router, cfg *config.Config) *http.Server {
	srv := &http.Server{
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
		Addr:         cfg.HttpAddress,
	}

	return srv
}
