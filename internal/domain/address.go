package domain

import "time"

type Address struct {
	ID        int64
	UserID    int64
	Street    string
	City      string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
