package loadtest

import "time"

type Config struct {
	TargetURL                string
	RequestMethod            string
	RequestFrequency         int
	FrequencyPerUnitDuration time.Duration
	TestDuration             time.Duration
}
