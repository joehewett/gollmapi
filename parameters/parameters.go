package parameters

import (
	"fmt"

	"github.com/joehewett/gollmapi/model"
)

// ErrModelNotSupported is returned when the model is not supported
var ErrModelNotSupported = fmt.Errorf("model not supported")

func New(modelName model.ModelName) *ModelParameters {
	params := &ModelParameters{
		temperature:       1.0,
		top_p:             1.0,
		top_k:             40,
		presence_penalty:  0.0,
		frequency_penalty: 0.0,
		stop_token:        "",
	}

	return params
}

type ModelParameters struct {
	temperature       float64
	top_p             float64
	top_k             int
	presence_penalty  float64
	frequency_penalty float64
	stop_token        string
}
