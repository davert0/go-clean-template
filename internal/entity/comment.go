// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// Comment -.
type Comment struct {
	Text       string    `json:"text"        example:"good job"`
	CreatedBy  string    `json:"created_by"  example:"user_kek"`
	CreatedAt  time.Time `json:"created_at"  example:"2025-06-15T12:34:56Z"`
	EntityID   string    `json:"entity_id"       example:"shot_12"`
	EntityType string    `json:"entity_type"     example:"shot"`
}
