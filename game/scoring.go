package game

// checkPaddleCollisions checks and handles collisions between ball and both paddles
// Returns true if a collision occurred
func (g *Game) checkPaddleCollisions() bool {
	// Check collision with left paddle (player 1)
	if collideBallWithPaddle(g.ball, g.paddle1, true) {
		g.currentRally++
		g.eventBus.Publish(Event{
			Type: EventBallHitPaddle,
			Data: BallHitPaddleData{
				X:         g.ball.X + g.ball.Size/2,
				Y:         g.ball.Y + g.ball.Size/2,
				IsPlayer1: true,
			},
		})
		return true
	}
	
	// Check collision with right paddle (player 2 or AI)
	if collideBallWithPaddle(g.ball, g.paddle2, false) {
		g.currentRally++
		g.eventBus.Publish(Event{
			Type: EventBallHitPaddle,
			Data: BallHitPaddleData{
				X:         g.ball.X + g.ball.Size/2,
				Y:         g.ball.Y + g.ball.Size/2,
				IsPlayer1: false,
			},
		})
		return true
	}
	
	return false
}

// checkScoring checks if the ball went out of bounds and updates scores
func (g *Game) checkScoring() {
	if g.ball.X < 0 {
		// Ball went past left side - player 2 scores
		g.score2++
		
		g.eventBus.Publish(Event{
			Type: EventScorePoint,
			Data: ScorePointData{
				Player: 2,
				X:      g.ball.X + g.ball.Size/2,
				Y:      g.ball.Y + g.ball.Size/2,
			},
		})
		
		// Update rally stats
		if g.currentRally > g.longestRally {
			g.longestRally = g.currentRally
		}
		g.currentRally = 0
		
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	} else if g.ball.X > ScreenWidth {
		// Ball went past right side - player 1 scores
		g.score1++
		
		g.eventBus.Publish(Event{
			Type: EventScorePoint,
			Data: ScorePointData{
				Player: 1,
				X:      g.ball.X + g.ball.Size/2,
				Y:      g.ball.Y + g.ball.Size/2,
			},
		})
		
		// Update rally stats
		if g.currentRally > g.longestRally {
			g.longestRally = g.currentRally
		}
		g.currentRally = 0
		
		g.checkWinCondition()
		if g.gameState != GameStateGameOver {
			g.gameState = GameStateWaitingToStart
		}
	}
}

// checkWinCondition checks if a player has reached the winning score
func (g *Game) checkWinCondition() {
	if g.score1 >= g.config.Game.WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 1
		g.recordGameStats()
	} else if g.score2 >= g.config.Game.WinningScore {
		g.gameState = GameStateGameOver
		g.winner = 2
		g.recordGameStats()
	}
}

// recordGameStats records the finished game in statistics
func (g *Game) recordGameStats() {
	duration := g.totalFrames / 60 // Convert frames to seconds
	g.stats.RecordGame(g.gameMode, g.winner, g.score1, g.score2, duration, g.longestRally)
	
	// Save stats to file
	SaveStats("saves/stats.json", g.stats)
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
	
	// Reset stats
	g.currentRally = 0
	g.longestRally = 0
	g.totalFrames = 0
	g.gameStartTime = 0
}
