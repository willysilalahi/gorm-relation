package model

type PostTag struct {
	ID     int  `json:"id"`
	TagID  int  `json:"tag_id"`
	Tag    Tag  `json:"tag" gorm:"foreignKey:TagID"`
	PostID int  `json:"post_id"`
	Post   Post `json:"post" gorm:"foreignKey:PostID"`
}

func (PostTag) TableName() string {
	return "post_tags"
}
