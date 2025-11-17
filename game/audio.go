package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sampleRate = 48000
)

// AudioManager manages all game sounds
type AudioManager struct {
	audioContext    *audio.Context
	paddleHitSound  []byte
	wallHitSound    []byte
	scoreSound      []byte
	enabled         bool
}

// NewAudioManager creates a new audio manager
func NewAudioManager() *AudioManager {
	am := &AudioManager{
		enabled: true,
	}
	
	// Try to initialize audio, but don't crash if it fails
	defer func() {
		if r := recover(); r != nil {
			am.enabled = false
		}
	}()
	
	audioContext := audio.NewContext(sampleRate)
	
	am.audioContext = audioContext
	am.paddleHitSound = generatePaddleHitSound()
	am.wallHitSound = generateWallHitSound()
	am.scoreSound = generateScoreSound()
	
	return am
}

func generatePaddleHitSound() []byte {
	return generateBeep(440.0, 0.08)
}

func generateWallHitSound() []byte {
	return generateBeep(330.0, 0.06)
}

func generateBeep(frequency, duration float64) []byte {
	length := int(sampleRate * duration)
	pcm := make([]byte, length*4)
	
	for i := 0; i < length; i++ {
		t := float64(i) / sampleRate
		sample := math.Sin(2 * math.Pi * frequency * t)
		
		envelope := 1.0
		fadeIn := length / 10
		fadeOut := length * 9 / 10
		
		if i < fadeIn {
			envelope = float64(i) / float64(fadeIn)
		} else if i > fadeOut {
			envelope = float64(length-i) / float64(length-fadeOut)
		}
		
		sample *= envelope * 0.3
		val := int16(sample * 32767)
		
		pcm[i*4] = byte(val)
		pcm[i*4+1] = byte(val >> 8)
		pcm[i*4+2] = byte(val)
		pcm[i*4+3] = byte(val >> 8)
	}
	
	return pcm
}

func generateScoreSound() []byte {
	tone1 := generateBeep(523.25, 0.1)
	tone2 := generateBeep(392.00, 0.15)
	
	pcm := make([]byte, len(tone1)+len(tone2))
	copy(pcm, tone1)
	copy(pcm[len(tone1):], tone2)
	
	return pcm
}

func (am *AudioManager) PlayPaddleHit() {
	if !am.enabled || am.audioContext == nil {
		return
	}
	player := am.audioContext.NewPlayerFromBytes(am.paddleHitSound)
	player.Play()
}

func (am *AudioManager) PlayWallHit() {
	if !am.enabled || am.audioContext == nil {
		return
	}
	player := am.audioContext.NewPlayerFromBytes(am.wallHitSound)
	player.Play()
}

func (am *AudioManager) PlayScore() {
	if !am.enabled || am.audioContext == nil {
		return
	}
	player := am.audioContext.NewPlayerFromBytes(am.scoreSound)
	player.Play()
}
