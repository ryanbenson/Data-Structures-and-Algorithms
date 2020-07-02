package rockpaperscissors

import (
	"testing"
)

func TestGetOpponentMove(t *testing.T) {
	move := getOpponentMove()

	if move != "rock" && move != "paper" && move != "scissors" {
		t.Errorf("When getting a random choice, is it not a valid move, got: %v.", move)
	}
}

func TestIsValidMove_IsValid(t *testing.T) {
	move := "scissors"
	isValid := isValidMove(move)

	if isValid != true {
		t.Errorf("When determining if a move is valid, it was incorrect, got: %v, want: %v.", isValid, true)
	}
}

func TestIsValidMove_IsInvalid(t *testing.T) {
	move := "lizard"
	isValid := isValidMove(move)

	if isValid != false {
		t.Errorf("When determining if a move is valid, it was incorrect, got: %v, want: %v.", isValid, false)
	}
}

func TestEvaluateGame_WinRock(t *testing.T) {
	userMove := "rock"
	opponentMove := "scissors"
	result := evaluateGame(userMove, opponentMove)
	expected := "win"

	if result != "win" {
		t.Errorf("When determining a winner, and should win, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_WinPaper(t *testing.T) {
	userMove := "paper"
	opponentMove := "rock"
	result := evaluateGame(userMove, opponentMove)
	expected := "win"

	if result != "win" {
		t.Errorf("When determining a winner, and should win, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_WinScissors(t *testing.T) {
	userMove := "scissors"
	opponentMove := "paper"
	result := evaluateGame(userMove, opponentMove)
	expected := "win"

	if result != "win" {
		t.Errorf("When determining a winner, and should win, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_LosePaper(t *testing.T) {
	userMove := "paper"
	opponentMove := "scissors"
	result := evaluateGame(userMove, opponentMove)
	expected := "lose"

	if result != "lose" {
		t.Errorf("When determining a winner, and should lose, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_LoseScissors(t *testing.T) {
	userMove := "scissors"
	opponentMove := "rock"
	result := evaluateGame(userMove, opponentMove)
	expected := "lose"

	if result != "lose" {
		t.Errorf("When determining a winner, and should lose, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_LoseRock(t *testing.T) {
	userMove := "rock"
	opponentMove := "paper"
	result := evaluateGame(userMove, opponentMove)
	expected := "lose"

	if result != "lose" {
		t.Errorf("When determining a winner, and should lose, it was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestEvaluateGame_Tie(t *testing.T) {
	userMove := "rock"
	opponentMove := "rock"
	result := evaluateGame(userMove, opponentMove)
	expected := "tie"

	if result != "tie" {
		t.Errorf("When determining a winner, and should tie, it was incorrect, got: %v, want: %v.", result, expected)
	}
}
