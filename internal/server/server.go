package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"trello-clone-backend/internal/database"
	"trello-clone-backend/internal/handlers"
	"trello-clone-backend/internal/routes"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	h := handlers.InitAllHandlers()
	return routes.Routes(h)
}

func NewServer() *http.Server {

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
