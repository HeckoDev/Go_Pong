package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Paddle represents a player's paddle with physics and movement
type Paddle struct {
	X            float64
	Y            float64
	Width        float64
	Height       float64
	Speed        float64
	VelocityY    float64
	Acceleration float64
	Friction     float64
	MaxSpeed     float64
}

// NewPaddle creates a new paddle at the specified position
func NewPaddle(x, y float64) *Paddle {
	return &Paddle{
		X:            x,
		Y:            y,
		Width:        PaddleWidth,
		Height:       PaddleHeight,
		Speed:        PaddleSpeed,
		VelocityY:    0,
		Acceleration: PaddleAcceleration,
		Friction:     PaddleFriction,
		MaxSpeed:     PaddleMaxSpeed,
	}
}

// Update updates paddle position with physics (friction, boundaries)
func (p *Paddle) Update() {
	p.Y += p.VelocityY
	p.VelocityY *= p.Friction

	if p.VelocityY > -PaddleMinVelocity && p.VelocityY < PaddleMinVelocity {
		p.VelocityY = 0
	}

	if p.Y < 0 {
		p.Y = 0
		p.VelocityY = 0
	}
	if p.Y+p.Height > ScreenHeight {
		p.Y = ScreenHeight - p.Height
		p.VelocityY = 0
	}
}

// MoveUp accelerates the paddle upward
func (p *Paddle) MoveUp() {
	p.VelocityY -= p.Acceleration
	if p.VelocityY < -p.MaxSpeed {
		p.VelocityY = -p.MaxSpeed
	}
}

// MoveDown accelerates the paddle downward
func (p *Paddle) MoveDown() {
	p.VelocityY += p.Acceleration
	if p.VelocityY > p.MaxSpeed {
		p.VelocityY = p.MaxSpeed
	}
}

// Draw renders the paddle to the screen
func (p *Paddle) Draw(screen *ebiten.Image) {
	maxX := int(p.X + p.Width)
	maxY := int(p.Y + p.Height)
	
	for x := int(p.X); x < maxX; x++ {
		for y := int(p.Y); y < maxY; y++ {
			if x >= 0 && x < ScreenWidth && y >= 0 && y < ScreenHeight {
				screen.Set(x, y, ColorForeground)
			}
		}
	}
}