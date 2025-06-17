// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// Comment -.
type Comment struct {
	EntityRefID   int64     `json:"entity_ref_id"`
	Text          string    `json:"text"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}
