package game

// updateAI controls the AI paddle movement in single player mode
func (g *Game) updateAI() {
	paddleCenterY := g.paddle2.Y + g.paddle2.Height/2
	ballCenterY := g.ball.Y + g.ball.Size/2
	
	// Only track the ball when it's moving toward the AI paddle
	if g.ball.VelX > 0 {
		deadzone := 10.0 // Prevent jittery movement
		
		if ballCenterY < paddleCenterY-deadzone {
			g.paddle2.MoveUp()
		} else if ballCenterY > paddleCenterY+deadzone {
			g.paddle2.MoveDown()
		}
	}
}
