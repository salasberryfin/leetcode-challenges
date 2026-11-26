package main

func canAliceWin(n int) bool {
	// Alice 10 stones first turn
	// Each new turn, one less stone than the previous
	// Whoever cannot make a move loses the game
	// result = true if Alice wins the game
	current := n
	isAliceTurn := true

	// Alice starts the game removing 10 stones, and then each turn removes one less stone
	for i := 10; i > 0; i-- {
		if i > current {
			// this player loses
			return !isAliceTurn
		}
		current = current - i
		isAliceTurn = !isAliceTurn
	}

	return true
}
