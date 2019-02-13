let processes = [];
let nextid = 0;

function exec(path) {
    let worker = new Worker("proc.js")
    worker.onerror = function(error) {
        console.log('proc error: ' + error.message);
    };

    worker.onmessage = function(evt) {
        let msg = JSON.parse(evt.data);
        if (msg.type !== 'status') {
            alert('CRITICAL: proc sent '+evt.data);
            return;
        }

        if (msg.value !== 'running') {
            alert('CRITICAL: proc failed to load');
            return;
        }

        worker.onmessage = onProcMessage;
        worker.postMessage(JSON.stringify({
            'type': 'load',
            'path': path
        }));
    };

    processes.push({
        'pid': ++nextid,
        'worker': worker,
    });
}

function onProcMessage(evt) {
    let msg = JSON.parse(evt.data);
    postMessage(JSON.stringify(msg));
    console.log(`sched received ${evt.data} from proc`)
}

// events: 
//   - start process
//   - terminate process
//   - send note to process
//   - etc
self.onmessage = function (evt) {
    let msg = JSON.parse(evt.data);

    switch(msg.type) {
    case 'exec':
        exec(msg.path);
        break;
    default:
        console.log(`sched received from kernel: `+  evt.data);
    }    
}

async function run() {
    await go.run(service);
}

// synchronize with kernel
postMessage(JSON.stringify({
    'type': 'status',
    'value': 'running'
}));
