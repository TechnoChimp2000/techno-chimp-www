System.register(['angular2/core'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
        var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
        if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
        else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
        return c > 3 && r && Object.defineProperty(target, key, r), r;
    };
    var __metadata = (this && this.__metadata) || function (k, v) {
        if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
    };
    var core_1;
    var canvasApp;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            }],
        execute: function() {
            canvasApp = (function () {
                function canvasApp() {
                    this.canvasWidth = 400;
                    this.canvasHeight = 400;
                    this.canvasStyle = 'border:1px solid #000000;';
                    this.imageWidth = 28;
                    this.imageHeight = 28;
                    this.imageMinColor = 0;
                    this.imageMaxColor = 255;
                    this.imageDefaultColor = 225;
                    this.mouseOver = false;
                    this.mouseDown = false;
                    this.pixelDx = this.canvasWidth / this.imageWidth;
                    this.pixelDy = this.canvasHeight / this.imageHeight;
                    this.imageContent = [];
                    for (var i = 0; i < this.imageWidth; ++i) {
                        this.imageContent.push([]);
                        for (var j = 0; j < this.imageHeight; ++j)
                            this.imageContent[i].push(this.imageDefaultColor);
                    }
                }
                canvasApp.prototype.ngAfterViewInit = function () {
                    this.mCanvas = document.getElementById('canvas-app-id');
                    this.mCanvasContext = this.mCanvas.getContext('2d');
                    this.resetCanvas();
                };
                canvasApp.prototype.updateMouseOverState = function (state) {
                    this.mouseOver = state;
                    this.mouseDown = this.mouseOver ? this.mouseDown : false;
                };
                canvasApp.prototype.updateMouseDownState = function (state) {
                    this.mouseDown = state;
                };
                canvasApp.prototype.updateImageContent = function (event) {
                    if (!this.mouseDown)
                        return;
                    var cur_x = event.clientX / this.canvasWidth;
                    cur_x = cur_x > 1.0 ? this.imageWidth : Math.round(cur_x * this.imageWidth) - 1;
                    var cur_y = event.clientY / this.canvasHeight;
                    cur_y = cur_y > 1.0 ? this.imageHeight : Math.round(cur_y * this.imageHeight) - 1;
                    this.redrawCanvas(false, cur_x, cur_y);
                };
                canvasApp.prototype.resetCanvas = function () {
                    this.redrawCanvas(true, 0, 0);
                };
                canvasApp.prototype.redrawCanvas = function (reset, x_pos, y_pos) {
                    if (!reset && this.imageContent[x_pos][y_pos] > 0) {
                        this.imageContent[x_pos][y_pos] -= 51;
                    }
                    this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                    for (var i = 0; i < this.imageWidth; ++i) {
                        for (var j = 0; j < this.imageHeight; ++j) {
                            if (reset)
                                this.imageContent[i][j] = this.imageDefaultColor;
                            this.mCanvasContext.fillStyle = "rgb(" +
                                this.imageContent[i][j] + ", " +
                                this.imageContent[i][j] + ", " +
                                this.imageContent[i][j] +
                                ")";
                            this.mCanvasContext.fillRect(i * this.pixelDx, j * this.pixelDy, this.pixelDx, this.pixelDy);
                        }
                    }
                };
                canvasApp = __decorate([
                    core_1.Component({
                        selector: 'canvas-app',
                        templateUrl: 'canvasapp.partial.html'
                    }), 
                    __metadata('design:paramtypes', [])
                ], canvasApp);
                return canvasApp;
            }());
            exports_1("canvasApp", canvasApp);
        }
    }
});
