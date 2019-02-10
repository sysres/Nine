        
var service = null;
var go = null;

// Listen for messages from the main thread. Because all messages to this thread come through
// this method, we need a way to know what is being asked of us which is why we included the
// MessagePurpose property.
self.onmessage = function (evt) {
    var objData = evt.data;
    var msgType = objData.msgType;
    if (msgType === "Hello") {
        Hello();
    } else if (msgType === "CompiledModule") {
        const go = new Go();
        go.argv = ["dist/"+wasmName];
        go.env = process.env;
        //go.exit = process.exit;

        WebAssembly.instantiate(objData.WasmModule, go.importObject).then(result => {
            service = result; // Hold onto the module's instance so that we can reuse it
            return go.run(result.instance)
        });
    }
}