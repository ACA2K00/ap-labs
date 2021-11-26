package scripts

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game : Main object of the scene. Parent of everything
type Game struct {
	snake        *Snake
	hud          *Hud
	foods        []*Food
	enemies      []*EnemySnake
	enemiesChan  []chan int
	snakeChannel chan int
	alive        bool
	crashed      bool
	numFood      int
	numEnemies   int
	score        int
	dotTime      int
	crashCount   int
}

// NewGame : Starts a new game assigning variables
func NewGame(nFood int, nEnemies int) Game {
	game := Game{
		alive:      true,
		crashed:    false,
		score:      0,
		dotTime:    0,
		crashCount: 10,
		numFood:    nFood,
		numEnemies: nEnemies,
	}

	foodArray := make([]*Food, game.numFood)
	for i := 0; i < game.numFood; i++ {
		foodArray[i] = GenFood()
		time.Sleep(20)
	}
	game.foods = foodArray
	game.snake = createSnake(&game)
	game.snakeChannel = make(chan int)
	go func() {
		err := game.snake.Behavior()
		if err != nil {
			fmt.Println("Error on main snake routine")
			os.Exit(1)
		}
	}()

	arrayEnemies := make([]*EnemySnake, game.numEnemies)
	for i := 0; i < len(arrayEnemies); i++ {
		arrayEnemies[i] = CreateEnemySnake(&game)
		time.Sleep(20)
	}
	enemiesChan := make([]chan int, game.numEnemies)
	for i := 0; i < len(enemiesChan); i++ {
		enemiesChan[i] = make(chan int)
		arrayEnemies[i].channelMovements = enemiesChan[i]
		go func() {
			err := arrayEnemies[i].ChannelPipe()
			if err != nil {
				fmt.Println("Error on enemy snake routine")
				os.Exit(1)
			}
		}()
		time.Sleep(20)
	}
	game.enemiesChan = enemiesChan
	game.enemies = arrayEnemies

	game.hud = initHud(&game)
	fmt.Printf("Food: %d \n", nFood)
	fmt.Printf("Enemies: %d \n", nEnemies)
	return game
}

func (game *Game) gameOver() {
	game.alive = false
	game.hud.game.crashCount = 0
}

func (game *Game) Crashed() {
	game.crashed = true
}

// Update the main process of the game
func (game *Game) Update() error {

	if game.alive {
		if game.numFood == 0 {
			game.hud.game.alive = false
			bigEnemy := game.enemies[0]
			for i := 1; i < len(game.enemies); i++ {
				if game.enemies[i].score > bigEnemy.score {
					bigEnemy = game.enemies[i]
				}
			}
			if game.snake.score > bigEnemy.score {
				game.hud.bigger = true
			} else {
				game.hud.bigger = false
			}

		}

		//update the channels
		game.dotTime = (game.dotTime + 1) % 5

		if err := game.snake.Update(game.dotTime); err != nil {
			game.snakeChannel <- game.dotTime
		}
		for i := 0; i < len(game.enemiesChan); i++ {
			game.enemiesChan[i] <- game.dotTime
		}
		xPos, yPos := game.snake.getHeadPos()
		for i := 0; i < len(game.foods); i++ {
			if xPos == game.foods[i].x && yPos == game.foods[i].y {
				game.foods[i].y = -20
				game.foods[i].x = -20
				game.hud.addPoint()
				game.numFood--
				game.snake.addPoint()
				break
			}
		}
		for j := 0; j < len(game.enemies); j++ {
			xPos, yPos := game.enemies[j].GetHeadPos()
			for i := 0; i < len(game.foods); i++ {
				if xPos == game.foods[i].x && yPos == game.foods[i].y {
					game.foods[i].y = -20
					game.foods[i].x = -20
					game.numFood--
					game.enemies[j].AddPoint()
					break
				}
			}
		}
	}
	for i := 0; i < game.numFood; i++ {
		if err := game.foods[i].Update(game.dotTime); err != nil {
			return err
		}
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) error {

	drawer := &ebiten.DrawImageOptions{}
	drawer.GeoM.Translate(0, 0)
	background, _, _ := ebitenutil.NewImageFromFile("img_res/bg.png", ebiten.FilterLinear)
	screen.DrawImage(background, drawer)

	if err := game.snake.Draw(screen, game.dotTime); err != nil {
		return err
	}

	for _, enemy := range game.enemies {
		if err := enemy.Draw(screen, game.dotTime); err != nil {
			return err
		}
	}

	if err := game.hud.Draw(screen); err != nil {
		return err
	}

	for i := 0; i < len(game.foods); i++ {
		if err := game.foods[i].Draw(screen, game.dotTime); err != nil {
			return err
		}
	}
	return nil
}
