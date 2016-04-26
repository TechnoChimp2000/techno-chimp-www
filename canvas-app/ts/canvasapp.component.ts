import {Component} from 'angular2/core';

@Component({
    selector: 'canvas-app',
    templateUrl: 'canvasapp.partial.html'
})

export class canvasApp {
   // defaults
   canvasWidth: number = 400;
   canvasHeight: number = 400;
   canvasStyle: string = 'border:1px solid #000000;';

   imageWidth: number = 28;
   imageHeight: number = 28;
   imageMinColor: number = 0;
   imageMaxColor: number = 255;
   imageDefaultColor: number = 225;
   imageContent: number[][];

   mouseOver: boolean = false;
   mouseDown: boolean = false;

   // pixel (box) size
   pixelDx: number = this.canvasWidth / this.imageWidth;
   pixelDy: number = this.canvasHeight / this.imageHeight;

   // event listener
   mCanvas;
   mCanvasContext;

   constructor(){
      // initialize to default color
      this.imageContent = [];
      for(var i: number = 0; i < this.imageWidth; ++i){
         this.imageContent.push([]);
         for(var j: number = 0; j< this.imageHeight; ++j)
            this.imageContent[i].push(this.imageDefaultColor);
      }
   }

   ngAfterViewInit(){
      // after DOM has been loaded
      this.mCanvas = document.getElementById('canvas-app-id');
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
      cur_x = cur_x > 1.0 ? this.imageWidth: Math.round(cur_x*this.imageWidth) - 1;
      var cur_y: number = event.clientY/this.canvasHeight;
      cur_y = cur_y > 1.0 ? this.imageHeight: Math.round(cur_y*this.imageHeight) - 1;

      this.redrawCanvas(false, cur_x, cur_y);
   }

   resetCanvas(){
      this.redrawCanvas(true, 0, 0);
   }

   // if reset is false then update and redraw canvas
   // otherwise reset it to default colour
   redrawCanvas(reset: boolean, x_pos: number, y_pos: number){
      if(!reset && this.imageContent[x_pos][y_pos] > 0) {
            this.imageContent[x_pos][y_pos] -= 51;
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
            this.mCanvasContext.fillRect(i*this.pixelDx, j*this.pixelDy, this.pixelDx, this.pixelDy);
         }
      }
   }
}
