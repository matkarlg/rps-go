package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playRound() {
	fmt.Print("Rock, Paper, Scissors? ")
	player := getHand()
	computer := hand(rand.Intn(3))
	result := playHands(player, computer)

	fmt.Println(show[player] + " vs " + show[computer])
	fmt.Println("You " + show[result] + "!")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		playRound()
	}
}
