var ws = new WebSocket(config.server)
var subscriber = null
console.log('api init')
ws.onopen = function() {
    socket.on('log:log', log => cb(log, null));
    socket.emit('log:subscribe', "log subscriber");
}

ws.onmessage = function(e) {
    if (subscriber) {
        subscriber(e.data);
    }
}

subscribeToLog = function(cb) {
    console.log('subscriber')
    subscriber = cb;
}