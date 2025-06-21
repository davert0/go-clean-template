package request

type Comment struct {
	Text       string `json:"text"        validate:"required"   example:"good job"`
	CreatedBy  string `json:"created_by"  validate:"required"   example:"user_kek"`
	EntityID   string `json:"entity_id"   validate:"required"   example:"shot_12"`
	EntityType string `json:"entity_type" validate:"required"   example:"shot"`
}
