package game

import (
	"encoding/json"
	"os"
	"time"
)

// Stats holds all game statistics
type Stats struct {
	// Overall stats
	TotalGamesPlayed   int     `json:"total_games_played"`
	TotalPlaytimeMinutes int     `json:"total_playtime_minutes"`
	
	// Player 1 stats
	Player1Wins        int     `json:"player1_wins"`
	Player1TotalScore  int     `json:"player1_total_score"`
	
	// Player 2 stats
	Player2Wins        int     `json:"player2_wins"`
	Player2TotalScore  int     `json:"player2_total_score"`
	
	// AI mode stats
	AIWins             int     `json:"ai_wins"`
	PlayerWins         int     `json:"player_wins"`
	
	// Records
	LongestRally       int     `json:"longest_rally"`
	ShortestGame       int     `json:"shortest_game_seconds"`     // Shortest winning game
	LongestGame        int     `json:"longest_game_seconds"`      // Longest winning game
	
	// High scores (per mode)
	HighScores         []HighScore `json:"high_scores"`
	
	// Last updated
	LastPlayed         time.Time   `json:"last_played"`
}

// HighScore represents a single high score entry
type HighScore struct {
	Mode           string    `json:"mode"`           // "2player" or "singleplayer"
	Winner         string    `json:"winner"`         // "Player 1", "Player 2", "You", "AI"
	FinalScore     string    `json:"final_score"`    // "5-3"
	Duration       int       `json:"duration"`       // seconds
	LongestRally   int       `json:"longest_rally"`  // hits
	Date           time.Time `json:"date"`
}

// DefaultStats returns a new Stats with default values
func DefaultStats() *Stats {
	return &Stats{
		HighScores:     []HighScore{},
		ShortestGame:   999999, // Initialize with very high value
		LastPlayed:     time.Now(),
	}
}

// LoadStats loads statistics from file
func LoadStats(path string) *Stats {
	stats := DefaultStats()
	
	data, err := os.ReadFile(path)
	if err != nil {
		// File doesn't exist, use defaults
		return stats
	}
	
	if err := json.Unmarshal(data, stats); err != nil {
		// Parse error, use defaults
		return stats
	}
	
	return stats
}

// SaveStats saves statistics to file
func SaveStats(path string, stats *Stats) error {
	// Update last played time
	stats.LastPlayed = time.Now()
	
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}

// RecordGame records a finished game in the statistics
func (s *Stats) RecordGame(mode GameMode, winner int, score1 int, score2 int, duration int, longestRally int) {
	s.TotalGamesPlayed++
	s.TotalPlaytimeMinutes += duration / 60
	
	// Update win counts
	if mode == GameModeTwoPlayer {
		if winner == 1 {
			s.Player1Wins++
		} else {
			s.Player2Wins++
		}
	} else {
		if winner == 1 {
			s.PlayerWins++
		} else {
			s.AIWins++
		}
	}
	
	// Update total scores
	s.Player1TotalScore += score1
	s.Player2TotalScore += score2
	
	// Update records
	if longestRally > s.LongestRally {
		s.LongestRally = longestRally
	}
	
	if duration < s.ShortestGame {
		s.ShortestGame = duration
	}
	
	if duration > s.LongestGame {
		s.LongestGame = duration
	}
	
	// Add to high scores (keep top 10)
	modeStr := "2player"
	winnerStr := "Player 1"
	if mode == GameModeSinglePlayer {
		modeStr = "singleplayer"
		if winner == 1 {
			winnerStr = "You"
		} else {
			winnerStr = "AI"
		}
	} else if winner == 2 {
		winnerStr = "Player 2"
	}
	
	finalScoreStr := ""
	if winner == 1 {
		finalScoreStr = jsonScore(score1, score2)
	} else {
		finalScoreStr = jsonScore(score2, score1)
	}
	
	highScore := HighScore{
		Mode:         modeStr,
		Winner:       winnerStr,
		FinalScore:   finalScoreStr,
		Duration:     duration,
		LongestRally: longestRally,
		Date:         time.Now(),
	}
	
	s.HighScores = append(s.HighScores, highScore)
	
	// Keep only the last 10 high scores
	if len(s.HighScores) > 10 {
		s.HighScores = s.HighScores[len(s.HighScores)-10:]
	}
}

// jsonScore formats a score as "5-3" string
func jsonScore(winner int, loser int) string {
	return string(rune('0'+winner)) + "-" + string(rune('0'+loser))
}
