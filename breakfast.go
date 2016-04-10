package main

import (
	//"fmt"
	terminal "github.com/buger/goterm"
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
	terminal.Clear()
	terminal.MoveCursor(1, 1)
	terminal.Println("I suggest the breakfast below:\n\n\n\t\t\t", breakfast_meals[r.Intn(len(breakfast_meals))])
	terminal.MoveCursor(10, 1)
	terminal.Flush()
	//	time.Sleep(time.Second)
}
