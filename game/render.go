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
	case GameStateWaitingToStart:
		g.drawWaitingMessages(screen)
		g.drawControls(screen)
	case GameStatePlaying:
		g.drawControls(screen)
	case GameStateGameOver:
		g.drawGameOver(screen)
	}
}

// drawMenu renders the main menu
func (g *Game) drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "PONG GAME", ScreenWidth/2-40, ScreenHeight/2+MenuTitleOffsetY)
	
	option1 := "  Two Player Mode"
	option2 := "  Single Player Mode"
	
	if g.menuSelection == 0 {
		option1 = "> Two Player Mode"
	} else {
		option2 = "> Single Player Mode"
	}
	
	ebitenutil.DebugPrintAt(screen, option1, ScreenWidth/2-70, ScreenHeight/2+MenuOption1OffsetY)
	ebitenutil.DebugPrintAt(screen, option2, ScreenWidth/2-70, ScreenHeight/2+MenuOption2OffsetY)
	ebitenutil.DebugPrintAt(screen, "Use W/S or ↑/↓ to navigate", ScreenWidth/2-90, ScreenHeight/2+MenuNavigateOffsetY)
	ebitenutil.DebugPrintAt(screen, "Press ENTER to select", ScreenWidth/2-70, ScreenHeight/2+MenuSelectOffsetY)
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
	ebitenutil.DebugPrintAt(screen, winnerText, ScreenWidth/2-60, ScreenHeight/2+MessageWinnerOffsetY)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Final Score: %d - %d", g.score1, g.score2), ScreenWidth/2-70, ScreenHeight/2+MessageScoreOffsetY)
	ebitenutil.DebugPrintAt(screen, "Press ENTER to return to menu", ScreenWidth/2-90, ScreenHeight/2+MessageMenuOffsetY)
}

// drawControls renders control hints at the bottom of the screen
func (g *Game) drawControls(screen *ebiten.Image) {
	if g.gameMode == GameModeTwoPlayer {
		ebitenutil.DebugPrint(screen, "Player 1: W/S | Player 2: ↑/↓ | Enter: Reset Ball | ESC: Menu")
	} else {
		ebitenutil.DebugPrint(screen, "Player: W/S | Enter: Reset Ball | ESC: Menu")
	}
}

// Layout returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
