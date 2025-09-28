package entity

import "time"

// Event представляет сущность события
type Event struct {
	UserID      uint64
	Title       string
	Description string
	Date        time.Time
}

// Request представляет сущность события из запроса
type Request struct {
	UserID      uint64
	Title       string
	Description string
	Date        string
}
