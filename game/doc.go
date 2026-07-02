// Package game implements the Pong game with a clean architecture.
//
// Architecture Overview:
//
// CORE (Business Logic & Configuration):
//   - config.go      : Configuration management (TOML loading, defaults)
//   - constants.go   : Game constants (screen size, UI offsets, colors)
//   - events.go      : Event system for decoupling (EventBus, EventType)
//   - stats.go       : Statistics tracking and persistence (HighScores, Records)
//
// ENTITIES (Game Objects):
//   - ball.go        : Ball entity with physics and trail rendering
//   - paddle.go      : Paddle entity with movement and collision
//   - explosion.go   : Explosion particle effect entity
//
// SYSTEMS (Game Logic):
//   - game.go        : Core game loop and state management
//   - input.go       : Input handling for all game states
//   - scoring.go     : Score management and win conditions
//   - ai.go          : AI opponent with difficulty levels
//
// RENDERING (UI & Graphics):
//   - render.go      : All rendering logic (menu, game, UI overlays)
//   - utils.go       : Shared rendering utilities (drawCircle)
//
// AUDIO:
//   - audio.go       : Audio management with goroutine safety
//
// Design Patterns:
//   - Event-driven architecture for decoupling
//   - Single Responsibility Principle (each file has one purpose)
//   - Dependency Injection (Config, Stats, EventBus passed to Game)
//   - Observer pattern (EventBus for audio/visual effects)
package game
