const API_ENDPOINTS = "http://127.0.0.1:8000";

export async function loginHTTPRequest(username, password) {
    const response = await fetch(`${API_ENDPOINTS}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username,
            password
        })
    });
    return await response.json();
}
export async function registerHTTPRequest(username, password) {
    const response = await fetch(`${API_ENDPOINTS}/registration`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username,
            password
        })
    });
    return await response.json();
}

export async function isUsernameAvailableHTTPRequest(username) {
    const response = await fetch(`${API_ENDPOINTS}/isUsernameAvailable/${username}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return await response.json();
}

export async function userSessionCheckHTTPRequest(username) {
    const response = await fetch(`${API_ENDPOINTS}/userSessionCheck/${username}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return await response.json();
}

export async function userCreateRoom(createRoomMessage) {
    try {
        const response = await fetch(`${API_ENDPOINTS}/CreateRoom`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(createRoomMessage)
        });
        return await response.json();
    } catch(err) {
    }

}

export async function userJoinRoom(joinRoomMessage) {
    try {
        const response = await fetch(`${API_ENDPOINTS}/JoinRoom`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(joinRoomMessage)
        });
        return await response.json();
    } catch(err) {
    }

}

export async function getConversationBetweenUsers(toUserID, fromUserID) {
    const response = await fetch(`${API_ENDPOINTS}/getConversation/${toUserID}/${fromUserID}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return await response.json();
}

export async function getConversationBetweenGroups(groupID, fromUserID) {
    const response = await fetch(`${API_ENDPOINTS}/getRoomChatConversation/${groupID}/${fromUserID}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return await response.json();
}