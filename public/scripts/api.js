var wsServer = document.location.href.replace("http", "ws") + "ws/";
var ws = new WebSocket(wsServer);
var logger = null
var message = ""
ws.onopen = function (e) {
    console.log("Connected to WebSocket server.",e);
    if (logger) {
        ws.send(message)
    }
} ;

ws.onclose = function (e) {
    console.log("Disconnected",e);
} ;

ws.onmessage = function(e) {
    let logbody = JSON.parse(e.data)
    if (logger) {
        logger(logbody);
    }
}

ws.onerror = function (e) {
    console.log('Error occured: ' + e.data,e);
} ;

function Regist(name, cb) {
    console.log("regist")
    logger = cb;
    console.log(ws.OPEN, ws.CONNECTING, ws.CLOSED, ws.CLOSING)
    console.log(ws.readyState)
    if (ws.readyState === ws.OPEN) {
        console.log("send")
        ws.send(name);
        return;
    }
    console.log("store")
    message = name;
}

function Close() {
    ws.close()
}