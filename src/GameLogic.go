package main

import (
	"fmt"
	"strings"
)

func playHands(player, computer hand) result {
	switch (3 + computer - player) % 3 {
	case 1:
		return lost
	case 2:
		return won
	default:
		return draw
	}
}

func getHand() hand {
	for {
		if hand, ok := read[strings.ToLower(readLn())]; ok {
			return hand
		}

		fmt.Println("Bad spelling, try again.")
	}
}
