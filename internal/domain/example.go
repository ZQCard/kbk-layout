package domain

import "time"

type Example struct {
	ID        int64
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
