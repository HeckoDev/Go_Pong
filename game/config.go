package game

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Config holds all game configuration
type Config struct {
	Game      GameConfig      `toml:"game"`
	Ball      BallConfig      `toml:"ball"`
	Paddle    PaddleConfig    `toml:"paddle"`
	Explosion ExplosionConfig `toml:"explosion"`
	AI        AIConfig        `toml:"ai"`
	Audio     AudioConfig     `toml:"audio"`
	Video     VideoConfig     `toml:"video"`
}

// GameConfig holds general game rules
type GameConfig struct {
	WinningScore int     `toml:"winning_score"`
	MaxVelocityY float64 `toml:"max_velocity_y"`
}

// BallConfig holds ball physics parameters
type BallConfig struct {
	Size           float64 `toml:"size"`
	BaseSpeed      float64 `toml:"base_speed"`
	SpeedIncrement float64 `toml:"speed_increment"`
	BounceLimit    int     `toml:"bounce_limit"`
	MaxTrailLength int     `toml:"max_trail_length"`
}

// PaddleConfig holds paddle parameters
type PaddleConfig struct {
	Width      float64 `toml:"width"`
	Height     float64 `toml:"height"`
	Speed      float64 `toml:"speed"`
	SpinFactor float64 `toml:"spin_factor"`
}

// ExplosionConfig holds explosion effect parameters
type ExplosionConfig struct {
	ParticleCount int     `toml:"particle_count"`
	Lifetime      int     `toml:"lifetime"`
	BaseVelocity  float64 `toml:"base_velocity"`
}

// AIConfig holds AI difficulty parameters
type AIConfig struct {
	EasyDeadzone          float64 `toml:"easy_deadzone"`
	EasyReactionSpeed     float64 `toml:"easy_reaction_speed"`
	MediumDeadzone        float64 `toml:"medium_deadzone"`
	MediumReactionSpeed   float64 `toml:"medium_reaction_speed"`
	HardDeadzone          float64 `toml:"hard_deadzone"`
	HardReactionSpeed     float64 `toml:"hard_reaction_speed"`
	HardPredictionEnabled bool    `toml:"hard_prediction_enabled"`
}

// AudioConfig holds audio settings
type AudioConfig struct {
	Enabled bool `toml:"enabled"`
}

// VideoConfig holds video/display settings
type VideoConfig struct {
	Width      int  `toml:"width"`
	Height     int  `toml:"height"`
	Fullscreen bool `toml:"fullscreen"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Game: GameConfig{
			WinningScore: 5,
			MaxVelocityY: 8.0,
		},
		Ball: BallConfig{
			Size:           10.0,
			BaseSpeed:      4.0,
			SpeedIncrement: 0.3,
			BounceLimit:    10,
			MaxTrailLength: 20,
		},
		Paddle: PaddleConfig{
			Width:      20.0,
			Height:     100.0,
			Speed:      6.0,
			SpinFactor: 5.0,
		},
		Explosion: ExplosionConfig{
			ParticleCount: 20,
			Lifetime:      30,
			BaseVelocity:  3.0,
		},
		AI: AIConfig{
			EasyDeadzone:          30.0,
			EasyReactionSpeed:     0.5,
			MediumDeadzone:        10.0,
			MediumReactionSpeed:   1.0,
			HardDeadzone:          3.0,
			HardReactionSpeed:     1.0,
			HardPredictionEnabled: true,
		},
		Audio: AudioConfig{
			Enabled: true,
		},
		Video: VideoConfig{
			Width:      800,
			Height:     600,
			Fullscreen: false,
		},
	}
}

// LoadConfig loads configuration from file, falls back to defaults if error
func LoadConfig(path string) *Config {
	config := DefaultConfig()
	
	// Try to load from file
	data, err := os.ReadFile(path)
	if err != nil {
		// File doesn't exist or can't be read, use defaults
		return config
	}
	
	// Parse TOML
	if err := toml.Unmarshal(data, config); err != nil {
		// Parse error, use defaults
		return config
	}
	
	return config
}

// SaveConfig saves configuration to file
func SaveConfig(path string, config *Config) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := toml.NewEncoder(file)
	return encoder.Encode(config)
}
