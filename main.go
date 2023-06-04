package main

import (
	"fmt"

	"github.com/joehewett/gollmapi/model"
	"github.com/joehewett/gollmapi/prompt"
)

func main() {
	fmt.Printf("Hello World")

	model, err := model.New(model.TextDaVinci003)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	prompt, err := prompt.New("Once upon a time")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	response, err := model.Complete(prompt)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Response: %s", response)
}
