package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	X       float64
	Y       float64
	VelX    float64
	VelY    float64
	Life    int
	MaxLife int
	Size    float64
	Color   color.RGBA
}

type Explosion struct {
	Particles []Particle
	Active    bool
}

func NewExplosion(x, y float64) *Explosion {
	particles := make([]Particle, 30)
	
	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		speed := 2.0 + rand.Float64()*4.0
		
		particles[i] = Particle{
			X:       x,
			Y:       y,
			VelX:    math.Cos(angle) * speed,
			VelY:    math.Sin(angle) * speed,
			Life:    30,
			MaxLife: 30,
			Size:    2 + rand.Float64()*3,
			Color:   color.RGBA{255, uint8(100 + rand.Intn(155)), 0, 255},
		}
	}
	
	return &Explosion{
		Particles: particles,
		Active:    true,
	}
}

func (e *Explosion) Update() {
	if !e.Active {
		return
	}
	
	allDead := true
	for i := range e.Particles {
		if e.Particles[i].Life > 0 {
			e.Particles[i].X += e.Particles[i].VelX
			e.Particles[i].Y += e.Particles[i].VelY
			e.Particles[i].VelY += 0.2
			e.Particles[i].VelX *= 0.98
			e.Particles[i].VelY *= 0.98
			e.Particles[i].Life--
			
			lifeRatio := float64(e.Particles[i].Life) / float64(e.Particles[i].MaxLife)
			e.Particles[i].Color.A = uint8(lifeRatio * 255)
			
			allDead = false
		}
	}
	
	if allDead {
		e.Active = false
	}
}

func (e *Explosion) Draw(screen *ebiten.Image) {
	if !e.Active {
		return
	}
	
	for _, p := range e.Particles {
		if p.Life > 0 {
			e.drawCircle(screen, p.X, p.Y, p.Size, p.Color)
		}
	}
}

func (e *Explosion) drawCircle(screen *ebiten.Image, cx, cy, radius float64, col color.RGBA) {
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
