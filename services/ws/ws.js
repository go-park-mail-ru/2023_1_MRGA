let userIdClients = new Map();
let id = 1;

module.exports = function onConnect(wsClient) {
    wsClient.id = id;
    id++;

    wsClient.send(JSON.stringify({
        flag: "CONN+ASK",
        status: 200,
    }));
    console.log(`CONN CLIENT ${wsClient.id}`);
    
    wsClient.on('message', function(message) {

        const jsonMSG = JSON.parse(message);

        switch (jsonMSG.flag) {
        case "REG":
            const userId = jsonMSG.body.userId;
            wsClient.userId = userId;

            let clientsByUser = userIdClients.get(userId);

            if (clientsByUser === undefined || clientsByUser.length === 0) {
                userIdClients.set(userId, [wsClient]);
            } else if (clientsByUser.length > 0) {
                clientsByUser.push(wsClient);
                userIdClients.set(userId, clientsByUser);
            } else {
                wsClient.send(JSON.stringify({
                    flag: "REG+ASK",
                    status: 404,
                    err: `Не удается зарегистрировать пользователя ${userId} на сервере.`,
                }));
                return;
            }
            
            wsClient.send(JSON.stringify({
                flag: "REG+ASK",
                status: 200,
            }));

            console.log(`REG ${wsClient.userId} ${wsClient.id}`);
            break;
        case "SEND":
            const msg = jsonMSG.body.msg;
            const userIds = jsonMSG.body.userIds;
            const chatId = jsonMSG.body.chatId;
            const sentAt = jsonMSG.body.sentAt;

            let sentAll = true;
            userIds.forEach(userId => {
                const clientsByUser = userIdClients.get(userId);

                const success = sendToClients(wsClient, clientsByUser, chatId, msg, sentAt);
                if (!success && sentAll) {
                    sentAll = false;
                }

            });
            if (sentAll) {
                const successSelf = sendToClients(wsClient, userIdClients.get(wsClient.userId), chatId, msg, sentAt);
                if (!successSelf) {
                    wsClient.send(JSON.stringify({
                        flag: "SEND+ASK",
                        status: 500,
                        err: "Не удалось отправить пользователю его сообщение",
                    }));
                } else {
                    wsClient.send(JSON.stringify({
                        flag: "SEND+ASK",
                        status: 200,
                    }));
                }
            }
            break;
        }
    });

    wsClient.on('close', () => {
        const userId = wsClient.userId;
        const id = wsClient.id;
        let clientsByUser = userIdClients.get(userId);
        
        if (clientsByUser !== undefined && clientsByUser.length > 0) {
            userIdClients.set(userId, clientsByUser.filter(client => client.id !== id));
        }
        console.log(`CLOSE ${userId} (${id})`);
    });
};

function sendToClients(wsClient, clients, chatId, msg, sentAt) {
    if (clients !== undefined && clients.length > 0) {
        clients.forEach(receiverWsClient => {
            if (receiverWsClient !== undefined) {
                const msgData = JSON.stringify({
                    flag: "SEND",
                    body: {
                        chatId,
                        senderId: wsClient.userId,
                        msg,
                        sentAt,
                    },
                });

                receiverWsClient.send(msgData);

                console.log(`SEND ${receiverWsClient.userId} (${receiverWsClient.id}) MSG ${msgData}`);
            }
        });
    } else {
        wsClient.send(JSON.stringify({
            flag: "SEND+ASK",
            status: 200,
        }));
        return true;
    }
    return true;
}
