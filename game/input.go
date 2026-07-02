package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// handleInput processes keyboard input based on current game state
func (g *Game) handleInput() {
	// Fullscreen toggle available in all states
	g.handleFullscreenToggle()

	switch g.gameState {
	case GameStateMenu:
		g.handleMenuInput()
	case GameStateOptions:
		g.handleOptionsInput()
	case GameStateWaitingToStart:
		g.handleWaitingInput()
	case GameStatePlaying:
		g.handlePlayingInput()
	case GameStatePaused:
		g.handlePausedInput()
	case GameStateGameOver:
		g.handleGameOverInput()
	}
}

// handleFullscreenToggle processes fullscreen toggle (F or F11 key)
func (g *Game) handleFullscreenToggle() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.isFullscreen = !g.isFullscreen
		ebiten.SetFullscreen(g.isFullscreen)
	}
}

// handleMenuInput processes input in menu state
func (g *Game) handleMenuInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.menuSelection--
		if g.menuSelection < 0 {
			g.menuSelection = 2
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.menuSelection++
		if g.menuSelection > 2 {
			g.menuSelection = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.menuSelection {
		case 0:
			g.gameMode = GameModeTwoPlayer
			g.gameState = GameStateWaitingToStart
		case 1:
			g.gameMode = GameModeSinglePlayer
			g.gameState = GameStateWaitingToStart
		case 2:
			g.gameState = GameStateOptions
			g.optionsSelection = 0
		}
	}
}

// handleOptionsInput processes input in options menu state
func (g *Game) handleOptionsInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.optionsSelection--
		if g.optionsSelection < 0 {
			g.optionsSelection = 2
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.optionsSelection++
		if g.optionsSelection > 2 {
			g.optionsSelection = 0
		}
	}
	
	// Toggle selected option with Enter or Arrow Left/Right
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || 
	   inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || 
	   inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
	   inpututil.IsKeyJustPressed(ebiten.KeyA) ||
	   inpututil.IsKeyJustPressed(ebiten.KeyD) {
		switch g.optionsSelection {
		case 0: // Audio toggle
			g.audioEnabled = !g.audioEnabled
			g.audioManager.enabled.Store(g.audioEnabled)
		case 1: // AI Difficulty
			if g.aiDifficulty == AIDifficultyEasy {
				g.aiDifficulty = AIDifficultyMedium
			} else if g.aiDifficulty == AIDifficultyMedium {
				g.aiDifficulty = AIDifficultyHard
			} else {
				g.aiDifficulty = AIDifficultyEasy
			}
		case 2: // Back to menu
			g.gameState = GameStateMenu
		}
	}
	
	// ESC to go back
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.gameState = GameStateMenu
	}
}

// handleWaitingInput processes input in waiting-to-start state
func (g *Game) handleWaitingInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.ball.Reset()
		g.gameState = GameStatePlaying
		
		// Initialize game start time if first round
		if g.gameStartTime == 0 {
			g.gameStartTime = int64(g.totalFrames)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.resetGame()
	}
}

// handlePlayingInput processes input during active gameplay
func (g *Game) handlePlayingInput() {
	// P or ESC to pause
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.gameState = GameStatePaused
		return
	}
	
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
}

// handlePausedInput processes input in paused state
func (g *Game) handlePausedInput() {
	// P or ESC to resume
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.gameState = GameStatePlaying
	}
	// Enter to return to menu
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.resetGame()
	}
}

// handleGameOverInput processes input in game over state
func (g *Game) handleGameOverInput() {
	// R to replay
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.score1 = 0
		g.score2 = 0
		g.winner = 0
		g.currentRally = 0
		g.longestRally = 0
		g.totalFrames = 0
		g.gameStartTime = 0
		g.ball.Reset()
		g.paddle1.VelocityY = 0
		g.paddle2.VelocityY = 0
		g.gameState = GameStateWaitingToStart
	}
	
	// M or ESC to return to menu
	if inpututil.IsKeyJustPressed(ebiten.KeyM) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.resetGame()
	}
}
