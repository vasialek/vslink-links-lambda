package models

import "time"

// Link is model for create/transfer links
type Link struct {
	LinkID         string    `json:"link_id"`
	UserID         string    `json:"user_id"`
	TypeID         int       `json:"type_id"`
	LinkCategoryID string    `json:"link_category_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	URL            string    `json:"url"`
	Rating         int       `json:"rating"`
	CreatedAt      time.Time `json:"created_at"`
}
