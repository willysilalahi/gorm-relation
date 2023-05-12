package model

type Tag struct {
	ID    int    `json:"id"`
	Nama  string `json:"name"`
	Posts []Post `json:"posts" gorm:"many2many:post_tags"`
}
