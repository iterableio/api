package core

import (
	"github.com/iterableio/api/db"
)

func startPipeline(frames []db.Frame) {
	db.BatchCreateFrames(frames)
}
