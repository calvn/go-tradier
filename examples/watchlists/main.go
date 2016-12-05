package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/calvn/go-tradier/tradier"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	// Load access token from .env
	err := godotenv.Load("examples/watchlists/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TRADIER_ACCESS_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := tradier.NewClient(tc)

	log.Println("Fetching all watchlists")
	watchlists, _, err := client.Watchlists.All()
	if err != nil {
		log.Fatalf("Error fetching order: %s", err)
	}

	fmt.Printf("ID from first watchlist: %+q\n", *(*watchlists)[0].ID)

	payload, err := json.Marshal(watchlists)
	if err != nil {
		log.Fatalf("Error marshaling orders to JSON: %s", err)
	}

	fmt.Println(string(payload))

	log.Println("Fetching default watchlist")
	watchlist, _, err := client.Watchlists.Get("default")
	if err != nil {
		log.Fatalf("Error fetching order: %s", err)
	}

	payload, err = json.Marshal(watchlist)
	if err != nil {
		log.Fatalf("Error marshaling orders to JSON: %s", err)
	}

	fmt.Println(string(payload))
}
