var socketio = io(document.location.href + 'socket.io/')

socketio.on('connect', function() {
    socketio.on('log:log', log => {
        subscriber.cb(log, null);
    });
});

socketio.on('disconnect', function () {
});

var SubscribeToLog = function(tag, cb) {
    socketio.emit('log:subscribe', tag);
}

var SubscribeTags = function(cb) {

}