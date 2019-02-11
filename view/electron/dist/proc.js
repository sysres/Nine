importScripts('./goloader.js');


var service = null;
var go = null;
var mod = null;

fetch("proc.wasm").then(response =>
    response.arrayBuffer()
).then(bytes =>
    WebAssembly.compile(bytes)
).then(WasmModule =>
    mod = WasmModule
);

// Listen for messages from the main thread. Because all messages to this thread come through
// this method, we need a way to know what is being asked of us which is why we included the
// MessagePurpose property.
self.onmessage = function (evt) {
    var objData = evt.data;
    var msgType = objData.msgType;
    if (msgType === "Hello") {
        Hello();
    } else if (msgType === "load") {
        go = new Go();
        go.argv = ["dist/proc"];
        go.env = process.env;
        //go.exit = process.exit;

        WebAssembly.instantiate(mod, go.importObject).then(result => {
            service = result; 
            process.on("exit", (code) => { // Node.js exits if no event handler is pending
                if (code === 0 && !go.exited) {
                    // deadlock, make Go print error and stack traces
                    go._pendingEvent = { id: 0 };
                    go._resume();
                }
            }); 
            run()
        });
    }
}

async function run() {
    await go.run(service);
}