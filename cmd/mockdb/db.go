package mockdb

import (
	"sync"
	"time"
)

// Match represents a ping pong match score.
type Match struct {
	Player1 string
	Player2 string
	Score1  int
	Score2  int
	Date    time.Time
}

// MockDB simulates a database for storing match scores.
type MockDB struct {
	matches []Match
	mu      sync.Mutex
}

var instance *MockDB = &MockDB{}

func GetDB() *MockDB {
	return instance
}

// New creates a new mock database.
func New() *MockDB {
	return &MockDB{
		matches: []Match{},
	}
}

// AddMatch adds a new match to the database.
func (db *MockDB) AddMatch(match Match) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.matches = append(db.matches, match)
}

// GetMatches returns all matches from the database.
func (db *MockDB) GetMatches() []Match {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.matches
}
