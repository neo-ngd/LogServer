var socketio = io(document.location.href + 'socket.io/')
var subscriber = {
    subscribed: false,
    cb: null,
}
socketio.on('connect', function() {
    socketio.on('log:log', log => {
        if (subscriber.subscribed) {
            subscriber.cb(log, null);
        }
    });
    if (subscriber.subscribed) {
        socketio.emit('log:subscribe', 'tag');
    }
});
socketio.on('disconnect', function () {
});

var SubscribeToLog = function(cb) {
    subscriber.cb = cb;
    subscriber.subscribed = true;
    if (socketio.connected) {
        socketio.emit('log:subscribe', 'tag');
    }
}