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
	err := godotenv.Load("examples/account_balances/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TRADIER_ACCESS_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := tradier.NewClient(tc)

	order, _, err := client.Order.Create("6YA05708", "equity", "AAPL", "day", "buy", 1, "market", 0, 0, "")
	if err != nil {
		log.Fatalf("Error fetching order: %s", err)
	}

	payload, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Error marshaling orders to JSON: %s", err)
	}

	fmt.Println(string(payload))
}
