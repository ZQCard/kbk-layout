package domain

import "time"

type Example struct {
	Id        int64
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
