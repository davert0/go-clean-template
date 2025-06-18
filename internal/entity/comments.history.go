// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// CommentsHistory -.
type CommentsHistory struct {
	Comments []Comment `json:"comment"`
}
