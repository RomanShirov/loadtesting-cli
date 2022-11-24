package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func SelectLoadConfigMode() (string, error) {
	prompt := promptui.Select{
		Label: "Choose a load configuration mode",
		Items: []string{"From .env configuration file", "Select manually"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Incorrect input error: %v\n", err)
		return "", err
	}

	return result, nil
}
