package loadtest

import (
	"fmt"
	v "github.com/tsenart/vegeta/lib"
	"log"
	"os"
	"time"
)

type Config struct {
	TargetServiceURL         string
	RequestMethod            string
	RequestFrequency         int
	FrequencyPerUnitDuration time.Duration
	TestDuration             time.Duration
}

var (
	outputDelimiter = "\n------------\n\n"
)

func debugTarget(service string, method string) v.Target {
	return v.Target{
		URL:    service,
		Method: method,
	}
}

func handleReporterError(err error) {
	if err != nil {
		log.Fatalf("Reporting error: %v", err)
	}
}

func RunLoadTest(c Config) {
	rate := v.Rate{Freq: c.RequestFrequency, Per: c.FrequencyPerUnitDuration}
	duration := c.TestDuration
	targeter := v.NewStaticTargeter(
		debugTarget(c.TargetServiceURL, c.RequestMethod),
	)

	attacker := v.NewAttacker()
	var metrics v.Metrics
	histogram := v.Histogram{
		Buckets: []time.Duration{
			0,
			10 * time.Millisecond,
			25 * time.Millisecond,
			50 * time.Millisecond,
			100 * time.Millisecond,
			1000 * time.Millisecond,
			3000 * time.Millisecond,
			5000 * time.Millisecond,
			10000 * time.Millisecond,
		},
	}

	fmt.Printf("Starting load test...\n\n")
	fmt.Printf("\u001B[32mREQUEST_URL: %s \nREQUEST_METHOD: %s \nTEST_DURATION: %s \nREQUEST_FREQUENCY: %d / %s\033[0m \n",
		c.TargetServiceURL, c.RequestMethod, c.TestDuration, c.RequestFrequency, c.FrequencyPerUnitDuration)

	for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
		metrics.Add(res)
		histogram.Add(res)
	}
	metrics.Close()

	fmt.Printf(outputDelimiter)

	// ToDo: Add optional reporters output to JSON

	textReporter := v.NewTextReporter(&metrics)
	err := textReporter(os.Stdout)
	handleReporterError(err)

	fmt.Printf(outputDelimiter)

	histogramReporter := v.NewHistogramReporter(&histogram)
	err = histogramReporter(os.Stdout)
	handleReporterError(err)
}
