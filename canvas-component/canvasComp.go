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

// MNIST images are 28x28 but contentSize (from cropped canvas) is smaller
// to account for some whitespace at the edges
var contentSize uint16 = 20;
var edgeSize uint16 = 28 - contentSize;

// data from ajax
type canvasData struct {
   Width uint16 `json:"width"`
   Height uint16 `json:"height"`
   Data map[string]uint8 `json:"data"`
   Keys []int `json:"-"` // calculated
}

func Handler(res http.ResponseWriter, req *http.Request) {
   if req.URL.Path == "/canvas-component/feedforward/" {
      // read input data
      inp_bytes,err:=ioutil.ReadAll(req.Body)
      if err!=nil{
         fmt.Printf("Error while reading request body: %v \n", err)
         io.WriteString(res, "Invalid request")
         return
      }
      // unpack input data into json
      var inp_data canvasData
      err = json.Unmarshal(inp_bytes, &inp_data)
      if err!=nil{
         fmt.Printf("Error when trying to unpack json: %v\n", err)
         io.WriteString(res, "Invalid request")
         return
      }
      // sort input pixels in ascending order
      for key,_:=range inp_data.Data {
         index, _ := strconv.Atoi(key)
         inp_data.Keys = append(inp_data.Keys, index)
      }
      sort.Ints(inp_data.Keys)

      inp_floats:=transformCanvasData(&inp_data)

      fmt.Printf("[")//debug
      for i:=0; i<len(inp_floats); i++{//debug
         fmt.Printf("%v, ", inp_floats[i])//debug
      }//debug
      fmt.Printf("]\n")//debug
      //fmt.Printf("%v\n",inp_floats) //debug

      feedforward_result:=feedforwardCanvas(inp_floats)
      io.WriteString(res, strconv.Itoa(feedforward_result))

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

// takes averages across canvasData and
// returns slice of 784 floats which represent 28x28 image
func transformCanvasData(cd *canvasData) []float32 {
   retval:=make([]float32, 784)
   out:=make([]uint8, len(cd.Data)/4)
   // take only alpha values (by canvas construction)
   for i:=3; i<len(cd.Data); i+=4 {
      val:=cd.Data[strconv.Itoa(cd.Keys[i])]
      out[(i+1)/4 - 1] = 255 - val
   }

   width_remainder := cd.Width%contentSize
   height_remainder := cd.Height%contentSize
   var cell_width uint16 = cd.Width/contentSize
   var additional_width uint16 = 0
   if width_remainder > 0 {
      for {
         additional_width++
         if (cd.Width + additional_width) % contentSize == 0 {
            break
         }
      }
      cell_width = (cd.Width + additional_width)/contentSize
   }

   var additional_height uint16 = 0
   var cell_height uint16 = cd.Height/contentSize;
   if height_remainder > 0 {
      for {
         additional_height++
         if (cd.Height + additional_height) % contentSize == 0 {
            break
         }
      }
      cell_height = (cd.Height +additional_height)/contentSize
   }

   // check for additional height or width
   if cell_width > cell_height {
      additional_height += (cd.Width + additional_width) - (cd.Height + additional_height)
   } else if cell_height > cell_width {
      additional_width += (cd.Height + additional_height) - (cd.Width + additional_width)
   }

   // first add additional columns
   enlarged_cols:=make([]uint8, len(out) + int(cd.Height*additional_width))
   if additional_width > 0 {
      var left uint16 = additional_width/2;
      var right uint16 = additional_width - left;
      for i:=uint16(0); i < cd.Height; i++ {
         for j:=uint16(0); j<left; j++ {
            enlarged_cols[i*(cd.Width + left + right)+j] = 255
         }
         for j:=uint16(0); j<cd.Width; j++{
            enlarged_cols[i*(cd.Width+left+right) + left +j] = out[i*cd.Width+j]
         }
         for j:=uint16(0); j<right; j++ {
            enlarged_cols[i*(cd.Width+left+right) + left + cd.Width +j] = 255
         }
      }
   }

   // then add additional rows
   enlarged_cols_and_rows:=make([]uint8, len(enlarged_cols) + int((cd.Width+additional_width)*additional_height))
   if additional_height > 0 {
      var top uint16 = additional_height/2;
      var bottom uint16 = additional_height - top;
      row:=uint16(cd.Width + additional_width)
      for i:=uint16(0); i<top; i++ {
         for j:=uint16(0); j<row; j++ {
            enlarged_cols_and_rows[i*row + j] = 255
         }
      }
      b1:=top*row
      for i:=uint16(0); i < uint16(len(enlarged_cols)); i++ {
         enlarged_cols_and_rows[b1 + i] = enlarged_cols[i]
      }
      b2:=top*row + uint16(len(enlarged_cols))
      for i:=uint16(0); i<bottom; i++ {
         for j:=uint16(0); j<row; j++ {
            enlarged_cols_and_rows[b2 + i*row + j] = 255
         }
      }
   }

   // pick bigger
   var cell_size uint16 = cell_height
   if cell_width > cell_height{
      cell_size = cell_width
   }

   // average out each cell and return
   row_size := cd.Width+additional_width
   for i:=uint16(0); i<contentSize; i++ {
      for j:=uint16(0); j<contentSize; j++ {
         cur_top_left := i*row_size*cell_size + j*cell_size
         var cell_sum uint32 = 0
         for k:=uint16(0); k<cell_size; k++ {
            for l:=uint16(0); l<cell_size;l++ {
               cell_sum += uint32(enlarged_cols_and_rows[cur_top_left + k*row_size + l])
            }
         }
         var avg float64 = float64(cell_sum)/float64(cell_size*cell_size)/255.0
         retval[(i+uint16(edgeSize/2))*28 + j + uint16(edgeSize/2)] = 1.0 - float32(avg)
      }
   }
   return retval
}
