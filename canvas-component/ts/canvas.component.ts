import {Component} from 'angular2/core';
import {Http, Response, HTTP_PROVIDERS} from 'angular2/http';

@Component({
    selector: 'canvas-component',
    templateUrl: './canvas-component/canvas.partial.html',
    providers: [HTTP_PROVIDERS]
})

export class canvasComponent {
   canvasId: string = "canvasFeedforwardDemo";

   // defaults
   canvasWidth: number = 400;
   canvasHeight: number = 400;
   canvasStyle: string = 'border:1px solid #000000;';
   
   imageWidth: number = 28;
   imageHeight: number = 28;
   imageMinColor: number = 0;
   imageMaxColor: number = 255;
   imageDefaultColor: number = 225;
   // color value of each box (columns x rows)
   imageContent: number[][];

   mouseOver: boolean = false;
   mouseDown: boolean = false;

   // pixel (box) size
   dx: number = this.canvasWidth / this.imageWidth;
   dy: number = this.canvasHeight / this.imageHeight;

   // event listener
   mCanvas;
   mCanvasContext;

   // connects to backend
   http: Http;

   constructor(http: Http){
      this.http = http;
      this.imageContent = [];
      // initialize default color
      for(var i: number = 0; i < this.imageWidth; ++i){
         this.imageContent.push([]);
         for(var j: number = 0; j< this.imageHeight; ++j)
            this.imageContent[i].push(this.imageDefaultColor);
      }
   }

   ngAfterViewInit(){
      // after DOM has been loaded
      this.mCanvas = document.getElementById(this.canvasId);
      this.mCanvasContext = this.mCanvas.getContext('2d');
      this.resetCanvas();
   }

   updateMouseOverState(state: boolean){
      this.mouseOver = state;
      this.mouseDown = this.mouseOver ? this.mouseDown : false;
   }

   updateMouseDownState(state: boolean){
      this.mouseDown = state;
   }

   updateImageContent(event){
      // don't update if not required
      if(!this.mouseDown)
         return;
      // determine in which square mouse is located
      var cur_x: number = event.clientX/this.canvasWidth;
      cur_x = cur_x > 1.0 ? this.imageWidth - 1 : Math.round(cur_x*this.imageWidth) - 1;
      var cur_y: number = event.clientY/this.canvasHeight;
      cur_y = cur_y > 1.0 ? this.imageHeight - 1 : Math.round(cur_y*this.imageHeight) - 1;

      this.redrawCanvas(false, cur_x, cur_y);
   }

   resetCanvas(){
      this.redrawCanvas(true);
   }

   // if reset is false then update and redraw canvas
   // otherwise reset it to default colour
   redrawCanvas(reset: boolean, x_pos?: number, y_pos?: number){
      if(!reset && this.imageContent[x_pos][y_pos] > 0) {
         var cur: number = this.imageContent[x_pos][y_pos];
         cur = cur > 50 ? cur - 50 : 0;
         this.imageContent[x_pos][y_pos] = cur;
      }

      this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
      for(var i: number = 0; i < this.imageWidth; ++i){
         for(var j: number = 0; j < this.imageHeight; ++j){
            if(reset)
               this.imageContent[i][j] = this.imageDefaultColor;
            this.mCanvasContext.fillStyle = "rgb(" +
               this.imageContent[i][j] + ", " + // R
               this.imageContent[i][j] + ", " + // G
               this.imageContent[i][j] +        // B
               ")";
            this.mCanvasContext.fillRect(i*this.dx, j*this.dy, this.dx, this.dy);
         }
      }
   }

   feedforward(){
      // loop through rows x columns and fill in data
      var data: number[] = [];
      for(var i: number = 0; i < this.imageWidth; ++i)
         for(var j: number = 0; j < this.imageHeight; ++j)
               data.push(this.imageContent[j][i]);

      this.http.post("/canvas-component/feedforward/", JSON.stringify(data))
         .map((res: Response) => console.log(res))
         .subscribe(
            data => console.log("success")
         );
   }
}
