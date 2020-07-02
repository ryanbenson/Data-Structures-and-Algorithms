package rockpaperscissors

import "math/rand"

// getOpponentMove returns a random move
func getOpponentMove() string {
	possibleMoves := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(possibleMoves))
	move := possibleMoves[randomIndex]
	return move
}

// isValidMove determines if a move is valid
func isValidMove(move string) bool {
	if move != "rock" && move != "paper" && move != "scissors" {
		return false
	}
	return true
}

// evaluateGame determines the result of the game comparing the user and opponment move
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
