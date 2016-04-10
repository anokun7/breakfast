package main

import (
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
	terminal.Println("I suggest the breakfast below:")
	terminal.MoveCursor(8, 20)
	terminal.Println(terminal.Background(terminal.Color(terminal.Bold(breakfast_meals[r.Intn(len(breakfast_meals))]), terminal.RED), terminal.BLACK))
	terminal.MoveCursor(10, 1)
	terminal.Flush()
}
