package main

type Game struct {
	broad      [9]string
	player     string
	turnNumber int
}

func main() {
	var game game
	game.player = "X"

	gameOver := false
	var winner string

	for gameOver != true {
		PrintBoard(game.board)
		move := askforplay()
		err := game.play(move)

	}
}
