package main

// #ifdef __cplusplus
// extern "C"
// {
// #endif
//     struct SamplesData
//     {
//         int numSamples;
//         int numInputs;
//         float *input;
//         float *output;
//     };
// #ifdef __cplusplus
// }
// #endif
import "C"

import (
	"fmt"
	"unsafe"
)

//export Hello
func Hello(data *C.struct_SamplesData) {
	// samples := make([][]C.float, data.numInputs)
	// for i := 0; i < int(data.numInputs); i++ {
	// 	// Get the pointer to the i-th row
	// 	rowPtr := (*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(data.input)) + uintptr(i)*unsafe.Sizeof(*data.input)))
	// 	// Create a Go slice for the row
	samples := (*[1 << 30]C.float)(unsafe.Pointer(data.input))[:data.numSamples:data.numSamples]
	// }
	// output := make([]C.float, data.numInputs)
	// for i := 0; i < int(data.numInputs); i++ {
	// 	// Get the pointer to the i-th row
	// 	rowPtr := (*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(data.output)) + uintptr(i)*unsafe.Sizeof(*data.output)))
	// 	// Create a Go slice for the row
	output := (*[1 << 30]C.float)(unsafe.Pointer(data.output))[:data.numSamples:data.numSamples]
	// }

	fmt.Printf("data from C: num samples %d, num inputs: %d\n", int(data.numSamples), int(data.numInputs))
	// fmt.Println(samples[0])
	for i, sample := range samples {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(sample)
		output[i] = C.float(float32(sample) * 2)
	}
	fmt.Print("\n")
}

func main() {
}
