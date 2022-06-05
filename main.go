package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"mission-control-center/controllers"
	"mission-control-center/repository"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	godotenv.Load("./config.env")
	// GET ENV VARIABLES //
	port := os.Getenv("API_PORT")

	dbConfig := repository.Config{
		ConnectionString: os.Getenv("DB_CONNECTIONSTRING"),
		Enabled:          true,
		Port:             os.Getenv("DB_PORT"),
		Database:         os.Getenv("DB_NAME"),
	}

	//INITIALIZE HANDLERS FOR DEPENDENCY INJECTION //
	//Order of initialization matters!!!
	logger := log.New(os.Stdout, "Template Service: ", log.LstdFlags)
	//Setup database and repository
	db := repository.NewPostgresDB(&dbConfig, logger)
	db.ConnectPostgresDB()

	hb := controllers.NewHeartbeat(logger)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// HANDLE ROUTES
	getRouter.HandleFunc("/{heartbeat|healthcheck}", hb.Heartbeat)

	//todo: Fetch from configuration file (MAY NOT BE NECESSARY)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      sm,
		IdleTimeout:  2 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		logger.Println(fmt.Sprintf("Starting Server on port: %s", port))
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
