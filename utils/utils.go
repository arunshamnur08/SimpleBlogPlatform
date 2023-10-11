package utils

import (
	"time"
)

type BlogDto struct {
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

type Blog struct {
	PostID    string    `json:"post_id" bson:"postid"`
	Title     string    `json:"title"  bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
}

type BlogDtoForUpdate struct {
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}
