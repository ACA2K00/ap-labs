package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Position : Struct for knowing the position of the body and head of the snake
type Position struct {
	X float64
	Y float64
}

type Snake struct {
	game         *Game
	len          int
	dir          int
	length       int
	newLength    int
	score        int
	behavior     chan int
	headImgRight ebiten.Image
	headImgLeft  ebiten.Image
	headImgDown  ebiten.Image
	headImgUp    ebiten.Image
	bodyImg      ebiten.Image
	parts        []Position
}

func createSnake(g *Game) *Snake {
	snake := Snake{
		game: g,
		len:  0,
		dir:  1,
	}
	snake.behavior = make(chan int)
	snake.parts = append(snake.parts, Position{300, 300})
	headimgright, _, _ := ebitenutil.NewImageFromFile("img_res/sHeadRight.png", ebiten.FilterLinear)
	headimgleft, _, _ := ebitenutil.NewImageFromFile("img_res/sHeadLeft.png", ebiten.FilterLinear)
	headimgdown, _, _ := ebitenutil.NewImageFromFile("img_res/sHeadDown.png", ebiten.FilterLinear)
	headimgup, _, _ := ebitenutil.NewImageFromFile("img_res/sHeadUp.png", ebiten.FilterLinear)
	bodyimg, _, _ := ebitenutil.NewImageFromFile("img_res/sBody.png", ebiten.FilterLinear)
	snake.headImgRight = *headimgright
	snake.headImgLeft = *headimgleft
	snake.headImgDown = *headimgdown
	snake.headImgUp = *headimgup
	snake.bodyImg = *bodyimg
	return &snake
}

func (snake *Snake) Behavior() error {
	dotTime := <-snake.behavior
	for {
		err := snake.Update(dotTime)
		if err != nil {
			return err
		}
		dotTime = <-snake.behavior
	}
}

func (snake *Snake) Update(dotTime int) error {
	// Left: 1
	// Down: 2
	// Right : 3
	// Up: 4
	if snake.game.alive {
		if ebiten.IsKeyPressed(ebiten.KeyRight) && snake.dir != 3 && snake.dir != 1 {
			snake.dir = 3
			return nil
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && snake.dir != 1 && snake.dir != 3 {
			snake.dir = 1
			return nil
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) && snake.dir != 2 && snake.dir != 5 {
			snake.dir = 2
			return nil
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) && snake.dir != 5 && snake.dir != 2 {
			snake.dir = 4
			return nil
		}

		if dotTime == 1 {
			xPos, yPos := snake.getHeadPos()
			if xPos < 20 || xPos > 560 || yPos < 60 || yPos > 660 || snake.ownCollision() {
				snake.game.Crashed()
				snake.game.gameOver()
			}
		}
	}
	return nil
}

func (snake *Snake) Draw(screen *ebiten.Image, dotTime int) error {
	if snake.game.alive {
		snake.UpdatePos(dotTime)
	}

	drawer := &ebiten.DrawImageOptions{}
	x, y := snake.getHeadPos()
	drawer.GeoM.Translate(x, y)

	if snake.dir == 4 {
		screen.DrawImage(&snake.headImgUp, drawer)
	} else if snake.dir == 2 {
		screen.DrawImage(&snake.headImgDown, drawer)
	} else if snake.dir == 3 {
		screen.DrawImage(&snake.headImgRight, drawer)
	} else if snake.dir == 1 {
		screen.DrawImage(&snake.headImgLeft, drawer)
	}

	for i := 0; i < snake.length; i++ {
		partDrawer := &ebiten.DrawImageOptions{}
		x, y := snake.getPart(i)
		partDrawer.GeoM.Translate(x, y)
		screen.DrawImage(&snake.bodyImg, partDrawer)
	}

	return nil
}

// 20 Value is due to tiles being 20x20
func (snake *Snake) UpdatePos(dotTime int) {
	if dotTime == 1 {
		if snake.newLength > 0 {
			snake.length++
			snake.newLength--
		}
		switch snake.dir {
		case 4:
			snake.turnDir(0, -20)
		case 2:
			snake.turnDir(0, +20)
		case 3:
			snake.turnDir(20, 0)
		case 1:
			snake.turnDir(-20, 0)
		}
	}
}

func (snake *Snake) addPoint() {
	snake.score++
	snake.newLength++
}

func (snake *Snake) getHeadPos() (float64, float64) {
	return snake.parts[0].X, snake.parts[0].Y
}

func (snake *Snake) getPart(pos int) (float64, float64) {
	return snake.parts[pos+1].X, snake.parts[pos+1].Y
}

func (snake *Snake) turnDir(newXPos, newYPos float64) {
	newX := snake.parts[0].X + newXPos
	newY := snake.parts[0].Y + newYPos
	snake.updateParts(newX, newY)
}

func (snake *Snake) updateParts(newX, newY float64) {
	snake.parts = append([]Position{{newX, newY}}, snake.parts...)
	snake.parts = snake.parts[:snake.length+1]
}

func (snake *Snake) ownCollision() bool {
	x, y := snake.getHeadPos()
	for i := 1; i < len(snake.parts); i++ {
		if x == snake.parts[i].X && y == snake.parts[i].Y {
			return true
		}
	}
	return false
}
