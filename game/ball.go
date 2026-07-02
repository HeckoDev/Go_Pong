package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ball represents the game ball with position, velocity, and visual trail
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

// TrailPoint represents a single point in the ball's visual trail
type TrailPoint struct {
	X     float64
	Y     float64
	Alpha float64
	Age   int
}

// NewBall creates a new ball at the specified position
func NewBall(x, y float64) *Ball {
	ball := &Ball{
		X:          x,
		Y:          y,
		Size:       BallSize,
		BaseSpeed:  BallBaseSpeed,
		SpeedBoost: BallSpeedBoost,
	}
	ball.Reset()
	return ball
}

// Reset resets the ball to center position with random direction
func (b *Ball) Reset() {
	b.X = ScreenWidth / 2
	b.Y = ScreenHeight / 2
	b.BounceCount = 0
	b.Trail = nil
	
	if rand.Intn(2) == 0 {
		b.VelX = -b.BaseSpeed
	} else {
		b.VelX = b.BaseSpeed
	}
	
	b.VelY = float64(rand.Intn(BallVelYRangeMax-BallVelYRangeMin+1) + BallVelYRangeMin)
}

// IncreaseBounce increments bounce count and increases speed every N bounces
func (b *Ball) IncreaseBounce() {
	b.BounceCount++
	
	if b.BounceCount%BallSpeedIncreaseEvery == 0 {
		if b.VelX > 0 {
			b.VelX += b.SpeedBoost
		} else {
			b.VelX -= b.SpeedBoost
		}
	}
}

// Update updates ball position, trail, and checks wall collisions
// Returns true if ball hit a wall
func (b *Ball) Update() bool {
	speed := b.HorizontalSpeed()
	trailLength := int(speed * BallTrailSpeedFactor)
	if trailLength < BallTrailMinLength {
		trailLength = BallTrailMinLength
	}
	if trailLength > BallTrailMaxLength {
		trailLength = BallTrailMaxLength
	}
	
	// Add new trail point
	b.Trail = append(b.Trail, TrailPoint{
		X:     b.X + b.Size/2,
		Y:     b.Y + b.Size/2,
		Alpha: 1.0,
		Age:   0,
	})
	
	// Update and filter trail points in-place (avoid reallocation)
	writeIdx := 0
	for i := range b.Trail {
		b.Trail[i].Age++
		b.Trail[i].Alpha = 1.0 - float64(b.Trail[i].Age)/float64(trailLength)
		if b.Trail[i].Age < trailLength {
			if writeIdx != i {
				b.Trail[writeIdx] = b.Trail[i]
			}
			writeIdx++
		}
	}
	b.Trail = b.Trail[:writeIdx]
	
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

// HorizontalSpeed returns the absolute horizontal speed of the ball
func (b *Ball) HorizontalSpeed() float64 {
	if b.VelX < 0 {
		return -b.VelX
	}
	return b.VelX
}

// Draw renders the ball and its trail to the screen
func (b *Ball) Draw(screen *ebiten.Image) {
	for _, point := range b.Trail {
		trailColor := ColorTrail
		trailColor.A = uint8(point.Alpha * BallTrailAlphaMax)
		drawCircle(screen, point.X, point.Y, b.Size*BallTrailSizeFactor, trailColor)
	}
	
	centerX := b.X + b.Size/2
	centerY := b.Y + b.Size/2
	radius := b.Size / 2
	
	drawCircle(screen, centerX, centerY, radius, ColorForeground)
}