package bowling

import (
	"testing"
)

func TestBowlingGame(t *testing.T) {
	// t.Run("A game is initialised with 10 frames and a score of 0", func(t *testing.T) {
	// 	game := NewGame()
	// 	gotFrameCount := len(game.Frames)
	// 	wantFrameCount := 10
	// 	if gotFrameCount != wantFrameCount {
	// 		t.Errorf("got %d, want %d", gotFrameCount, wantFrameCount)
	// 	}

	// 	gotScore := game.GetScore()
	// 	wantScore := 0
	// 	if gotScore != wantScore {
	// 		t.Errorf("got %d, want %d", gotScore, wantScore)
	// 	}
	// })
	// t.Run("A strike gives 10 points including the total of the following two rolls", func(t *testing.T) {
	// 	game := NewGame()
	// 	game.AddFrameScore(10, 0)
	// 	game.AddFrameScore(5, 4)
	// 	got := game.GetScore()
	// 	want := 28
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	// t.Run("A strike's bonus points when it accounts for 2 strikes in a row", func(t *testing.T) {
	// 	game := NewGame()
	// 	game.AddFrameScore(10, 0)
	// 	game.AddFrameScore(10, 0)
	// 	game.AddFrameScore(5, 4)
	// 	got := game.GetScore()
	// 	want := 53
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	// t.Run("Spare bonus score is calculated with the next roll's score", func(t *testing.T) {
	// 	game := NewGame()
	// 	game.AddFrameScore(0, 10)
	// 	game.AddFrameScore(5, 0)
	// 	got := game.GetScore()
	// 	want := 20
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })
	t.Run("Perfect game score should be 300", func(t *testing.T) {
		game := NewGame()
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 0)
		game.AddFrameScore(10, 10, 10)
		got := game.GetScore()
		want := 300
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
