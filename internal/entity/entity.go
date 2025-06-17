// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Comment -.
type Entity struct {
	EntityID   string `json:"entity_id"       example:"shot_12"`
	EntityType string `json:"entity_type"     example:"shot"`
}
