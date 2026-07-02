package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// drawCircle draws a filled circle at (cx, cy) with given radius and color
// This is the performance-critical rendering function called many times per frame
// TODO: Consider pre-rendering to ebiten.Image for better performance
func drawCircle(screen *ebiten.Image, cx, cy, radius float64, col color.Color) {
	r2 := radius * radius
	minX := int(cx - radius)
	maxX := int(cx + radius)
	minY := int(cy - radius)
	maxY := int(cy + radius)
	
	// Clamp to screen bounds to avoid unnecessary iterations
	if minX < 0 {
		minX = 0
	}
	if maxX >= ScreenWidth {
		maxX = ScreenWidth - 1
	}
	if minY < 0 {
		minY = 0
	}
	if maxY >= ScreenHeight {
		maxY = ScreenHeight - 1
	}
	
	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			dx := float64(px) - cx
			dy := float64(py) - cy
			if dx*dx+dy*dy <= r2 {
				screen.Set(px, py, col)
			}
		}
	}
}

// collideBallWithPaddle checks and handles collision between ball and paddle
// Returns true if collision occurred
func collideBallWithPaddle(ball *Ball, paddle *Paddle, isLeftPaddle bool) bool {
	// Check AABB collision
	if ball.X+ball.Size < paddle.X || ball.X > paddle.X+paddle.Width ||
		ball.Y+ball.Size < paddle.Y || ball.Y > paddle.Y+paddle.Height {
		return false
	}
	
	// Collision detected - bounce ball
	ball.VelX = -ball.VelX
	
	// Position ball outside paddle to prevent sticking
	if isLeftPaddle {
		ball.X = paddle.X + paddle.Width
	} else {
		ball.X = paddle.X - ball.Size
	}
	
	// Increase bounce count and speed
	ball.IncreaseBounce()
	
	// Apply spin effect based on where ball hit the paddle
	hitPos := (ball.Y + ball.Size/2) - (paddle.Y + paddle.Height/2)
	ball.VelY += hitPos * PaddleSpinFactor
	
	// Clamp vertical velocity
	if ball.VelY > MaxVelY {
		ball.VelY = MaxVelY
	} else if ball.VelY < -MaxVelY {
		ball.VelY = -MaxVelY
	}
	
	return true
}

