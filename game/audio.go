package game

import (
	"log"
	"math"
	"os"
	"strings"
	"sync/atomic"
	"time"

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
	enabled         atomic.Bool
}

// isWSL checks if we're running under WSL
func isWSL() bool {
	// Check /proc/version for WSL
	if data, err := os.ReadFile("/proc/version"); err == nil {
		version := strings.ToLower(string(data))
		if strings.Contains(version, "microsoft") || strings.Contains(version, "wsl") {
			return true
		}
	}
	
	// Check WSL_DISTRO_NAME environment variable
	if os.Getenv("WSL_DISTRO_NAME") != "" {
		return true
	}
	
	return false
}

// NewAudioManager creates a new audio manager
// Returns an error if audio initialization fails
func NewAudioManager() (*AudioManager, error) {
	am := &AudioManager{}
	am.enabled.Store(false)
	
	// Disable audio on WSL by default (ALSA doesn't work well there)
	if isWSL() {
		log.Println("🔇 WSL detected: Audio disabled (use pong.exe for audio support)")
		return am, nil
	}
	
	audioContext := audio.NewContext(sampleRate)
	if audioContext == nil {
		log.Println("⚠️  Audio initialization failed: context is nil")
		return am, nil // Return disabled audio, not a fatal error
	}
	
	am.audioContext = audioContext
	am.paddleHitSound = generatePaddleHitSound()
	am.wallHitSound = generateWallHitSound()
	am.scoreSound = generateScoreSound()
	am.enabled.Store(true)
	log.Println("🔊 Audio initialized successfully")
	
	return am, nil
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
	if !am.enabled.Load() || am.audioContext == nil {
		return
	}
	
	// Protect against audio errors
	defer func() {
		if r := recover(); r != nil {
			// Disable audio if playback fails
			am.enabled.Store(false)
			log.Printf("⚠️  Audio playback failed, disabling audio: %v\n", r)
		}
	}()
	
	player := am.audioContext.NewPlayerFromBytes(am.paddleHitSound)
	if player == nil {
		return
	}
	player.Play()
	
	go func() {
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				if !player.IsPlaying() {
					player.Close()
					return
				}
			case <-timeout:
				player.Close()
				return
			}
		}
	}()
}

func (am *AudioManager) PlayWallHit() {
	if !am.enabled.Load() || am.audioContext == nil {
		return
	}
	
	// Protect against audio errors
	defer func() {
		if r := recover(); r != nil {
			// Disable audio if playback fails
			am.enabled.Store(false)
			log.Printf("⚠️  Audio playback failed, disabling audio: %v\n", r)
		}
	}()
	
	player := am.audioContext.NewPlayerFromBytes(am.wallHitSound)
	if player == nil {
		return
	}
	player.Play()
	
	go func() {
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				if !player.IsPlaying() {
					player.Close()
					return
				}
			case <-timeout:
				player.Close()
				return
			}
		}
	}()
}

func (am *AudioManager) PlayScore() {
	if !am.enabled.Load() || am.audioContext == nil {
		return
	}
	
	// Protect against audio errors
	defer func() {
		if r := recover(); r != nil {
			// Disable audio if playback fails
			am.enabled.Store(false)
			log.Printf("⚠️  Audio playback failed, disabling audio: %v\n", r)
		}
	}()
	
	player := am.audioContext.NewPlayerFromBytes(am.scoreSound)
	if player == nil {
		return
	}
	player.Play()
	
	go func() {
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				if !player.IsPlaying() {
					player.Close()
					return
				}
			case <-timeout:
				player.Close()
				return
			}
		}
	}()
}
