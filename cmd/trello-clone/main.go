package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"trello-clone-backend/internal/config"
	"trello-clone-backend/pkg/utils"
)

var logger *slog.Logger

type application struct {
	config config.Config
}

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := utils.ConnectToDatabase()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println("database connection established")

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
	}

	fmt.Printf("starting server", "addr", srv.Addr, "env")

	srv.ListenAndServe()

}
