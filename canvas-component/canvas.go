package canvas

import (
   "fmt"
   "io"
   "io/ioutil"
   "strconv"
   "net/http"
   "encoding/json"
   "github.com/Go-Feedforward-Neural-Network/network"
)


func CanvasHandler(res http.ResponseWriter, req *http.Request) {
   if(req.URL.Path == "/canvas-component/feedforward/") {
      inp_bytes,_:=ioutil.ReadAll(req.Body)
      inp_floats := make([]float32, 784)
      json.Unmarshal(inp_bytes, &inp_floats)
      out:=feedforwardCanvas(inp_floats)
      io.WriteString(res, strconv.Itoa(out))
      return
   }

   http.ServeFile(res, req, req.URL.Path[1:])
}

func feedforwardCanvas(input []float32) int {
   // json_file path relative to server.go
   json_file := "./canvas-component/files/trained_network.json"
   nn := network.LoadNetworkFromFile( json_file )

   // produce a prediction on a sample
   predictionVector := nn.Calculate(input)
   fmt.Printf("Prediction Vector:%v\n", predictionVector )

   predictionVectorNormalized := network.Predict(predictionVector)
   finalResult := GetInteger(predictionVectorNormalized)

   fmt.Printf("Prediction Vector Normalized:%v\n", predictionVectorNormalized )
   fmt.Printf("Final Result:%v\n", finalResult )
   return finalResult
}

func GetInteger(input []float32) int {
   for i,j := range input {
      if j == 1 {
         return i+1
      }
   }
   return -1
}
