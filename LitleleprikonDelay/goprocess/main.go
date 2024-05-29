package main

// #ifdef __cplusplus
// extern "C"
// {
// #endif
//     struct DelayProcessData
//     {
//         int numSamples;
//         double delaySize;
//         double mix;
//         double sampleRate;
//         float *input;
//         float *output;
//     };
// #ifdef __cplusplus
// }
// #endif
import "C"
import (
	"unsafe"
)

func main() {
	numSamples := 10
	input := make([]float32, numSamples)
	output := make([]float32, numSamples)
	counter := 0
	for i := 0; i < 20; i++ {
		for j := 0; j < numSamples; j++ {
			input[j] = float32(counter) * 1e-6
			counter++
		}
		applyDelay(numSamples, 10, 1, input, output)
	}
}

const (
	ResultCodeOk int = iota
	ResultCodeDelayTooBig
)

// func init() {
// 	debug.SetGCPercent(-1)
// 	debug.SetMemoryLimit(math.MaxInt64)
// }

const delaySecondsLimit = 5

const sampleRateLimit = 100000
const maxBufferSize = delaySecondsLimit * sampleRateLimit

var writeIndex = 0

var buffer = [maxBufferSize]float32{}

//export ProcessDelay
func ProcessDelay(data *C.struct_DelayProcessData) C.int {
	// Calculate delay samples
	delaySamples := int(data.delaySize * data.sampleRate)
	if delaySamples > maxBufferSize {
		return C.int(ResultCodeDelayTooBig)
	}

	numSamples := int(data.numSamples)
	mix := float32(data.mix)

	// Convert input and output pointers to Go slices
	input := (*[1 << 30]float32)(unsafe.Pointer(data.input))[:numSamples:numSamples]
	output := (*[1 << 30]float32)(unsafe.Pointer(data.output))[:numSamples:numSamples]

	applyDelay(numSamples, delaySamples, mix, input, output)

	return C.int(ResultCodeOk)
}

func applyDelay(numSamples, delaySamples int, mix float32, input, output []float32) {
	for i := 0; i < numSamples; i++ {
		// Write current sample to the buffer
		buffer[writeIndex] = input[i]

		// Find delayed sample
		readIndex := (writeIndex + maxBufferSize - delaySamples) % maxBufferSize
		delayedSample := buffer[readIndex]

		// Apply delay with mix
		output[i] = input[i]*1 - mix + delayedSample*mix

		// Increment writeIndex and wrap around if necessary
		writeIndex = (writeIndex + 1) % maxBufferSize
	}
}
