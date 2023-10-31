package vm

import (
	"errors"
	"fmt"
	"unsafe"

)

type GPUContext struct {
	context *gocu.Context
}

// Initializes and returns a new GPU context
func NewGPUContext() (*GPUContext, error) {
	ctx, err := gocu.NewContext(gocu.Device(0), gocu.SchedAuto)
	if err != nil {
		return nil, err
	}

	return &GPUContext{
		context: ctx,
	}, nil
}

// RunInference runs the inference on the GPU
func (gc *GPUContext) RunInference(data []float32, modelPath string) ([]float32, error) {
	if gc.context == nil {
		return nil, errors.New("GPU context is not initialized")
	}

	// TODO: Generric model loading logic TBD

	// Allocate device memory and copy the data to the GPU
	dData, err := gocu.BytesToDevice(data)
	if err != nil {
		return nil, err
	}
	defer gocu.Free(dData)

	// TODO: We'll have to implement model execution here later

	// Copy the results back from the GPU
	res := make([]float32, len(data)) // Adjust the length based on your model's output
	err = gocu.DeviceToHost(res, dData)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// release resouces
func (gc *GPUContext) Cleanup() {
	if gc.context != nil {
		gc.context.Destroy()
	}
}
