package models

import "time"

type CompositionMetadata struct {
	CompositionID int
	Genre         string
	Tags          []string
	ListenCount   int
	LikeCount     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	Limit         int
	Offset        int
}
