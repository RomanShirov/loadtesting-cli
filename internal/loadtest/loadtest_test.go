package loadtest

import (
	"testing"
	"time"
)

func TestRunLoadTest(t *testing.T) {
	debugConfig := Config{
		"http://localhost:8080",
		"GET",
		1000,
		time.Second,
		30 * time.Second,
	}

	RunLoadTest(debugConfig)
}
