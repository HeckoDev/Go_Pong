package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	WinningScore = 5
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStateWaitingToStart
	GameStateGameOver
)

type GameMode int

const (
	GameModeTwoPlayer GameMode = iota
	GameModeSinglePlayer
)

type Game struct {
	ball          *Ball
	paddle1       *Paddle
	paddle2       *Paddle
	score1        int
	score2        int
	gameState     GameState
	gameMode      GameMode
	winner        int
	menuSelection int
	audioManager  *AudioManager
	explosion     *Explosion
}

func NewGame() *Game {
	return &Game{
		ball:          NewBall(ScreenWidth/2, ScreenHeight/2),
		paddle1:       NewPaddle(30, ScreenHeight/2-50),
		paddle2:       NewPaddle(ScreenWidth-50, ScreenHeight/2-50),
		gameState:     GameStateMenu,
		gameMode:      GameModeTwoPlayer,
		audioManager:  NewAudioManager(),
	}
}

func (g *Game) Update() error {
	g.handleInput()

	if g.explosion != nil && g.explosion.Active {
		g.explosion.Update()
	}

	if g.gameState == GameStatePlaying {
		if g.ball.Update() {
			g.audioManager.PlayWallHit()
		}

		g.paddle1.Update()
		g.paddle2.Update()

		if g.gameMode == GameModeSinglePlayer {
			g.updateAI()
		}

		if g.checkPaddleCollisions() {
			g.audioManager.PlayPaddleHit()
		}

		g.checkScoring()
	}

	return nil
}

func (g *Game) handleInput() {
	switch g.gameState {
	case GameStateMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
			g.menuSelection = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
			g.menuSelection = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.menuSelection == 0 {
				g.gameMode = GameModeTwoPlayer
			} else {
				g.gameMode = GameModeSinglePlayer
			}
			g.gameState = GameStateWaitingToStart
		}

	case GameStateWaitingToStart:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.gameState = GameStatePlaying
			g.ball.Reset()
		}

	case GameStateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.resetGame()
		}

	case GameStatePlaying:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.ball.Reset()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.gameState = GameStateMenu
		}
	}

	if g.gameState == GameStatePlaying {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.paddle1.MoveUp()
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.paddle1.MoveDown()
		}

		if g.gameMode == GameModeTwoPlayer {
			if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
				g.paddle2.MoveUp()
			}
			if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
				g.paddle2.MoveDown()
			}
		}
	}
}

func (g *Game) checkPaddleCollisions() bool {
	hitPaddle := false

	if g.ball.X <= g.paddle1.X+g.paddle1.Width &&
		g.ball.X+g.ball.Size >= g.paddle1.X &&
		g.ball.Y <= g.paddle1.Y+g.paddle1.Height &&
		g.ball.Y+g.ball.Size >= g.paddle1.Y {
		g.ball.VelX = -g.ball.VelX
		g.ball.X = g.paddle1.X + g.paddle1.Width
		g.ball.IncreaseBounce()
		hitPaddle = true
		
		hitPos := (g.ball.Y + g.ball.Size/2) - (g.paddle1.Y + g.paddle1.Height/2)
		g.ball.VelY += hitPos * 0.1
	}

	if g.ball.X+g.ball.Size >= g.paddle2.X &&
		g.ball.X <= g.paddle2.X+g.paddle2.Width &&
		g.ball.Y <= g.paddle2.Y+g.paddle2.Height &&
		g.ball.Y+g.ball.Size >= g.paddle2.Y {
		g.ball.VelX = -g.ball.VelX
		g.ball.X = g.paddle2.X - g.ball.Size
		g.ball.IncreaseBounce()
		hitPaddle = true
		
		hitPos := (g.ball.Y + g.ball.Size/2) - (g.paddle2.Y + g.paddle2.Height/2)
		g.ball.VelY += hitPos * 0.1
	}

	return hitPaddle
}

func (g *Game) checkScoring() {
	if g.ball.X < 0 {
		g.explosion = NewExplosion(g.ball.X+g.ball.Size/2, g.ball.Y+g.ball.Size/2)
		g.score2++
		g.audioManager.PlayScore()
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	} else if g.ball.X > ScreenWidth {
		g.explosion = NewExplosion(g.ball.X+g.ball.Size/2, g.ball.Y+g.ball.Size/2)
		g.score1++
		g.audioManager.PlayScore()
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	}
}

func (g *Game) checkWinCondition() {
	if g.score1 >= WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 1
	} else if g.score2 >= WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 2
	}
}

func (g *Game) resetGame() {
	g.score1 = 0
	g.score2 = 0
	g.gameState = GameStateMenu
	g.winner = 0
	g.menuSelection = 0
	g.ball.X = ScreenWidth / 2
	g.ball.Y = ScreenHeight / 2
	g.ball.VelX = 0
	g.ball.VelY = 0
}

func (g *Game) updateAI() {
	paddleCenterY := g.paddle2.Y + g.paddle2.Height/2
	ballCenterY := g.ball.Y + g.ball.Size/2
	aiSpeed := g.paddle2.Speed * 0.8
	
	if g.ball.VelX > 0 {
		if ballCenterY < paddleCenterY-10 {
			g.paddle2.Y -= aiSpeed
			if g.paddle2.Y < 0 {
				g.paddle2.Y = 0
			}
		} else if ballCenterY > paddleCenterY+10 {
			g.paddle2.Y += aiSpeed
			if g.paddle2.Y+g.paddle2.Height > ScreenHeight {
				g.paddle2.Y = ScreenHeight - g.paddle2.Height
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.drawCenterLine(screen)
	g.paddle1.Draw(screen)
	g.paddle2.Draw(screen)
	g.ball.Draw(screen)

	if g.explosion != nil && g.explosion.Active {
		g.explosion.Draw(screen)
	}

	g.drawScore(screen)
	g.drawGameStateMessages(screen)
}

func (g *Game) drawCenterLine(screen *ebiten.Image) {
	centerX := ScreenWidth / 2
	dashHeight := 10
	dashGap := 10

	for y := 0; y < ScreenHeight; y += dashHeight + dashGap {
		for i := 0; i < dashHeight && y+i < ScreenHeight; i++ {
			for j := 0; j < 2; j++ {
				screen.Set(centerX+j, y+i, color.RGBA{255, 255, 255, 255})
			}
		}
	}
}

func (g *Game) drawScore(screen *ebiten.Image) {
	score1Text := fmt.Sprintf("%d", g.score1)
	ebitenutil.DebugPrintAt(screen, score1Text, ScreenWidth/2-100, 50)
	
	score2Text := fmt.Sprintf("%d", g.score2)
	ebitenutil.DebugPrintAt(screen, score2Text, ScreenWidth/2+80, 50)
}

func (g *Game) drawGameStateMessages(screen *ebiten.Image) {
	switch g.gameState {
	case GameStateMenu:
		g.drawMenu(screen)
	case GameStateWaitingToStart:
		if g.score1 == 0 && g.score2 == 0 {
			modeText := "TWO PLAYER MODE"
			if g.gameMode == GameModeSinglePlayer {
				modeText = "SINGLE PLAYER MODE"
			}
			ebitenutil.DebugPrintAt(screen, modeText, ScreenWidth/2-70, ScreenHeight/2-60)
			ebitenutil.DebugPrintAt(screen, "Press ENTER to start", ScreenWidth/2-70, ScreenHeight/2-20)
		} else {
			ebitenutil.DebugPrintAt(screen, "Press ENTER to continue", ScreenWidth/2-80, ScreenHeight/2-10)
		}
		g.drawControls(screen)
	case GameStatePlaying:
		g.drawControls(screen)
	case GameStateGameOver:
		winnerText := fmt.Sprintf("PLAYER %d WINS!", g.winner)
		if g.gameMode == GameModeSinglePlayer {
			if g.winner == 1 {
				winnerText = "YOU WIN!"
			} else {
				winnerText = "AI WINS!"
			}
		}
		ebitenutil.DebugPrintAt(screen, winnerText, ScreenWidth/2-60, ScreenHeight/2-40)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Final Score: %d - %d", g.score1, g.score2), ScreenWidth/2-70, ScreenHeight/2-10)
		ebitenutil.DebugPrintAt(screen, "Press ENTER to return to menu", ScreenWidth/2-90, ScreenHeight/2+20)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "PONG GAME", ScreenWidth/2-40, ScreenHeight/2-100)
	
	option1 := "  Two Player Mode"
	option2 := "  Single Player Mode"
	
	if g.menuSelection == 0 {
		option1 = "> Two Player Mode"
	} else {
		option2 = "> Single Player Mode"
	}
	
	ebitenutil.DebugPrintAt(screen, option1, ScreenWidth/2-70, ScreenHeight/2-30)
	ebitenutil.DebugPrintAt(screen, option2, ScreenWidth/2-70, ScreenHeight/2)
	ebitenutil.DebugPrintAt(screen, "Use W/S or ↑/↓ to navigate", ScreenWidth/2-90, ScreenHeight/2+60)
	ebitenutil.DebugPrintAt(screen, "Press ENTER to select", ScreenWidth/2-70, ScreenHeight/2+80)
}

func (g *Game) drawControls(screen *ebiten.Image) {
	if g.gameMode == GameModeTwoPlayer {
		ebitenutil.DebugPrint(screen, "Player 1: W/S | Player 2: ↑/↓ | Enter: Reset Ball | ESC: Menu")
	} else {
		ebitenutil.DebugPrint(screen, "Player: W/S | Enter: Reset Ball | ESC: Menu")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}