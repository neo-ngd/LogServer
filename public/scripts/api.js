var socketio = io.connect(document.location.href + 'socket.io/')
var subscriber = null
socketio.on('connect', function() {
    socketio.on('log:log', log => subscriber(log, null));
    socketio.send('log:subscribe', "subscribe tag");
});
socketio.on('disconnect', function () {
});

var SubscribeToLog = function(cb) {
    subscriber = cb;
}