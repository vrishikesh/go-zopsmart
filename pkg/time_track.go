package pkg

import (
	"log"
	"time"
)

func TimeTrack(start time.Time) {
	log.Printf("Execution time %s", time.Since(start))
}
