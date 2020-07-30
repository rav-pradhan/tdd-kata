package tennis_test

import (
	"fmt"
	"testing"
)

type Match struct {
	PlayerOne  Player
	PlayerTwo  Player
	GameWinner string
}

func (m *Match) GetScore() string {
	var result string

	if m.PlayerOne.Score+m.PlayerTwo.Score <= 6 {
		playerOneTranslatedScore := m.TranslateScore(m.PlayerOne.Score)
		playerTwoTranslatedScore := m.TranslateScore(m.PlayerTwo.Score)
		result = fmt.Sprintf("%s-%s", playerOneTranslatedScore, playerTwoTranslatedScore)
	} else {
		result = m.calculateDeuceScenario()
	}

	return result
}

func (m *Match) PlayerOneWinsAPoint() {
	m.PlayerOne.Score++
}

func (m *Match) PlayerTwoWinsAPoint() {
	m.PlayerTwo.Score++
}

func (m *Match) calculateDeuceScenario() string {
	if m.PlayerOne.Score == 4 && m.PlayerTwo.Score == 3 {
		return fmt.Sprintf("%s-%s", "A", "40")
	} else if m.PlayerOne.Score == 3 && m.PlayerTwo.Score == 4 {
		return fmt.Sprintf("%s-%s", "40", "A")
	}
	return "40-40"
}

func (m *Match) TranslateScore(score int) string {
	switch score {
	case 0:
		return "0"
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	default:
		return ""
	}
}

type Player struct {
	Name  string
	Score int
}

// New creates an instance of a Match struct
func New(p1, p2 string) Match {
	return Match{
		PlayerOne: Player{
			Name:  p1,
			Score: 0,
		},
		PlayerTwo: Player{
			Name:  p2,
			Score: 0,
		},
	}
}

func TestTennis(t *testing.T) {
	t.Run("Test that a tennis match is created with two players", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")

		if match.PlayerOne.Name != "Rafael Nadal" && match.PlayerTwo.Name != "Andy Murray" {
			t.Errorf("Players not created. Expected %s and %s, got %s and %s", "Rafael Nadal", "Andy Murray", match.PlayerOne.Name, match.PlayerTwo.Name)
		}
	})

	t.Run("Test that the score is 0-0 when a tennis match is created", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")

		if match.GetScore() != "0-0" {
			t.Errorf("Match score not correct. Expected %s and got %s", "0-0", match.GetScore())
		}
	})

	t.Run("Test that the score is 15-0 when player one wins a point", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 1, 0)

		if match.GetScore() != "15-0" {
			t.Errorf("Match score not correct. Expected %s and got %s", "15-0", match.GetScore())
		}
	})

	t.Run("Test that the score is 0-15 when player two wins a point", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 0, 1)

		if match.GetScore() != "0-15" {
			t.Errorf("Match score not correct. Expected %s and got %s", "0-15", match.GetScore())
		}
	})

	t.Run("Test that the score is 30-0 when player one wins two points", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 2, 0)

		if match.GetScore() != "30-0" {
			t.Errorf("Match score not correct. Expected %s and got %s", "30-0", match.GetScore())
		}
	})

	t.Run("Test that player one wins the game when winning a point on 40-0", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 4, 0)

		if match.GetScore() != "Game won by Rafael Nadal" {
			t.Errorf("Incorrect score readout. Expected %s and got %s", "Game won by Rafael Nadal", match.GetScore())
		}
	})

	t.Run("Test that the score is 40-40 when both players win 3 points each", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 3, 3)

		if match.GetScore() != "40-40" {
			t.Errorf("Match score not correct. Expected %s and got %s", "40-40", match.GetScore())
		}
	})

	t.Run("Test that the score is A-40 when both players win 3 points each and player one wins the following point", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 4, 3)

		if match.GetScore() != "A-40" {
			t.Errorf("Match score not correct. Expected %s and got %s", "A-40", match.GetScore())
		}
	})

	t.Run("Test that the score is 40-A when both players win 3 points each and player two wins the following point", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 3, 4)

		if match.GetScore() != "40-A" {
			t.Errorf("Match score not correct. Expected %s and got %s", "40-A", match.GetScore())
		}
	})

	t.Run("Test that the score returns to 40-40 when player one is on Advantage and player two wins the next point", func(t *testing.T) {
		match := New("Rafael Nadal", "Novak Djokovic")
		setScore(&match, 4, 3)
		match.PlayerTwoWinsAPoint()

		if match.GetScore() != "40-40" {
			t.Errorf("Match score not correct. Expected %s and got %s", "40-40", match.GetScore())
		}
	})
}

func setScore(match *Match, p1Score, p2Score int) {
	for i := 0; i < p1Score; i++ {
		match.PlayerOneWinsAPoint()
	}

	for i := 0; i < p2Score; i++ {
		match.PlayerTwoWinsAPoint()
	}
}
