package canvasComp

import (
   "fmt"
   "io"
   "io/ioutil"
   "strconv"
   "sort"
   "net/http"
   "encoding/json"
   "github.com/Go-Feedforward-Neural-Network/network"
)


// load only once
type canvasNN struct {
   loaded bool
   nn *network.NeuralNetwork
}
var canvasNet canvasNN = canvasNN{loaded:false}

// data from ajax
type canvasData struct {
   Width uint16 `json:"width"`
   Height uint16 `json:"height"`
   Data map[string]uint8 `json:"data"`
   Keys []int `json:"-"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
   if req.URL.Path == "/canvas-component/feedforward/" {
      inp_bytes,err:=ioutil.ReadAll(req.Body)
      if err!=nil{
         fmt.Printf("Error while reading request body: %v \n", err)
         io.WriteString(res, "Invalid request")
         return
      }

      var inp_data canvasData
      err = json.Unmarshal(inp_bytes, &inp_data)

      if err!=nil{
         fmt.Printf("Error when trying to unpack json: %v\n", err)
         io.WriteString(res, "Invalid request")
         return
      }
      for key,_:=range inp_data.Data {
         index, _ := strconv.Atoi(key)
         inp_data.Keys = append(inp_data.Keys, index)
      }
      sort.Ints(inp_data.Keys)

      /*
      feedforward_result:=feedforwardCanvas(inp_floats)
      */
      return
   }
   http.ServeFile(res, req, req.URL.Path[1:])
}

func feedforwardCanvas(input []float32) int {
   // load only once
   if !canvasNet.loaded{
      // json_file path relative to main.go
      json_file := "canvas-component/trained_network.json"
      canvasNet.nn = network.LoadNetworkFromFile( json_file )
      canvasNet.loaded = true
   }
   predictionVector := canvasNet.nn.Calculate(input)
   fmt.Printf("Prediction Vector:%v\n", predictionVector )

   predictionVectorNormalized := network.Predict(predictionVector)
   finalResult, _ := getOne(predictionVectorNormalized)

   fmt.Printf("Prediction Vector Normalized:%v\n", predictionVectorNormalized )
   fmt.Printf("Final Result:%v\n", finalResult )
   return finalResult
}

func getOne(input []float32) (int, string) {
   for i,j := range input {
      if j == 1 {
         return (i+1)%10, ""
      }
   }
   return -1, "Array doesn't contain num 1"
}
