package config

import (
	"github.com/RomanShirov/loadtesting-cli/internal/loadtest"
	"log"
	"os"
	"strconv"
	"time"
)

func GetFrequencyTimeUnit() time.Duration {
	switch os.Getenv("FREQ_TIME_UNIT") {
	case "SECOND":
		return time.Second
	case "MINUTE":
		return time.Minute
	default:
		return time.Second
	}
}

func GetConfigFromEnvironment() loadtest.Config {
	targetURL := os.Getenv("TARGET_URL")
	requestMethod := os.Getenv("REQUEST_METHOD")
	requestFrequency, err := strconv.Atoi(os.Getenv("REQ_FREQUENCY"))
	if err != nil {
		log.Fatalf(".env parsing error: %v", err)
	}

	frequencyTimeUnit := GetFrequencyTimeUnit()

	testDurationEnv, err := strconv.Atoi(os.Getenv("TEST_DURATION"))
	if err != nil {
		log.Fatalf(".env parsing error: %v", err)
	}
	testDurationTime := time.Duration(testDurationEnv) * time.Second

	config := loadtest.Config{
		TargetServiceURL:         targetURL,
		RequestMethod:            requestMethod,
		RequestFrequency:         requestFrequency,
		FrequencyPerUnitDuration: frequencyTimeUnit,
		TestDuration:             testDurationTime,
	}

	return config
}
