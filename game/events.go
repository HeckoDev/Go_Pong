package game

// EventType represents different types of game events
type EventType int

const (
	EventBallHitWall EventType = iota
	EventBallHitPaddle
	EventScorePoint
	EventGameWon
	EventGameStarted
	EventGamePaused
	EventGameResumed
)

// Event represents a game event with optional data
type Event struct {
	Type EventType
	Data interface{}
}

// BallHitWallData contains data for ball hitting wall
type BallHitWallData struct {
	X float64
	Y float64
}

// BallHitPaddleData contains data for ball hitting paddle
type BallHitPaddleData struct {
	X         float64
	Y         float64
	IsPlayer1 bool
}

// ScorePointData contains data for scoring
type ScorePointData struct {
	Player int // 1 or 2
	X      float64
	Y      float64
}

// GameWonData contains data for game won
type GameWonData struct {
	Winner int // 1 or 2
	Score1 int
	Score2 int
}

// EventBus handles event publishing and subscribing
type EventBus struct {
	listeners map[EventType][]func(Event)
}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[EventType][]func(Event)),
	}
}

// Subscribe adds a listener for a specific event type
func (eb *EventBus) Subscribe(eventType EventType, handler func(Event)) {
	eb.listeners[eventType] = append(eb.listeners[eventType], handler)
}

// Publish sends an event to all subscribed listeners
func (eb *EventBus) Publish(event Event) {
	if handlers, ok := eb.listeners[event.Type]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
