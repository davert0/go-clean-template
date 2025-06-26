package request

type Entity struct {
	EntityID   string `json:"entity_id"   validate:"required"   example:"shot_12"`
	EntityType string `json:"entity_type" validate:"required"   example:"shot"`
	OrderBy    string `json:"order_by"                       example:"ASC"`
}
