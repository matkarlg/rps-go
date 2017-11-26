package main

type hand int
type result int

const (
	rock hand = iota
	paper
	scissors
	// deriving show read

	won result = iota
	lost
	draw
	// deriving show
)

// Instances

var read = map[string]hand{
	"rock":     rock,
	"paper":    paper,
	"scissors": scissors,
}

var show = [...]string{
	rock:     "Rock",
	paper:    "Paper",
	scissors: "Scissors",

	won:  "Won",
	lost: "Lost",
	draw: "Draw",
}
