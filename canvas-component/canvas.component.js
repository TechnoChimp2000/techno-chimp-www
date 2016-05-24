System.register(['angular2/core', 'angular2/http'], function(exports_1, context_1) {
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
    var core_1, http_1;
    var canvasComponent;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (http_1_1) {
                http_1 = http_1_1;
            }],
        execute: function() {
            canvasComponent = (function () {
                function canvasComponent(http) {
                    this.canvasId = "canvasFeedforwardDemo";
                    this.secondCanvasId = "secondCanvas";
                    this.canvasWidth = 300;
                    this.canvasHeight = 300;
                    this.canvasStyle = 'border:1px solid #000000; float: left;';
                    this.secondCanvasStyle = 'border:1px solid #000000; float: left;';
                    this.mouseOver = false;
                    this.mouseDown = false;
                    this.min_x = this.canvasWidth;
                    this.max_x = 0;
                    this.min_y = this.canvasHeight;
                    this.max_y = 0;
                    this.showResult = false;
                    this.http = http;
                }
                canvasComponent.prototype.ngAfterViewInit = function () {
                    this.mCanvas = document.getElementById(this.canvasId);
                    this.mCanvasContext = this.mCanvas.getContext('2d');
                    this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                    this.mSecondCanvas = document.getElementById(this.secondCanvasId);
                    this.mSecondCanvasContext = this.mSecondCanvas.getContext('2d');
                    this.mSecondCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                    var pic = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.20462185, 0.5, 0.5, 0.122709095, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.4187675, 1, 1, 0.15318125, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.09539819, 0.7103641, 0.5766107, 0.0037415028, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.7244898, 0.78571427, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0.78571427, 0.7244898, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0.71428573, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.21482593, 0.78571427, 0.78571427, 0.2944578, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.2310524, 1, 1, 0.23217285, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.00048017502, 0.37056822, 0.36998802, 0.00042015314, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
                    this.updateSecondCanvas(pic);
                };
                canvasComponent.prototype.updateMouseOverState = function (state) {
                    this.mouseOver = state;
                    this.mouseDown = this.mouseOver ? this.mouseDown : false;
                    if (this.mouseOver)
                        this.mCanvas.style.cursor = "pointer";
                    else
                        this.mCanvas.style.cursor = "auto";
                };
                canvasComponent.prototype.updateMouseDownState = function (event, state) {
                    this.mouseDown = state;
                    this.updateImageContent(event);
                };
                canvasComponent.prototype.updateImageContent = function (event) {
                    if (!this.mouseDown)
                        return;
                    var posx = event.clientX;
                    var posy = event.clientY;
                    this.mCanvasContext.beginPath();
                    for (var i = 0; i < 5; ++i) {
                        var alpha = (5.0 - i) / 5.0;
                        this.mCanvasContext.fillStyle = 'rgba(0, 0, 0, ' + alpha + ')';
                        var radius = 6.0 / 100.0 * this.canvasWidth;
                        var left = posx - radius;
                        var right = posx + radius;
                        var up = posy - radius;
                        var down = posy + radius;
                        if (left < this.min_x)
                            this.min_x = left < 0 ? 0 : left;
                        if (right > this.max_x)
                            this.max_x = right > this.canvasWidth ? this.canvasWidth : right;
                        if (up < this.min_y)
                            this.min_y = up < 0 ? 0 : up;
                        if (down > this.max_y)
                            this.max_y = down > this.canvasHeight ? this.canvasHeight : down;
                        this.mCanvasContext.arc(posx, posy, radius, 0, Math.PI * 2);
                        this.mCanvasContext.fill();
                    }
                };
                canvasComponent.prototype.updateSecondCanvas = function (pic28x28) {
                    var dx = this.canvasWidth / 28.0;
                    var dy = this.canvasHeight / 28.0;
                    for (var i = 0; i < 28; i++) {
                        for (var j = 0; j < 28; j++) {
                            var color = 255 - Math.floor(pic28x28[28 * i + j] * 255.0);
                            this.mSecondCanvasContext.fillStyle = 'rgb(' + color + ', ' + color + ', ' + color + ')';
                            this.mSecondCanvasContext.fillRect(j * dx, i * dy, dx, dy);
                        }
                    }
                };
                canvasComponent.prototype.resetCanvas = function () {
                    this.showResult = false;
                    this.min_x = this.canvasWidth;
                    this.max_x = 0;
                    this.min_y = this.canvasHeight;
                    this.max_y = 0;
                    this.mCanvasContext.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                };
                canvasComponent.prototype.feedforward = function () {
                    var _this = this;
                    var width = this.max_x - this.min_x;
                    var height = this.max_y - this.min_y;
                    var canvas_data = this.mCanvasContext.getImageData(this.min_x, this.min_y, width, height);
                    var json_data = JSON.stringify({
                        'width': canvas_data.width,
                        'height': canvas_data.height,
                        'data': canvas_data.data
                    });
                    this.http.post("/canvas-component/feedforward/", json_data)
                        .map(function (res) {
                        _this.result = res.json();
                        _this.showResult = true;
                    })
                        .subscribe(function (data) { return console.log("success"); });
                };
                canvasComponent = __decorate([
                    core_1.Component({
                        selector: 'canvas-component',
                        templateUrl: 'canvas-component/canvas.partial.html',
                        providers: [http_1.HTTP_PROVIDERS]
                    }), 
                    __metadata('design:paramtypes', [(typeof (_a = typeof http_1.Http !== 'undefined' && http_1.Http) === 'function' && _a) || Object])
                ], canvasComponent);
                return canvasComponent;
                var _a;
            }());
            exports_1("canvasComponent", canvasComponent);
        }
    }
});
