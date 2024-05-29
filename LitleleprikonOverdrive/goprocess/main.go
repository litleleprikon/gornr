package main

// #ifdef __cplusplus
// extern "C"
// {
// #endif
//     struct OverdriveProcessData
//     {
//         int numSamples;
//         double gain;
//         double volume;
//         double tone;
//         double sampleRate;
//         float *input;
//         float *output;
//     };
// #ifdef __cplusplus
// }
// #endif
import "C"
import (
	"math"
	"unsafe"
)

func main() {

}

const (
	ResultCodeOk int = iota
)

// func init() {
// 	debug.SetGCPercent(-1)
// 	debug.SetMemoryLimit(math.MaxInt64)
// }

//export ProcessOverdrive
func ProcessOverdrive(data *C.struct_OverdriveProcessData) C.int {
	numSamples := int(data.numSamples)
	gain := float32(data.gain)
	volume := float32(data.volume)
	tone := float32(data.tone)
	sampleRate := int(data.sampleRate)

	// Convert input and output pointers to Go slices
	input := (*[1 << 30]float32)(unsafe.Pointer(data.input))[:numSamples:numSamples]
	output := (*[1 << 30]float32)(unsafe.Pointer(data.output))[:numSamples:numSamples]

	for i := 0; i < numSamples; i++ {
		output[i] = processTS(input[i], gain, volume, tone, sampleRate)
	}

	return C.int(ResultCodeOk)
}

// ProcessTS applies the Tube Screamer effect to an input signal
func processTS(input float32, gain float32, volume float32, tone float32, sampleRate int) float32 {
	// Calculate overdrive amount
	overdrive := input * gain

	// Apply tone control
	overdrive = applyTone(overdrive, tone, gain, sampleRate)

	// Apply soft clipping
	overdrive = float32(softClip(float64(overdrive)))

	// Apply output level control
	output := volume * overdrive

	return output
}

// applyTone applies the tone control to the overdrive signal
func applyTone(input float32, tone float32, gain float32, sampleRate int) float32 {
	// Apply a simple high-pass filter for the tone control
	cutoff := 1000.0 * math.Pow(10, float64(2*tone-1))
	rc := 1.0 / (cutoff * 2 * math.Pi)
	dt := 1.0 / float32(sampleRate)
	alpha := dt / (float32(rc) + dt)
	return (1-alpha)*input + alpha*gain
	// panic("not implemented")
}

// softClip applies soft clipping to the input signal
func softClip(input float64) float64 {
	// const threshold1 = 0.5
	// const threshold2 = 0.75
	// const a = 1.5
	// const b = 0.75
	// const c = 0.25
	// sign := input / math.Abs(input)

	// output := input

	// if input > threshold2 {
	// 	output = sign*(a*input)/(1+math.Exp(b*(input-threshold2))) + c
	// } else if input > threshold1 {
	// 	output = sign * (a * input) / (1 + math.Exp(b*(input-threshold1)))
	// }
	// return output

	// const threshold = 0.3

	// // Apply cubic soft clipping
	// if input > threshold {
	// 	return threshold + (1-threshold)*math.Pow(input-threshold, 1/3.0)
	// } else if input < -threshold {
	// 	return -threshold + (1-threshold)*math.Pow(input+threshold, 1/3.0)
	// }
	// return input
	const gain = 2.0 // Gain factor before clipping, adjust as needed

	// Apply the gain to the input signal
	amplifiedInput := input * gain

	// Apply soft clipping using the tanh function
	clippedOutput := math.Tanh(amplifiedInput)

	return clippedOutput
}
