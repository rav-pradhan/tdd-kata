package bowling

// The game consists of 10 frames as shown above. In each frame the player has two opportunities to knock down 10 pins. The score for the frame is the total number of pins knocked down, plus bonuses for strikes and spares.

// A spare is when the player knocks down all 10 pins in two tries. The bonus for that frame is the number of pins knocked down by the next roll. So in frame 3 above, the score is 10 (the total number knocked down) plus a bonus of 5 (the number of pins knocked down on the next roll.)

// A strike is when the player knocks down all 10 pins on his first try. The bonus for that frame is the value of the next two balls rolled.

// In the tenth frame a player who rolls a spare or strike is allowed to roll the extra balls to complete the frame. However no more than three balls can be rolled in tenth frame.

const (
	totalPins  = 10
	finalFrame = 9
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
	for frame := range g.Frames {
		if frame == finalFrame {
			g.calculateFinalFrameRolls(rolls)
			return
		}
		if g.Frames[frame].Played == false {
			g.addRollsToNextFrame(frame, rolls)
			return
		}
	}
}

func (g *Game) calculateFinalFrameRolls(rolls []int) {
	g.Frames[finalFrame].roll1 = rolls[0]
	g.Frames[finalFrame].roll2 = rolls[1]
	if len(rolls) == 3 {
		g.Frames[finalFrame].roll3 = rolls[2]
	}
	g.Frames[finalFrame].Played = true

}

func (g *Game) addRollsToNextFrame(frame int, rolls []int) {
	g.Frames[frame].roll1 = rolls[0]
	g.Frames[frame].roll2 = rolls[1]
	g.Frames[frame].Played = true
}

func (g *Game) GetScore() int {
	return g.calculateTotalScore()
}

func (g *Game) calculateTotalScore() int {
	for i := range g.Frames {
		if g.frameHadStrike(i) {
			g.score += (totalPins + g.calculateStrikeBonusScore(i))
			continue
		}

		if g.frameHadSpare(i) {
			g.score += (totalPins + g.calculateSpareBonusScore(i))
			continue
		}

		g.score += (g.Frames[i].roll1 + g.Frames[i].roll2)
	}
	return g.score
}

func (g *Game) frameHadStrike(frameIndex int) bool {
	return g.Frames[frameIndex].roll1 == totalPins
}

func (g *Game) frameHadSpare(frameIndex int) bool {
	return g.Frames[frameIndex].roll1+g.Frames[frameIndex].roll2 == totalPins
}

func (g *Game) calculateStrikeBonusScore(frameIndex int) int {
	if frameIndex < finalFrame {
		return g.calculateMidFrameStrikeBonus(frameIndex + 1)
	}
	return g.Frames[finalFrame].roll2 + g.Frames[finalFrame].roll3
}

func (g *Game) calculateMidFrameStrikeBonus(nextFrameIndex int) int {
	if g.Frames[nextFrameIndex].roll1 == totalPins && nextFrameIndex+1 <= finalFrame {
		return g.Frames[nextFrameIndex].roll1 + g.Frames[nextFrameIndex+1].roll1
	}
	return g.Frames[nextFrameIndex].roll1 + g.Frames[nextFrameIndex].roll2
}

func (g *Game) calculateSpareBonusScore(frameIndex int) int {
	if frameIndex == finalFrame {
		return g.Frames[finalFrame].roll3
	}
	return g.Frames[frameIndex+1].roll1
}
