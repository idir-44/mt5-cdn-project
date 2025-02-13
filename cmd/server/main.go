package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/idir-44/mt5-cdn-project/internal/controllers"
	"github.com/idir-44/mt5-cdn-project/internal/repositories"
	"github.com/idir-44/mt5-cdn-project/internal/services"
	"github.com/idir-44/mt5-cdn-project/pkg/database"
	"github.com/idir-44/mt5-cdn-project/pkg/server"
)

func main() {
	srv := server.New(server.Config{Port: 8080})
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to init store: %s", err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)

	v1 := srv.NewGroup("/v1")

	controllers.RegisterHandlers(v1, service)

	data, err := json.MarshalIndent(srv.Router.Routes(), "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal routes: %s", err)
	}

	fmt.Println(string(data))

	for _, route := range srv.Router.Routes() {
		fmt.Printf("Route enregistrée : %s %s\n", route.Method, route.Path)
	}

	srv.Run()
}
