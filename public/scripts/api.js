var socketio = io.connect(document.location.href + 'socket.io/')
var subscriber = function(){}
socketio.on('connect', function() {
    socketio.on('log:log', log => subscriber(log, null));
    setTimeout(() => {
        socketio.emit('log:subscribe', 'tag');
    }, 500);
});
socketio.on('disconnect', function () {
});

var SubscribeToLog = function(cb) {
    subscriber = cb;
}