package canvas

import (
	"fmt"
	"os"
	"github.com/Go-Feedforward-Neural-Network/network"
	"strconv"
	"net/http"
)


func CanvasHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}

func canvas() {
	fmt.Printf("info:Starting the program...\n")

	var input []float32
	input = make([]float32, 784)

	if len(os.Args) <= 1 {
		fmt.Printf("error:This script can only be run with arguments, please fix.\n")
		return
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
	return -1
}
