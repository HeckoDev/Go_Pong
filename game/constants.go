package game

import "image/color"

// Screen dimensions
const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

// Game rules
const (
	WinningScore = 5
	MaxVelY      = 8.0
)

// Ball constants
const (
	BallSize              = 10.0
	BallBaseSpeed         = 4.0
	BallSpeedBoost        = 0.5
	BallSpeedIncreaseEvery = 5    // Bounce count interval for speed increase
	BallVelYRangeMin      = -3
	BallVelYRangeMax      = 3
	BallTrailMinLength    = 3
	BallTrailMaxLength    = 15
	BallTrailSpeedFactor  = 2.0   // Trail length = speed * factor
	BallTrailSizeFactor   = 0.333 // Trail circle size = ball size / 3
	BallTrailAlphaMax     = 200
)

// Paddle constants
const (
	PaddleWidth       = 20.0
	PaddleHeight      = 100.0
	PaddleSpeed       = 6.0
	PaddleAcceleration = 1.2
	PaddleFriction    = 0.82
	PaddleMaxSpeed    = 12.0
	PaddleMinVelocity = 0.1 // Velocity threshold to stop paddle
	PaddleSpinFactor  = 0.1 // Effect of paddle position on ball angle
)

// Explosion constants
const (
	ExplosionParticleCount  = 30
	ExplosionSpeedMin       = 2.0
	ExplosionSpeedMax       = 6.0
	ExplosionParticleLife   = 30
	ExplosionParticleSizeMin = 2.0
	ExplosionParticleSizeMax = 5.0
	ExplosionGravity        = 0.2
	ExplosionAirResistance  = 0.98
	ExplosionGreenMin       = 100
	ExplosionGreenMax       = 255
)

// UI constants
const (
	CenterLineWidth     = 2
	CenterLineDashHeight = 10
	CenterLineDashGap   = 10
	
	Score1OffsetX = -100
	Score2OffsetX = 80
	ScoreOffsetY  = 50
	
	MenuTitleOffsetY       = -100
	MenuOption1OffsetY     = -30
	MenuOption2OffsetY     = 0
	MenuNavigateOffsetY    = 60
	MenuSelectOffsetY      = 80
	
	MessageModeOffsetY     = -60
	MessageStartOffsetY    = -20
	MessageContinueOffsetY = -10
	MessageWinnerOffsetY   = -40
	MessageScoreOffsetY    = -10
	MessageMenuOffsetY     = 20
)

// AI constants
const (
	AIReactionDelay = 0 // Frames to delay AI reaction (0 = instant)
	AITargetOffset  = 0.0 // Offset for AI paddle target (for difficulty)
)

// Colors
var (
	ColorBackground = color.RGBA{0, 0, 0, 255}
	ColorForeground = color.RGBA{255, 255, 255, 255}
	ColorTrail      = color.RGBA{255, 0, 0, 255} // Red trail (alpha varies)
)
