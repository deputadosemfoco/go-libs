package domain

import "time"

// Entity is base struct for all domain entities
type Entity struct {
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
}
