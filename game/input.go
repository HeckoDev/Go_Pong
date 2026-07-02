package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// handleInput processes keyboard input based on current game state
func (g *Game) handleInput() {
	switch g.gameState {
	case GameStateMenu:
		g.handleMenuInput()
	case GameStateWaitingToStart:
		g.handleWaitingInput()
	case GameStatePlaying:
		g.handlePlayingInput()
	case GameStateGameOver:
		g.handleGameOverInput()
	}
}

// handleMenuInput processes input in menu state
func (g *Game) handleMenuInput() {
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
}

// handleWaitingInput processes input in waiting-to-start state
func (g *Game) handleWaitingInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.ball.Reset()
		g.gameState = GameStatePlaying
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.resetGame()
	}
}

// handlePlayingInput processes input during active gameplay
func (g *Game) handlePlayingInput() {
	// Player 1 controls (W/S)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.paddle1.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.paddle1.MoveDown()
	}

	// Player 2 controls (Arrow keys) - only in two player mode
	if g.gameMode == GameModeTwoPlayer {
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			g.paddle2.MoveUp()
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			g.paddle2.MoveDown()
		}
	}

	// ESC to return to menu
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.resetGame()
	}
}

// handleGameOverInput processes input in game over state
func (g *Game) handleGameOverInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.resetGame()
	}
}
