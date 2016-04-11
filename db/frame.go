package db

import (
	"time"
)

type Frame struct {
	Id           string    `json:"id"`
	Content      string    `json:"content"`
	Editor       string    `json:"editor"`
	FrameTaken   time.Time `json:"frameTaken"`
	FileId       string    `json:"fileId"`
	CurrentHash  string    `json:"currentHash"`
	PreviousHash string    `json:"previousHash"`
	InsertedAt   time.Time `json:"insertedAt"`
}

func BatchCreateFrames(frames []Frame) error {
	return nil
}
