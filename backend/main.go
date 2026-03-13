package main

import (
	"combat-sim/internal/app"
	"combat-sim/internal/transport"
	"log"
	"net/http"
)

func main() {

	store := app.NewCampaignStore()

	service := app.NewCampaignService(store)

	handler := transport.NewHandler(service)

	http.HandleFunc("/campaign/start", handler.StartCampaign)
	http.HandleFunc("/fight/start", handler.StartFight)
	http.HandleFunc("/fight/action", handler.PerformAction)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
