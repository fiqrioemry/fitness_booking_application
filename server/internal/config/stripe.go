package config

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v75"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	log.Println("Stripe client initialized")
}
