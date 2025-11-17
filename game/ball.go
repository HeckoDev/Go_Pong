package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	X           float64
	Y           float64
	VelX        float64
	VelY        float64
	Size        float64
	BounceCount int
	BaseSpeed   float64
	SpeedBoost  float64
	Trail       []TrailPoint
}

type TrailPoint struct {
	X     float64
	Y     float64
	Alpha float64
	Age   int
}

func NewBall(x, y float64) *Ball {
	ball := &Ball{
		X:          x,
		Y:          y,
		Size:       10,
		BaseSpeed:  4.0,
		SpeedBoost: 0.5,
	}
	ball.Reset()
	return ball
}

func (b *Ball) Reset() {
	b.X = ScreenWidth / 2
	b.Y = ScreenHeight / 2
	b.BounceCount = 0
	
	if rand.Intn(2) == 0 {
		b.VelX = -b.BaseSpeed
	} else {
		b.VelX = b.BaseSpeed
	}
	
	b.VelY = float64(rand.Intn(6) - 3)
}

func (b *Ball) IncreaseBounce() {
	b.BounceCount++
	
	if b.BounceCount%5 == 0 {
		if b.VelX > 0 {
			b.VelX += b.SpeedBoost
		} else {
			b.VelX -= b.SpeedBoost
		}
	}
}

func (b *Ball) Update() bool {
	speed := b.GetSpeed()
	trailLength := int(speed * 2)
	if trailLength < 3 {
		trailLength = 3
	}
	if trailLength > 15 {
		trailLength = 15
	}
	
	b.Trail = append(b.Trail, TrailPoint{
		X:     b.X + b.Size/2,
		Y:     b.Y + b.Size/2,
		Alpha: 1.0,
		Age:   0,
	})
	
	newTrail := make([]TrailPoint, 0, len(b.Trail))
	for i := range b.Trail {
		b.Trail[i].Age++
		b.Trail[i].Alpha = 1.0 - float64(b.Trail[i].Age)/float64(trailLength)
		if b.Trail[i].Age < trailLength {
			newTrail = append(newTrail, b.Trail[i])
		}
	}
	b.Trail = newTrail
	
	b.X += b.VelX
	b.Y += b.VelY

	hitWall := false

	if b.Y <= 0 || b.Y+b.Size >= ScreenHeight {
		b.VelY = -b.VelY
		hitWall = true
		
		if b.Y <= 0 {
			b.Y = 0
		}
		if b.Y+b.Size >= ScreenHeight {
			b.Y = ScreenHeight - b.Size
		}
	}

	return hitWall
}

func (b *Ball) GetSpeed() float64 {
	if b.VelX < 0 {
		return -b.VelX
	}
	return b.VelX
}

func (b *Ball) Draw(screen *ebiten.Image) {
	for _, point := range b.Trail {
		b.drawCircle(screen, point.X, point.Y, b.Size/3, color.RGBA{
			255, 0, 0, uint8(point.Alpha * 200),
		})
	}
	
	centerX := b.X + b.Size/2
	centerY := b.Y + b.Size/2
	radius := b.Size / 2
	
	b.drawCircle(screen, centerX, centerY, radius, color.RGBA{255, 255, 255, 255})
}

func (b *Ball) drawCircle(screen *ebiten.Image, cx, cy, radius float64, col color.RGBA) {
	r2 := radius * radius
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= r2 {
				px := int(cx + x)
				py := int(cy + y)
				if px >= 0 && px < ScreenWidth && py >= 0 && py < ScreenHeight {
					screen.Set(px, py, col)
				}
			}
		}
	}
}