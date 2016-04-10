package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	breakfast_meals := []string{
		"Uppithu",
		"Avalakki",
		"Rave Idli",
		"Shavige",
		"Raagi Dose",
		"Omelette",
		"Chitranna (Lemon rice)",
		"Bread toast",
		"Fruits Salad",
	}
	fmt.Println("I suggest the breakfast below:\n", breakfast_meals[r.Intn(len(breakfast_meals))])
}
