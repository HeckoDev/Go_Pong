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
	particles := make([]Particle, ExplosionParticleCount)
	
	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		speed := ExplosionSpeedMin + rand.Float64()*(ExplosionSpeedMax-ExplosionSpeedMin)
		
		particles[i] = Particle{
			X:       x,
			Y:       y,
			VelX:    math.Cos(angle) * speed,
			VelY:    math.Sin(angle) * speed,
			Life:    ExplosionParticleLife,
			MaxLife: ExplosionParticleLife,
			Size:    ExplosionParticleSizeMin + rand.Float64()*(ExplosionParticleSizeMax-ExplosionParticleSizeMin),
			Color:   color.RGBA{255, uint8(ExplosionGreenMin + rand.Intn(ExplosionGreenMax-ExplosionGreenMin+1)), 0, 255},
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
			e.Particles[i].VelY += ExplosionGravity
			e.Particles[i].VelX *= ExplosionAirResistance
			e.Particles[i].VelY *= ExplosionAirResistance
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
			drawCircle(screen, p.X, p.Y, p.Size, p.Color)
		}
	}
}
