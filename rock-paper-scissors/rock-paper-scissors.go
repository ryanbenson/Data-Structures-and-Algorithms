package rockpaperscissors

func getOpponentMove() string {
	return "rock"
}

func isValidMove(move string) bool {
	if move != "rock" && move != "paper" && move != "scissors" {
		return false
	}
	return true
}

func evaluateGame(userMove string, opponentMove string) string {
	return "win"
}
