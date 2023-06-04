package model

import (
	"github.com/joehewett/gollmapi/parameters"
)

type Model struct {
	name       ModelName
	mode       string
	parameters *parameters.ModelParameters
}

// New returns a new llm that can be used to generate text
func New(modelName ModelName) (*Model, error) {
	return &Model{
		name:       modelName,
		parameters: parameters.New(modelName),
	}, nil
}

// Complete returns a completed prompt
func (m *Model) Complete(prompt string) (string, error) {
	// Hit the openAI API:
	// response := openAI.complete(prompt, m.parameters)
	return "Completed text", nil

}

// Enum for models
type ModelName string

const (
	// Davinci is the Davinci model
	Davinci ModelName = "davinci"
	// Curie is the Curie model
	Curie ModelName = "curie"
	// Babbage is the Babbage model
	Babbage ModelName = "babbage"
	// Ada is the Ada model
	Ada ModelName = "ada"
	// DaVinciInstructBeta is the DaVinci Instruct Beta model
	DaVinciInstructBeta ModelName = "davinci-instruct-beta"
	// CurieInstructBeta is the Curie Instruct Beta model
	CurieInstructBeta ModelName = "curie-instruct-beta"
	// BabbageInstructBeta is the Babbage Instruct Beta model
	TextDaVinci003 ModelName = "text-davinci-003"
)
