package game

// GameState represents the current state of the game
type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStateWaitingToStart
	GameStateGameOver
	GameStatePaused
	GameStateOptions
)

// GameMode represents the game mode (single or two player)
type GameMode int

const (
	GameModeTwoPlayer GameMode = iota
	GameModeSinglePlayer
)

// AIDifficulty represents the AI opponent difficulty level
type AIDifficulty int

const (
	AIDifficultyEasy AIDifficulty = iota
	AIDifficultyMedium
	AIDifficultyHard
)

// Game represents the main game structure with all game state
type Game struct {
	ball             *Ball
	paddle1          *Paddle
	paddle2          *Paddle
	score1           int
	score2           int
	gameState        GameState
	gameMode         GameMode
	winner           int
	menuSelection    int
	optionsSelection int
	audioManager     *AudioManager
	explosion        *Explosion
	isFullscreen     bool
	audioEnabled     bool
	aiDifficulty     AIDifficulty
	config           *Config   // Game configuration
	stats            *Stats    // Game statistics
	eventBus         *EventBus // Event bus for decoupling
	
	// Game stats
	gameStartTime    int64  // Frame count when game started
	currentRally     int    // Current rally length (consecutive hits)
	longestRally     int    // Longest rally in this game
	totalFrames      int    // Total frames played (for duration)
}

// NewGame creates and initializes a new game instance
func NewGame() *Game {
	// Load configuration
	config := LoadConfig("config.toml")
	
	// Load statistics
	stats := LoadStats("saves/stats.json")
	
	// Create event bus
	eventBus := NewEventBus()
	
	audioManager, err := NewAudioManager()
	if err != nil {
		// Audio initialization failed, but game can continue without audio
		audioManager = &AudioManager{}
	}
	
	game := &Game{
		ball:          NewBall(float64(config.Video.Width)/2, float64(config.Video.Height)/2),
		paddle1:       NewPaddle(30, float64(config.Video.Height)/2-config.Paddle.Height/2),
		paddle2:       NewPaddle(float64(config.Video.Width)-50, float64(config.Video.Height)/2-config.Paddle.Height/2),
		gameState:     GameStateMenu,
		gameMode:      GameModeTwoPlayer,
		audioManager:  audioManager,
		audioEnabled:  config.Audio.Enabled,
		aiDifficulty:  AIDifficultyMedium,
		config:        config,
		stats:         stats,
		eventBus:      eventBus,
		isFullscreen:  config.Video.Fullscreen,
	}
	
	// Subscribe to events
	game.setupEventHandlers()
	
	return game
}

// setupEventHandlers configures event listeners
func (g *Game) setupEventHandlers() {
	// Audio events
	g.eventBus.Subscribe(EventBallHitWall, func(e Event) {
		g.audioManager.PlayWallHit()
	})
	
	g.eventBus.Subscribe(EventBallHitPaddle, func(e Event) {
		g.audioManager.PlayPaddleHit()
	})
	
	g.eventBus.Subscribe(EventScorePoint, func(e Event) {
		if data, ok := e.Data.(ScorePointData); ok {
			g.explosion = NewExplosion(data.X, data.Y)
		}
		g.audioManager.PlayScore()
	})
}

// Update updates game logic (60 times per second)
func (g *Game) Update() error {
	g.handleInput()

	if g.explosion != nil && g.explosion.Active {
		g.explosion.Update()
	}

	if g.gameState == GameStatePlaying {
		// Track total frames played
		g.totalFrames++
		
		// Update ball and check for wall collision
		if g.ball.Update() {
			g.eventBus.Publish(Event{
				Type: EventBallHitWall,
				Data: BallHitWallData{
					X: g.ball.X + g.ball.Size/2,
					Y: g.ball.Y + g.ball.Size/2,
				},
			})
		}

		// Update paddles
		g.paddle1.Update()
		g.paddle2.Update()

		// Update AI in single player mode
		if g.gameMode == GameModeSinglePlayer {
			g.updateAI()
		}

		// Check paddle collisions
		if g.checkPaddleCollisions() {
			// Event is published inside checkPaddleCollisions
		}

		// Check scoring
		g.checkScoring()
	}

	return nil
}
