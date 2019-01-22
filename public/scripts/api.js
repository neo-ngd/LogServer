var socketio = io(document.location.href + 'socket.io/')
var dispatch = {
    logger: null,
}
socketio.on('connect', function() {
    socketio.on('log:log', log => {
        if (dispatch.logger) {
            dispatch.logger(log);
        }
    });
});

socketio.on('disconnect', function () {
});

var SubscribeToLog = function(tag, cb) {
    socketio.emit('log:subscribe', tag);
    dispatch.logger = cb;
}