package models

import "time"

type URL struct {
	ID          int
	ShortURL    string
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	VisitCount  int
}
