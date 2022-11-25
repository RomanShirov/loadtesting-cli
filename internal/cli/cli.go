package cli

import (
	"fmt"
	"github.com/RomanShirov/loadtesting-cli/internal/loadtest"
	"github.com/manifoldco/promptui"
	"strconv"
	"time"
)

func SelectLoadConfigMode() (string, error) {
	confModeFlags := map[string]string{"From .env configuration file": "ENV", "Select manually": "MANUAL"}
	prompt := promptui.Select{
		Label: "Choose a load configuration mode",
		Items: []string{"From .env configuration file", "Select manually"},
	}

	_, result, err := prompt.Run()

	configMode := confModeFlags[result]

	if err != nil {
		fmt.Printf("Incorrect input error: %v\n", err)
		return "", err
	}

	return configMode, nil
}

func GatherConfigData() loadtest.Config {

	// Target URL
	targetPrompt := promptui.Prompt{
		Label: "Target URL",
	}
	targetURL, _ := targetPrompt.Run()

	// Request method
	reqMethodPrompt := promptui.Select{
		Label: "Request Method",
		Items: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}
	_, requestMethod, _ := reqMethodPrompt.Run()

	// Requests frequency
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		return err
	}
	reqFrequencyPrompt := promptui.Prompt{
		Label:    "Requests Frequency",
		Validate: validate,
	}
	reqFrequency, _ := reqFrequencyPrompt.Run()
	reqFrequencyCount, _ := strconv.Atoi(reqFrequency)

	// Request frequency time unit
	frequencyUnits := map[string]time.Duration{"RPS": time.Second, "RPM": time.Minute}
	freqTimeUnitPrompt := promptui.Select{
		Label: "Request per (Second, Minute)?",
		Items: []string{"RPS", "RPM"},
	}
	_, freqTimeUnit, _ := freqTimeUnitPrompt.Run()
	freqTimeUnitDuration := frequencyUnits[freqTimeUnit]

	// Test duration
	testDurationValidate := func(input string) error {
		_, err := strconv.Atoi(input)
		return err
	}
	testDurationPrompt := promptui.Prompt{
		Label:    "Test Duration (In seconds)",
		Validate: testDurationValidate,
	}
	testDuration, _ := testDurationPrompt.Run()
	testDurationTime, _ := strconv.Atoi(testDuration)

	config := loadtest.Config{
		TargetServiceURL:         targetURL,
		RequestMethod:            requestMethod,
		RequestFrequency:         reqFrequencyCount,
		FrequencyPerUnitDuration: freqTimeUnitDuration,
		TestDuration:             time.Duration(testDurationTime) * time.Second,
	}

	return config
}
