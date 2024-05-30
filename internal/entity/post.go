package entity

import (
	"github.com/lib/pq"
)

type Post struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Tags        []Tag       `json:"tags"`
	Status      string      `json:"status"`
	PublishDate pq.NullTime `json:"publish_date"`
}
