package scripts

import (
	"image/color"
	"strconv"

	"golang.org/x/image/font/inconsolata"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Hud struct {
	game   *Game
	score  int
	bigger bool
}

// initHud Constructor
func initHud(g *Game) *Hud {
	hud := Hud{
		game:   g,
		score:  0,
		bigger: false,
	}

	return &hud
}

func (hud *Hud) addPoint() {
	hud.score++
}

// Draw the hud
func (hud *Hud) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Health: "+strconv.Itoa(hud.game.crashCount)+" | Score: "+strconv.Itoa(hud.score), inconsolata.Bold8x16, 15, 25, color.Black)
	if hud.game.alive == false {
		textGameOver := ""
		textGameOverDescription := ""
		if hud.game.crashed {
			textGameOver = "GAME OVER"
			textGameOverDescription = "You Crashed Into Something!"
		} else if hud.game.crashCount == 0 {
			textGameOver = "GAME OVER"
			textGameOverDescription = "Your Health dropped to 0!"
		} else if hud.bigger {
			textGameOver = "YOU WIN!"
		} else {
			textGameOver = "GAME OVER"
			textGameOverDescription = "Get More Food Next Time!"

		}
		text.Draw(screen, textGameOver, inconsolata.Bold8x16, 250, 350, color.Black)
		text.Draw(screen, textGameOverDescription, inconsolata.Bold8x16, 180, 380, color.Black)
	}
	return nil
}
