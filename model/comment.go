package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	AdminID   int       `json:"admin_id"`
	Admin     Admin     `json:"admin" gorm:"foreignKey:AdminID"`
	PostID    int       `json:"post_id"`
	Post      Post      `json:"post" gorm:"foreignKey:PostID"`
	CreatedAt time.Time `json:"created_at"`
}
