package game

// GameState represents the current state of the game
type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStateWaitingToStart
	GameStateGameOver
)

// GameMode represents the game mode (single or two player)
type GameMode int

const (
	GameModeTwoPlayer GameMode = iota
	GameModeSinglePlayer
)

// Game represents the main game structure with all game state
type Game struct {
	ball          *Ball
	paddle1       *Paddle
	paddle2       *Paddle
	score1        int
	score2        int
	gameState     GameState
	gameMode      GameMode
	winner        int
	menuSelection int
	audioManager  *AudioManager
	explosion     *Explosion
}

// NewGame creates and initializes a new game instance
func NewGame() *Game {
	audioManager, err := NewAudioManager()
	if err != nil {
		// Audio initialization failed, but game can continue without audio
		audioManager = &AudioManager{}
	}
	
	return &Game{
		ball:          NewBall(ScreenWidth/2, ScreenHeight/2),
		paddle1:       NewPaddle(30, ScreenHeight/2-50),
		paddle2:       NewPaddle(ScreenWidth-50, ScreenHeight/2-50),
		gameState:     GameStateMenu,
		gameMode:      GameModeTwoPlayer,
		audioManager:  audioManager,
	}
}

// Update updates game logic (60 times per second)
func (g *Game) Update() error {
	g.handleInput()

	if g.explosion != nil && g.explosion.Active {
		g.explosion.Update()
	}

	if g.gameState == GameStatePlaying {
		// Update ball and check for wall collision
		if g.ball.Update() {
			g.audioManager.PlayWallHit()
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
			g.audioManager.PlayPaddleHit()
		}

		// Check scoring
		g.checkScoring()
	}

	return nil
}
