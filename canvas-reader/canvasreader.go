package main

import (
	"fmt"
	"os"
	network "github.com/Go-Feedforward-Neural-Network/network"
	"strconv"
)



func main() {

	fmt.Printf("info:Starting the program...\n")

	var input []float32
	input = make([]float32, 784)

	if len(os.Args) <= 1 {
		fmt.Printf("error:This script can only be run with arguments, please fix.\n")
		return
		//input = []float32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,45,118,147,254,254,254,254,206,118,95,0,0,58,70,0,0,0,0,0,0,0,0,0,0,0,0,0,0,8,164,253,253,253,253,253,253,253,249,234,234,243,245,234,200,81,0,0,0,0,0,0,0,0,0,0,0,0,12,70,177,177,216,253,253,253,253,253,253,253,253,253,253,247,81,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,32,61,61,61,61,61,139,131,139,253,253,222,23,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,17,198,253,245,109,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,131,253,253,191,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,79,253,253,253,121,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,27,251,253,253,113,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,151,253,253,217,22,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,151,237,253,219,35,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,77,235,253,252,105,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,59,237,253,250,144,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,168,253,253,222,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,177,253,253,230,46,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,54,250,253,253,75,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,57,237,253,253,148,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,232,253,253,253,172,163,193,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,39,229,253,253,253,253,242,212,115,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,56,253,253,253,233,110,70,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,56,253,171,116,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	} else {

		fmt.Printf("info:Length of the input:%v\n", len(os.Args[1:]))

		// allocate each part of the argument into the input, but convert it from string first into float32
		for indexArg, arg := range os.Args[1:] {

			if n, err := strconv.ParseFloat(arg, 32); err == nil {
				input[indexArg] = float32(n)
			}
		}
	}

	// create the network
	json_file := "./files/trained_network.json"
	nn := network.LoadNetworkFromFile( json_file )

	// produce a prediction on a sample
	predictionVector := nn.Calculate(input)
	fmt.Printf("Prediction Vector:%v\n", predictionVector )

	predictionVectorNormalized := network.Predict(predictionVector)
	finalResult := GetInteger(predictionVectorNormalized)

	fmt.Printf("Prediction Vector Normalized:%v\n", predictionVectorNormalized )
	fmt.Printf("Final Result:%v\n", finalResult )

}

func GetInteger(input []float32) (prediction int) {

	for i,j := range input {
		if j == 1 {
			return i+1
		}
	}
	return
}
