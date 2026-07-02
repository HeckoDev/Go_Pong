package game

// checkPaddleCollisions checks and handles collisions between ball and both paddles
// Returns true if a collision occurred
func (g *Game) checkPaddleCollisions() bool {
	// Check collision with left paddle (player 1)
	if collideBallWithPaddle(g.ball, g.paddle1, true) {
		return true
	}
	
	// Check collision with right paddle (player 2 or AI)
	if collideBallWithPaddle(g.ball, g.paddle2, false) {
		return true
	}
	
	return false
}

// checkScoring checks if the ball went out of bounds and updates scores
func (g *Game) checkScoring() {
	if g.ball.X < 0 {
		// Ball went past left side - player 2 scores
		g.explosion = NewExplosion(g.ball.X+g.ball.Size/2, g.ball.Y+g.ball.Size/2)
		g.score2++
		g.audioManager.PlayScore()
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	} else if g.ball.X > ScreenWidth {
		// Ball went past right side - player 1 scores
		g.explosion = NewExplosion(g.ball.X+g.ball.Size/2, g.ball.Y+g.ball.Size/2)
		g.score1++
		g.audioManager.PlayScore()
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	}
}

// checkWinCondition checks if a player has reached the winning score
func (g *Game) checkWinCondition() {
	if g.score1 >= WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 1
	} else if g.score2 >= WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 2
	}
}

// resetGame resets all game state to initial values
func (g *Game) resetGame() {
	g.score1 = 0
	g.score2 = 0
	g.gameState = GameStateMenu
	g.winner = 0
	g.menuSelection = 0
	g.ball.Reset()
	g.paddle1.VelocityY = 0
	g.paddle2.VelocityY = 0
}
