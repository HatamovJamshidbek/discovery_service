package models

import "time"

type UserInteraction struct {
	ID              int
	UserID          int
	CompositionID   int
	InteractionType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	Limit           int
	Offset          int
}
