package request

import "time"

type Comment struct {
	Text       string    `json:"text"        validate:"required"    example:"good job"`
	CreatedBy  string    `json:"created_by"  validate:"required"   example:"user_kek"`
	EntityID   string    `json:"entity_id"   validate:"required"  example:"shot_12"`
	EntityType string    `json:"entity_type" validate:"required"   example:"shot"`
	CreatedAt  time.Time `json:"created_at"  example:"2025-06-15T12:34:56Z"`
}
