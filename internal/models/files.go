package models

import "time"

type File struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Filename  string    `json:"filename"`
	Filepath  string    `json:"filepath"`
	Filesize  int64     `json:"filesize"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
