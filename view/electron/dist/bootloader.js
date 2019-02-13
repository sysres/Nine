const kernBinary = "kern.wasm";
var PIXEL_RATIO;
var kchanProc = new BroadcastChannel('kchan-proc');

PIXEL_RATIO = (function () {
  var ctx = document.createElement("canvas").getContext("2d"),
      dpr = window.devicePixelRatio || 1,
      bsr = ctx.webkitBackingStorePixelRatio ||
            ctx.mozBackingStorePixelRatio ||
            ctx.msBackingStorePixelRatio ||
            ctx.oBackingStorePixelRatio ||
            ctx.backingStorePixelRatio || 1;
  return dpr / bsr;
})();

window.addEventListener('load', boot);

// Create a canvas scaled up to the devicePixelRatio.
// Otherwise Nine will have to deal with the aspect ratio
// when drawing.
// https://www.html5rocks.com/en/tutorials/canvas/hidpi/
function createHiDPICanvas(w, h, ratio) {
    if (!ratio) { ratio = PIXEL_RATIO; }
    var can = document.createElement("canvas");
    can.width = w * ratio;
    can.height = h * ratio;
    can.style.width = w + "px";
    can.style.height = h + "px";
    can.getContext("2d").setTransform(ratio, 0, 0, ratio, 0, 0);
    return can;
}

function loadKern() {    
    let go = new Go();
    let mod, inst;
    
    WebAssembly.instantiateStreaming(fetch(kernBinary), go.importObject).then((result) => {
        mod = result.module;
        inst = result.instance;

        process.on("exit", (code) => { // Node.js exits if no event handler is pending
            if (code === 0 && !go.exited) {
                // deadlock, make Go print error and stack traces
                go._pendingEvent = { id: 0 };
                go._resume();
            }
        });    
      
        run();
    });

    async function run() {
      await go.run(inst);
    }

    kchanProc.onmessage = function(e) {
        alert('Message from process!!!'+e.data);
    }
}

function sched() {
    var worker = new Worker("sched.js");
    worker.onerror = function (evt) { alert(`Error from sched service: ${evt.message}`); };
    worker.onmessage = function (evt) {
        let msg = JSON.parse(evt.data);
        if(msg.type !== 'status') {
            alert(`CRITICAL: kernel received ${msg} from sched`);
            return;
        }

        if(msg.value !== 'running') {
            alert(`CRITICAL: sched service not started`);
            return;
        }

        // TODO(i4k): handle this from kernel
        worker.onmessage = function(evt) {
            console.log(`WARNING::js kernel:: Message from sched: ${evt.data}`);
        };

        worker.onerror = function(evt) {
            console.log(`WARNING::js kernel:: Error from sched: ${evt.message}`);
        }

        worker.postMessage(JSON.stringify({
            'type': 'exec',
            'path': 'wm/wm.wasm'
        }));
    };
}

function boot() {        
  loadKern();
  sched();
}