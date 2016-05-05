import {Component} from 'angular2/core';
import {Http, Response, HTTP_PROVIDERS} from 'angular2/http';

@Component({
    selector: 'canvas-component',
    // templateUrl path relative to main.go
    templateUrl: 'canvas-component/canvas.partial.html',
    providers: [HTTP_PROVIDERS]
})

export class canvasComponent {
   canvasId: string = "canvasFeedforwardDemo";

   // defaults
   canvasWidth: number = 300;
   canvasHeight: number = 300;
   canvasStyle: string = 'border:1px solid #000000;';

   mouseOver: boolean = false;
   mouseDown: boolean = false;

   // event listener
   mCanvas;
   mCanvasContext;

   // connects to backend
   http: Http;

   // crop image by this before sending to backend
   min_x: number = this.canvasWidth;
   max_x: number = 0;
   min_y: number = this.canvasHeight;
   max_y: number = 0;


   showResult: boolean = false;
   result: number;

   constructor(http: Http){
      this.http = http;
   }

   ngAfterViewInit(){
      // after DOM has been loaded
      this.mCanvas = document.getElementById(this.canvasId);
      this.mCanvasContext = this.mCanvas.getContext('2d');
      this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
   }

   updateMouseOverState(state: boolean){
      this.mouseOver = state;
      this.mouseDown = this.mouseOver ? this.mouseDown : false;
      if(this.mouseOver)
         this.mCanvas.style.cursor = "pointer";
      else
         this.mCanvas.style.cursor = "auto";
   }

   updateMouseDownState(event, state: boolean){
      this.mouseDown = state;
      this.updateImageContent(event);
   }

   updateImageContent(event){
      // don't update if not required
      if(!this.mouseDown)
         return;

      var posx = event.clientX;
      var posy = event.clientY;

      this.mCanvasContext.beginPath();
      for(var i: number=0; i<5; ++i) {
         var alpha:number = (5.0-i)/5.0;
         this.mCanvasContext.fillStyle = 'rgba(0, 0, 0, '+alpha+')';
         var radius:number = (1.0 + Math.log(1.0 + this.canvasWidth/100.0*i))/100.0*this.canvasWidth;
         var left = posx - radius;
         var right = posx + radius;
         var up = posy - radius;
         var down = posy + radius;
         if(left < this.min_x)
            this.min_x = left < 0 ? 0 : left;
         if(right > this.max_x)
            this.max_x = right > this.canvasWidth ? this.canvasWidth : right;
         if(up < this.min_y)
            this.min_y = up < 0 ? 0 : up;
         if(down > this.max_y)
            this.max_y = down > this.canvasHeight ? this.canvasHeight : down;

         this.mCanvasContext.arc(posx, posy, radius,0, Math.PI*2);
         this.mCanvasContext.fill();
      }
   }

   resetCanvas(){
      this.showResult = false;
      this.min_x  = this.canvasWidth;
      this.max_x = 0;
      this.min_y = this.canvasHeight;
      this.max_y = 0;
      this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
   }

   feedforward(){
      var width: number = this.max_x - this.min_x;
      var height: number = this.max_y - this.min_y;
      var canvas_data = this.mCanvasContext.getImageData(this.min_x, this.min_y, width, height);
      var json_data = JSON.stringify({
         'width' : canvas_data.width,
         'height' : canvas_data.height,
         'data' : canvas_data.data
      });
      this.http.post("/canvas-component/feedforward/", json_data)
         .map((res: Response) => {
            //this.result = res.json();
            this.showResult = true;
         })
         .subscribe(
            data => console.log("success")
         );
   }
}
