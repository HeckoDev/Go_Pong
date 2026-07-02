package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw renders all game elements to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(ColorBackground)
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

// drawCenterLine renders the dashed center line
func (g *Game) drawCenterLine(screen *ebiten.Image) {
	centerX := ScreenWidth / 2

	for y := 0; y < ScreenHeight; y += CenterLineDashHeight + CenterLineDashGap {
		for i := 0; i < CenterLineDashHeight && y+i < ScreenHeight; i++ {
			for j := 0; j < CenterLineWidth; j++ {
				screen.Set(centerX+j, y+i, ColorForeground)
			}
		}
	}
}

// drawScore renders the current scores
func (g *Game) drawScore(screen *ebiten.Image) {
	score1Text := fmt.Sprintf("%d", g.score1)
	ebitenutil.DebugPrintAt(screen, score1Text, ScreenWidth/2+Score1OffsetX, ScoreOffsetY)
	
	score2Text := fmt.Sprintf("%d", g.score2)
	ebitenutil.DebugPrintAt(screen, score2Text, ScreenWidth/2+Score2OffsetX, ScoreOffsetY)
}

// drawGameStateMessages renders state-specific UI messages
func (g *Game) drawGameStateMessages(screen *ebiten.Image) {
	switch g.gameState {
	case GameStateMenu:
		g.drawMenu(screen)
	case GameStateOptions:
		g.drawOptions(screen)
	case GameStateWaitingToStart:
		g.drawWaitingMessages(screen)
		g.drawControls(screen)
	case GameStatePlaying:
		g.drawControls(screen)
	case GameStatePaused:
		g.drawPauseOverlay(screen)
	case GameStateGameOver:
		g.drawGameOver(screen)
	}
}

// drawMenu renders the main menu
func (g *Game) drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "PONG GAME", ScreenWidth/2-40, ScreenHeight/2+MenuTitleOffsetY)
	
	option1 := "  Two Player Mode"
	option2 := "  Single Player Mode"
	option3 := "  Options"
	
	if g.menuSelection == 0 {
		option1 = "> Two Player Mode"
	} else if g.menuSelection == 1 {
		option2 = "> Single Player Mode"
	} else {
		option3 = "> Options"
	}
	
	ebitenutil.DebugPrintAt(screen, option1, ScreenWidth/2-70, ScreenHeight/2+MenuOption1OffsetY)
	ebitenutil.DebugPrintAt(screen, option2, ScreenWidth/2-70, ScreenHeight/2+MenuOption2OffsetY)
	ebitenutil.DebugPrintAt(screen, option3, ScreenWidth/2-70, ScreenHeight/2+MenuOption2OffsetY+30)
	ebitenutil.DebugPrintAt(screen, "Use W/S or ↑/↓ to navigate", ScreenWidth/2-90, ScreenHeight/2+MenuNavigateOffsetY+30)
	ebitenutil.DebugPrintAt(screen, "Press ENTER to select | F to toggle fullscreen", ScreenWidth/2-125, ScreenHeight/2+MenuSelectOffsetY+30)
}

// drawOptions renders the options menu
func (g *Game) drawOptions(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "OPTIONS", ScreenWidth/2-30, ScreenHeight/2+MenuTitleOffsetY)
	
	// Audio option
	audioText := "  Audio: "
	if g.audioEnabled {
		audioText += "ON"
	} else {
		audioText += "OFF"
	}
	if g.optionsSelection == 0 {
		audioText = "> " + audioText[2:]
	}
	
	// AI Difficulty option
	difficultyText := "  AI Difficulty: "
	switch g.aiDifficulty {
	case AIDifficultyEasy:
		difficultyText += "EASY"
	case AIDifficultyMedium:
		difficultyText += "MEDIUM"
	case AIDifficultyHard:
		difficultyText += "HARD"
	}
	if g.optionsSelection == 1 {
		difficultyText = "> " + difficultyText[2:]
	}
	
	// Back option
	backText := "  Back to Menu"
	if g.optionsSelection == 2 {
		backText = "> " + backText[2:]
	}
	
	ebitenutil.DebugPrintAt(screen, audioText, ScreenWidth/2-70, ScreenHeight/2+MenuOption1OffsetY)
	ebitenutil.DebugPrintAt(screen, difficultyText, ScreenWidth/2-70, ScreenHeight/2+MenuOption2OffsetY)
	ebitenutil.DebugPrintAt(screen, backText, ScreenWidth/2-70, ScreenHeight/2+MenuOption2OffsetY+30)
	ebitenutil.DebugPrintAt(screen, "Use W/S or ↑/↓ to navigate", ScreenWidth/2-90, ScreenHeight/2+MenuNavigateOffsetY+30)
	ebitenutil.DebugPrintAt(screen, "Press ENTER or ←/→ to toggle | ESC to go back", ScreenWidth/2-135, ScreenHeight/2+MenuSelectOffsetY+30)
}

// drawWaitingMessages renders messages when waiting to start
func (g *Game) drawWaitingMessages(screen *ebiten.Image) {
	if g.score1 == 0 && g.score2 == 0 {
		modeText := "TWO PLAYER MODE"
		if g.gameMode == GameModeSinglePlayer {
			modeText = "SINGLE PLAYER MODE"
		}
		ebitenutil.DebugPrintAt(screen, modeText, ScreenWidth/2-70, ScreenHeight/2+MessageModeOffsetY)
		ebitenutil.DebugPrintAt(screen, "Press ENTER to start", ScreenWidth/2-70, ScreenHeight/2+MessageStartOffsetY)
	} else {
		ebitenutil.DebugPrintAt(screen, "Press ENTER to continue", ScreenWidth/2-80, ScreenHeight/2+MessageContinueOffsetY)
	}
}

// drawGameOver renders the game over screen
func (g *Game) drawGameOver(screen *ebiten.Image) {
	winnerText := fmt.Sprintf("PLAYER %d WINS!", g.winner)
	if g.gameMode == GameModeSinglePlayer {
		if g.winner == 1 {
			winnerText = "YOU WIN!"
		} else {
			winnerText = "AI WINS!"
		}
	}
	
	// Calculate game duration in seconds
	gameDuration := g.totalFrames / 60 // 60 FPS
	minutes := gameDuration / 60
	seconds := gameDuration % 60
	
	// Victory message
	ebitenutil.DebugPrintAt(screen, winnerText, ScreenWidth/2-60, ScreenHeight/2-100)
	
	// Final score
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Final Score: %d - %d", g.score1, g.score2), ScreenWidth/2-70, ScreenHeight/2-60)
	
	// Game stats
	ebitenutil.DebugPrintAt(screen, "GAME STATS", ScreenWidth/2-45, ScreenHeight/2-20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Game Duration: %d:%02d", minutes, seconds), ScreenWidth/2-70, ScreenHeight/2+10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Longest Rally: %d hits", g.longestRally), ScreenWidth/2-70, ScreenHeight/2+30)
	
	// Options
	ebitenutil.DebugPrintAt(screen, "R - Replay   M - Menu   ESC - Exit", ScreenWidth/2-110, ScreenHeight/2+70)
}

// drawPauseOverlay renders the pause screen overlay
func (g *Game) drawPauseOverlay(screen *ebiten.Image) {
	// Semi-transparent overlay
	for y := ScreenHeight/2 - 80; y < ScreenHeight/2+80; y++ {
		for x := ScreenWidth/2 - 150; x < ScreenWidth/2+150; x++ {
			if x >= 0 && x < ScreenWidth && y >= 0 && y < ScreenHeight {
				r, g, b, a := screen.At(x, y).RGBA()
				// Darken the background
				screen.Set(x, y, ColorBackground)
				if a > 0 {
					screen.Set(x, y, ColorBackground)
				}
				_ = r
				_ = g
				_ = b
			}
		}
	}
	
	ebitenutil.DebugPrintAt(screen, "PAUSED", ScreenWidth/2-30, ScreenHeight/2-40)
	ebitenutil.DebugPrintAt(screen, "Press P or ESC to resume", ScreenWidth/2-80, ScreenHeight/2)
	ebitenutil.DebugPrintAt(screen, "Press ENTER to return to menu", ScreenWidth/2-90, ScreenHeight/2+30)
}

// drawControls renders control hints at the bottom of the screen
func (g *Game) drawControls(screen *ebiten.Image) {
	if g.gameMode == GameModeTwoPlayer {
		ebitenutil.DebugPrint(screen, "Player 1: W/S | Player 2: ↑/↓ | P/ESC: Pause")
	} else {
		ebitenutil.DebugPrint(screen, "Player: W/S | P/ESC: Pause")
	}
}

// Layout returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
