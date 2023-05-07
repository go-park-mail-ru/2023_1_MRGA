const path = require('path');
const ws = require('ws');
const ws_connection = require('./ws');

const public = path.join(path.resolve(__dirname, '..'), 'dist');

const config = {
    PORT: 3000,
    WS_PORT: 9090,
    HOST: "0.0.0.0"
}

const wsServer = new ws.Server({host: config.HOST, port: config.WS_PORT}, () => {
    console.log(`WebSocketServer is running on ${config.HOST}:${config.WS_PORT}`);
});
wsServer.on('connection', ws_connection);
