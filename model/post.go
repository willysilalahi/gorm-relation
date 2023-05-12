package model

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Slug        string    `json:"slug"`
	AdminID     int       `json:"admin_id"`
	Tags        []Tag     `gorm:"many2many:post_tags"`
	Admin       Admin     `json:"admin" gorm:"foreignKey:AdminID"`
	Comments    []Comment `json:"comments" gorm:"foreignKey:PostID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_At"`
}
