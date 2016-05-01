package canvasComp

import (
   "fmt"
   "io"
   "io/ioutil"
   "strconv"
   "net/http"
   "encoding/json"
   "github.com/Go-Feedforward-Neural-Network/network"
)

func Handler(res http.ResponseWriter, req *http.Request) {
   if req.URL.Path == "/canvas-component/feedforward/" {
      inp_bytes,err:=ioutil.ReadAll(req.Body)
      if err!=nil{
         fmt.Printf("Error while reading request body: %v \n", err)
         io.WriteString(res, "Invalid request")
         return
      }
      inp_floats := make([]float32, 784)
      err = json.Unmarshal(inp_bytes, &inp_floats)
      if err!=nil{
         fmt.Printf("Error while trying to unpack json: %v \n", err)
         io.WriteString(res, "Invalid request")
         return
      }
      feedforward_result:=feedforwardCanvas(inp_floats)
      io.WriteString(res, strconv.Itoa(feedforward_result))
      return
   }
   http.ServeFile(res, req, req.URL.Path[1:])
}

func feedforwardCanvas(input []float32) int {
   // json_file path relative to main.go
   json_file := "canvas-component/trained_network.json"
   nn := network.LoadNetworkFromFile( json_file )

   predictionVector := nn.Calculate(input)
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
