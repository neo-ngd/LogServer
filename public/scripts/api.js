var socketio = io.connect(Config.Server)
var subscriber = null
console.log('api init')
socketio.on('connect', function() {
    socketio.on('log:log', log => subscriber(log, null));
    socketio.send('log:subscribe', "log subscriber");
});
socketio.on('disconnect', function () {
});

var SubscribeToLog = function(cb) {
    subscriber = cb;
}