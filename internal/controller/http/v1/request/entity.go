package request

type Entity struct {
	EntityID   string `json:"entity_id"   validate:"required"   example:"shot_12"`
	EntityType string `json:"entity_type" validate:"required"   example:"shot"`
	Offset     int    `json:"offset"                            example:"20"`
	Limit      int    `json:"limit"                             example:"10"`
	OrderBy    string `json:"order_by"                          example:"ASC"`
}
