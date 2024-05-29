// Your First C++ Program

#include "gohello.h"
#include <stdlib.h>
#include <random>
#include <iostream>

const int NUM_INPUTS = 2;
const int NUM_SAMPLES = 10;

int main()
{
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_real_distribution<float> dis(0.0, 1.0); // Distribution in range [0.0, 1.0)
    float *samples = new float[NUM_SAMPLES];
    for (int i = 0; i < NUM_SAMPLES; ++i)
    {
        samples[i] = i * 0.1;
    }

    std::cout << "The value of input is:" << std::endl;

    for (int j = 0; j < NUM_SAMPLES; ++j)
    {
        if (j != 0)
        {
            std::cout << ", ";
        }
        std::cout << samples[j];
    }
    std::cout << std::endl;

    float *output = new float[NUM_SAMPLES];
    for (int i = 0; i < NUM_SAMPLES; ++i)
    {
        output[i] = 0.0;
    }

    SamplesData samplesData = {
        NUM_SAMPLES,
        NUM_INPUTS,
        samples,
        output,
    };

    Hello(&samplesData);
    std::cout << "The value of output is:" << std::endl;

    for (int j = 0; j < NUM_SAMPLES; ++j)
    {
        if (j != 0)
        {
            std::cout << ", ";
        }
        std::cout << output[j];
    }
    std::cout << std::endl;
    return 0;
}