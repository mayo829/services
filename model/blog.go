package model

import (
	"time"
)

type Post struct {
	PostID     	uint64     `json:"post_id"`
	PostName 		string		 `json:"post_name"`
	PostDesc		string		 `json:"post_description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"created_at"`
}