package tennis

import "fmt"

const (
	deuceGameScore = 6
)

// Match struct contains the Players and GameWinner
type Match struct {
	PlayerOne  Player
	PlayerTwo  Player
	GameWinner string
}

// Player struct contains name and current game score
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

// GetScore fetches the current game's score
func (m *Match) GetScore() string {
	if m.gameHasAWinner() {
		return fmt.Sprintf("Game won by %s", m.GameWinner)
	}
	return m.calculateScore()
}

// PlayerOneWinsAPoint increments player one's score
func (m *Match) PlayerOneWinsAPoint() {
	m.PlayerOne.Score++
}

// PlayerTwoWinsAPoint increments player two's score
func (m *Match) PlayerTwoWinsAPoint() {
	m.PlayerTwo.Score++
}

func (m *Match) calculateScore() string {
	if m.notADeuceGame() {
		return fmt.Sprintf("%s-%s", m.translateScore(m.PlayerOne.Score), m.translateScore(m.PlayerTwo.Score))
	}
	return m.calculateDeuceScenario()
}

func (m *Match) notADeuceGame() bool {
	return m.PlayerOne.Score+m.PlayerTwo.Score <= deuceGameScore
}

func (m *Match) calculateDeuceScenario() string {
	if m.PlayerOne.Score == 4 && m.PlayerTwo.Score == 3 {
		return fmt.Sprintf("%s-%s", "A", "40")
	} else if m.PlayerOne.Score == 3 && m.PlayerTwo.Score == 4 {
		return fmt.Sprintf("%s-%s", "40", "A")
	}
	return "40-40"
}

func (m *Match) translateScore(score int) string {
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

func (m *Match) gameHasAWinner() bool {
	if m.PlayerOne.Score >= 4 && m.PlayerTwo.Score+2 <= m.PlayerOne.Score {
		m.GameWinner = m.PlayerOne.Name
		return true
	} else if m.PlayerTwo.Score >= 4 && m.PlayerOne.Score+2 <= m.PlayerTwo.Score {
		m.GameWinner = m.PlayerTwo.Name
		return true
	}
	return false
}
