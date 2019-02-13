importScripts('./goloader.js');

let go;
let path; // program path
let mod;
let inst;
var kchan = new BroadcastChannel('kchan-proc');

self.onmessage = function(evt) {
    let msg = JSON.parse(evt.data);

    switch(msg.type) {
    case 'load':
        load(msg.path, function() {
            start();
            postMessage(JSON.stringify({
                'type': 'status',
                'value': 'running'
            }));
        }, function(error) {
            postMessage(JSON.stringify({
                'type': 'status',
                'value': 'error',
                'message': error.message
            }));
        });
        break;
    case 'start':
        start();
        postMessage(JSON.stringify({
            'type': 'status',
            'value': 'running'
        }));
        break;
    default:
        console.log(`proc received unknown: ${msg} from sched`);
    }
}

function load(path, successcb, errorcb) {
    self.path = path;
    fetch(path).then(response => 
        response.arrayBuffer()
    ).then(bytes =>
        WebAssembly.compile(bytes)
    ).then(WasmModule => {
        self.mod = WasmModule
        return successcb()
    }).catch(errorcb);
}

function start() {
    go = new Go();
    go.argv = [self.path];
    go.env = process.env;
    //go.exit = process.exit;

    WebAssembly.instantiate(self.mod, go.importObject).then(result => {
        self.inst = result; 
        process.on("exit", (code) => { // Node.js exits if no event handler is pending
            if (code === 0 && !go.exited) {
                // deadlock, make Go print error and stack traces
                go._pendingEvent = { id: 0 };
                go._resume();
            }

            //TODO(i4k): send an event termination event to proc to terminate the worker
            close();
        }); 
        run()
    });
}

async function run() {
    await go.run(self.inst);
}

// synchronize with kernel
postMessage(JSON.stringify({
    'type': 'status',
    'value': 'running'
}));