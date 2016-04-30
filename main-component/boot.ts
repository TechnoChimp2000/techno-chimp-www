import {bootstrap} from 'angular2/platform/browser';
import {canvasComponent} from 'canvas-component/canvas.component.js'; // SystemJS cannot handle without .js extension
import 'rxjs/add/operator/map';

bootstrap(canvasComponent);
