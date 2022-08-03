package model

import "time"

type Event struct {
	EventID     int       `json:"event_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
