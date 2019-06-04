var socketio = null;
var dispatch = {
    logger: null,
}
var SubscribeToLog = function(tag, cb) {
    if (socketio === null) {
        socketio = io(document.location.href + 'socket.io/');
        socketio.on('connect', function() {
    
        });
        socketio.on('log:log', log => {
            if (dispatch.logger) {
                dispatch.logger(log);
            }
        });
        socketio.on('disconnect', function () {
            socketio.close();
            socketio = null;
        });
        socketio.emit('log:subscribe', tag);
        dispatch.logger = cb;
    } else {
        socketio.emit('log:subscribe', tag);
        dispatch.logger = cb;
    }
    
}
var SocketClose = function() {
    socketio.close();
    socketio = null;
}