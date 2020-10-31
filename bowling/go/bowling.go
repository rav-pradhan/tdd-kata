package bowling

// The game consists of 10 frames as shown above. In each frame the player has two opportunities to knock down 10 pins. The score for the frame is the total number of pins knocked down, plus bonuses for strikes and spares.

// A spare is when the player knocks down all 10 pins in two tries. The bonus for that frame is the number of pins knocked down by the next roll. So in frame 3 above, the score is 10 (the total number knocked down) plus a bonus of 5 (the number of pins knocked down on the next roll.)

// A strike is when the player knocks down all 10 pins on his first try. The bonus for that frame is the value of the next two balls rolled.

// In the tenth frame a player who rolls a spare or strike is allowed to roll the extra balls to complete the frame. However no more than three balls can be rolled in tenth frame.

const (
	totalPins = 10
)

type Game struct {
	Frames [10]Frame
	score  int
}

type Frame struct {
	roll1  int
	roll2  int
	roll3  int
	Played bool
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) AddFrameScore(rolls ...int) {
	for i := range g.Frames {
		if i == 9 {
			if rolls[0] == totalPins {
				g.Frames[i].roll1 = rolls[0]
				g.Frames[i].roll2 = rolls[1]
				g.Frames[i].roll3 = rolls[2]
			}
		}
		if g.Frames[i].Played == false {
			g.Frames[i].roll1 = rolls[0]
			g.Frames[i].roll2 = rolls[1]
			g.Frames[i].Played = true
			break
		}
	}
}

func (g *Game) GetScore() int {
	g.calculateTotalScore()
	return g.score
}

func (g *Game) calculateTotalScore() {
	for i := range g.Frames {
		if g.frameHadStrike(i) {
			g.score += (totalPins + g.calculateStrikeBonusScore(i+1))
			continue
		}

		if g.frameHadSpare(i) {
			g.score += (totalPins + g.calculateSpareBonusScore(i+1))
			continue
		}

		g.score += (g.Frames[i].roll1 + g.Frames[i].roll2)
	}
}

func (g *Game) frameHadStrike(frameIndex int) bool {
	return g.Frames[frameIndex].roll1 == totalPins
}

func (g *Game) frameHadSpare(frameIndex int) bool {
	return g.Frames[frameIndex].roll1+g.Frames[frameIndex].roll2 == totalPins
}

func (g *Game) calculateStrikeBonusScore(nextFrameIndex int) int {
	if nextFrameIndex < 10 {
		if g.Frames[nextFrameIndex].roll1 == 10 && nextFrameIndex+1 < 11 {
			return g.Frames[nextFrameIndex].roll1 + g.Frames[nextFrameIndex+1].roll1
		}
		return g.Frames[nextFrameIndex].roll1 + g.Frames[nextFrameIndex].roll2
	}

	if nextFrameIndex == 10 {
		return g.Frames[9].roll2 + g.Frames[9].roll3
	}
	return 0
}

func (g *Game) calculateSpareBonusScore(nextFrameIndex int) int {
	if nextFrameIndex < 10 {
		return g.Frames[nextFrameIndex].roll1 + g.Frames[nextFrameIndex+1].roll1
	}
	return 0
}
