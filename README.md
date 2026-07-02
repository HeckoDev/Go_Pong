# Pong Master

A modern, feature-rich Pong game implementation in Go using the Ebiten game engine.

## Features

### Game Modes
- 🎮 **Two-player mode** - Play against a friend on the same keyboard
- 🤖 **Single-player mode** - Play against an AI opponent

### Visual Effects
- 🎨 **Round ball** with smooth circular rendering
- 🔴 **Dynamic red trail** - Length increases with ball speed
- 💥 **Explosion effects** on scoring with particle system
- ⚡ **60 FPS** smooth rendering
- 🏓 **Clean interface** with center line and score display

### Gameplay Mechanics
- 🚀 **Progressive difficulty** - Ball speed increases every 5 paddle bounces
- 🎯 **Precise collision detection** with paddle spin effect
- � **Score system** - First to 5 points wins
- 🎲 **Random ball direction** on each reset
- 🎮 **Smooth paddle physics** with acceleration and friction

### Audio
- 🔊 **Paddle hit sounds** - High-pitched beep (440 Hz)
- � **Wall hit sounds** - Lower beep (330 Hz)
- 🔊 **Score sounds** - Dual-tone celebration

## Controls

### Player 1 (Left paddle)
- **W** : Move up
- **S** : Move down

### Player 2 (Right paddle)
- **↑** (Arrow Up) : Move up
- **↓** (Arrow Down) : Move down

### Other controls
- **Enter** : Start game / Launch ball / Continue after score
- **ESC** : Return to menu
- **W/S or ↑/↓** : Navigate menu

## Game Rules

1. **Objective**: Score points by getting the ball past the opponent's paddle
2. **Movement**: Paddles move with smooth acceleration and deceleration
3. **Ball Physics**: 
   - Bounces off paddles and top/bottom walls
   - Speed increases every 5 paddle hits
   - Paddle position affects ball angle (spin effect)
   - Speed resets after each point
4. **Scoring**: Ball exits left/right side = point for opponent
5. **Winning**: First player to reach 5 points wins
6. **Explosions**: Ball explodes into particles when a point is scored

## Installation

### Prerequisites
- **Go 1.21+**
- **Linux**: X11 and ALSA development libraries

### 🪟 WSL (Windows Subsystem for Linux) - **RECOMMANDÉ**

Si vous êtes sur WSL, utilisez la version Windows pour un fonctionnement optimal :

```bash
# Configurer le proxy Go (nécessaire une seule fois)
export GOPROXY="https://proxy.golang.org,direct"

# Compiler pour Windows
make build-windows

# Lancer le jeu
./pong.exe

# OU créer un alias pour lancer depuis n'importe où
echo "alias pong='cd /chemin/vers/PongMaster && ./pong.exe'" >> ~/.zshrc
source ~/.zshrc
pong
```

> **Note** : L'audio fonctionne parfaitement avec `pong.exe` sous WSL/Windows.

### 🐧 Linux Natif

#### Dépendances (Debian/Ubuntu)
```bash
sudo apt-get update
sudo apt-get install -y libx11-dev libxcursor-dev libxrandr-dev \
    libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev \
    libasound2-dev libasound2 libasound2-plugins alsa-utils
```

#### Démarrage rapide
```bash
# Installer les dépendances Go
export GOPROXY="https://proxy.golang.org,direct"
go mod tidy

# Build et lancement (recommandé)
make run

# OU compilation seule
make build

# OU exécution directe sans compilation
make dev
```

## Using Makefile

The project includes a comprehensive Makefile for easy project management:

```bash
make build         # Compile the game
make run           # Compile and run the game
make dev           # Run without compiling (go run)
make clean         # Remove compiled binaries
make deps          # Install Go dependencies
make install-deps  # Install system dependencies (Linux)
make build-linux   # Cross-compile for Linux
make build-windows # Cross-compile for Windows
make build-mac     # Cross-compile for macOS
make build-all     # Build for all platforms
make help          # Show all available commands
```

## Project Structure

```
PongMaster/
├── main.go              # Entry point
├── Makefile            # Build automation
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── README.md           # Documentation
└── game/               # Game logic package
    ├── game.go         # Core game state and logic
    ├── ball.go         # Ball physics and rendering
    ├── paddle.go       # Paddle physics and rendering
    ├── audio.go        # Sound generation and playback
    └── explosion.go    # Particle explosion effects
```

## Code Architecture

### Clean Code Principles
- **Single Responsibility**: Each file handles one aspect of the game
- **DRY**: Reusable functions for common operations (circle drawing, beep generation)
- **Performance**: Optimized rendering with pre-calculated values
- **Simplicity**: Clear, concise code without unnecessary comments

### Package: `game`

#### `game.go`
- Game state management (menu, playing, game over)
- Input handling for both players and menu navigation
- Game modes (two-player, single-player with AI)
- Collision detection and scoring system
- Rendering pipeline

#### `ball.go`
- Ball physics with progressive speed system
- Visual trail effect that grows with speed
- Circular rendering with optimized algorithms
- Bounce count tracking for difficulty scaling

#### `paddle.go`
- Smooth physics with acceleration and friction
- Velocity-based movement for natural feel
- Boundary collision handling

#### `audio.go`
- Procedural sound generation (sine wave synthesis)
- Three distinct sounds (paddle hit, wall hit, score)
- Graceful degradation if audio unavailable

#### `explosion.go`
- Particle system with 30 particles per explosion
- Physics simulation (gravity, friction)
- Alpha fading for smooth disappearance

## Game States

1. **Menu**: Choose between two-player or single-player mode
2. **Waiting to Start**: Initial state and after each point
3. **Playing**: Active gameplay with ball movement
4. **Game Over**: When a player reaches 5 points

## Performance Optimizations

- Pre-calculated radius squared for circle rendering
- Efficient trail management with pre-allocated slices
- Minimal memory allocations during gameplay
- Optimized collision detection
- Sound buffer reuse to avoid repeated player creation

## Technologies

- **Go 1.24**: Modern, fast, compiled language
- **Ebiten v2.9**: Cross-platform 2D game engine
- **Ebiten Audio**: Audio playback system
- **Pure Go**: No external C dependencies beyond system libraries

## Development

### Code Quality
- Clean, idiomatic Go code
- No unnecessary comments
- Consistent naming conventions
- Single responsibility principle
- DRY (Don't Repeat Yourself)

### Building
```bash
# Development build
go build -o pong

# Optimized release build
go build -ldflags="-s -w" -o pong
```

## Troubleshooting

### 🚨 Erreur "401 Unauthorized" lors du build

**Symptôme** : `github.com/ebitengine/oto/v3@v3.4.0: reading https://europe-west1-go.pkg.dev/... 401 Unauthorized`

**Solution** : Configurer le proxy Go public
```bash
export GOPROXY="https://proxy.golang.org,direct"
go clean -modcache
go mod download
```

Pour rendre permanent, ajoutez à votre `~/.zshrc` ou `~/.bashrc` :
```bash
export GOPROXY="https://proxy.golang.org,direct"
```

### 🐧 WSL : Erreurs ALSA (Audio)

**Symptôme** : `ALSA lib pcm.c: Unknown PCM default`

**Solution recommandée** : Utilisez la version Windows
```bash
make build-windows
./pong.exe
```

**Alternative** : Désactiver l'audio (Linux WSL)
```bash
export SDL_AUDIODRIVER=dummy
./pong
```

### 🖼️ WSL : Erreurs X11/Display

**Symptôme** : `Could not get authority info: open /home/user/.Xauthority: no such file or directory`

**Solution** : Mettre à jour WSL pour avoir WSLg
```powershell
# Depuis PowerShell (Windows)
wsl --update
wsl --shutdown
```

Vérifier que WSLg est installé :
```bash
wsl.exe --version
```

### 🔇 Linux Natif : Pas de son

Si vous n'avez pas de son sous Linux :
```bash
sudo apt-get install libasound2 libasound2-plugins alsa-utils
```

### 📦 Linux Natif : Erreurs X11

Si vous avez des erreurs de bibliothèques X11 :
```bash
sudo apt-get install libx11-dev libxcursor-dev libxrandr-dev \
    libxinerama-dev libxi-dev libgl1-mesa-dev libxxf86vm-dev
```
