package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func NewPaddle(x, y float64) *Paddle {
	return &Paddle{
		X:            x,
		Y:            y,
		Width:        20,
		Height:       100,
		Speed:        6,
		VelocityY:    0,
		Acceleration: 1.2,
		Friction:     0.82,
		MaxSpeed:     12,
	}
}

func (p *Paddle) Update() {
	p.Y += p.VelocityY
	p.VelocityY *= p.Friction

	if p.VelocityY > -0.1 && p.VelocityY < 0.1 {
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

func (p *Paddle) MoveUp() {
	p.VelocityY -= p.Acceleration
	if p.VelocityY < -p.MaxSpeed {
		p.VelocityY = -p.MaxSpeed
	}
}

func (p *Paddle) MoveDown() {
	p.VelocityY += p.Acceleration
	if p.VelocityY > p.MaxSpeed {
		p.VelocityY = p.MaxSpeed
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	white := color.RGBA{255, 255, 255, 255}
	maxX := int(p.X + p.Width)
	maxY := int(p.Y + p.Height)
	
	for x := int(p.X); x < maxX; x++ {
		for y := int(p.Y); y < maxY; y++ {
			if x >= 0 && x < ScreenWidth && y >= 0 && y < ScreenHeight {
				screen.Set(x, y, white)
			}
		}
	}
}