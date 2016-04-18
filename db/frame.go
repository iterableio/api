package db

import (
	"time"
)

type Frame struct {
	Id           string    `json:"id" column:"id"`
	Content      string    `json:"content" column:"content"`
	Editor       string    `json:"editor" column:"editor"`
	FrameTaken   time.Time `json:"frameTaken" column:"frame_taken"`
	FileId       string    `json:"fileId" column:"file_id"`
	CurrentHash  string    `json:"currentHash" column:"current_hash"`
	PreviousHash string    `json:"previousHash" column:"previous_hash"`
	InsertedAt   time.Time `json:"insertedAt" column:"inserted_at"`
}

func createColvaluesFromFrames(frames []Frame) map[string]interface{} {
	return nil
}

func BatchCreateFrames(frames []Frame) error {
	return nil
}
