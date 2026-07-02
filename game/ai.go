package game

import "math/rand"

// updateAI controls the AI paddle movement in single player mode with difficulty levels
func (g *Game) updateAI() {
	paddleCenterY := g.paddle2.Y + g.paddle2.Height/2
	ballCenterY := g.ball.Y + g.ball.Size/2
	
	// Only track the ball when it's moving toward the AI paddle
	if g.ball.VelX > 0 {
		var deadzone float64
		var reactionSpeed float64
		var predictEnabled bool
		
		switch g.aiDifficulty {
		case AIDifficultyEasy:
			deadzone = g.config.AI.EasyDeadzone
			reactionSpeed = g.config.AI.EasyReactionSpeed
			if rand.Float64() > reactionSpeed {
				return
			}
		case AIDifficultyMedium:
			deadzone = g.config.AI.MediumDeadzone
			reactionSpeed = g.config.AI.MediumReactionSpeed
		case AIDifficultyHard:
			deadzone = g.config.AI.HardDeadzone
			reactionSpeed = g.config.AI.HardReactionSpeed
			predictEnabled = g.config.AI.HardPredictionEnabled
			
			// Predict where ball will be (simple linear prediction)
			if predictEnabled && g.ball.VelX > 0 {
				distanceX := (g.paddle2.X - g.ball.X)
				timeToReach := distanceX / g.ball.VelX
				predictedY := g.ball.Y + (g.ball.VelY * timeToReach)
				ballCenterY = predictedY + g.ball.Size/2
			}
		}
		
		if ballCenterY < paddleCenterY-deadzone {
			g.paddle2.MoveUp()
		} else if ballCenterY > paddleCenterY+deadzone {
			g.paddle2.MoveDown()
		}
	}
}
