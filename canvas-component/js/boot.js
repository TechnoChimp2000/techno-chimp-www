System.register(['angular2/platform/browser', './canvas.component', 'rxjs/add/operator/map'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var browser_1, canvas_component_1;
    return {
        setters:[
            function (browser_1_1) {
                browser_1 = browser_1_1;
            },
            function (canvas_component_1_1) {
                canvas_component_1 = canvas_component_1_1;
            },
            function (_1) {}],
        execute: function() {
            browser_1.bootstrap(canvas_component_1.canvasComponent);
        }
    }
});
