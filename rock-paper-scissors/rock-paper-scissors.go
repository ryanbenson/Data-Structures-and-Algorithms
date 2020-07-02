package rockpaperscissors

import "math/rand"

func getOpponentMove() string {
	possibleMoves := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(possibleMoves))
	move := possibleMoves[randomIndex]
	return move
}

func isValidMove(move string) bool {
	if move != "rock" && move != "paper" && move != "scissors" {
		return false
	}
	return true
}

func evaluateGame(userMove string, opponentMove string) string {
	if userMove == opponentMove {
		return "tie"
	}

	// list out the win scenarios
	if userMove == "rock" && opponentMove == "scissors" {
		return "win"
	}

	if userMove == "paper" && opponentMove == "rock" {
		return "win"
	}

	if userMove == "scissors" && opponentMove == "paper" {
		return "win"
	}

	// anything else means the user lost
	return "lose"
}
