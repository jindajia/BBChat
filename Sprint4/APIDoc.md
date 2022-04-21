# BackEnd API
## Customer API

### 1. Account Register API
```
POST    http://localhost:8000/registration
```

#### Examples:</br>
*Example1:*
Normal case.
##### Body
```
    {
      "Username": "zascauchy",
      "Password": "kkkna784984",
    }
```
##### Response:
```
{
      "code": 200,
      "status": "OK",
      "message": "User Registration Completed.",
      "response": {
          "username": "zascauchy",
          "userID": "622261e5f61f9a3ee8b25101",
          "online": ""
      }
    }
```
*Example2:*
wrong case without password input:
##### Body
```
{
    "Username": "ioih",
    "Password": ""
}
```
##### Response:
```
{
    "code": 500,
    "status": "Internal Server Error",
    "message": "Password can't be empty.",
    "response": null
}
```


### 2. User LogIn API
```
POST    http://localhost:8000/login
```
#### Examples:</br>
*Example 1:*</br>
Normal Case 
##### Body
```
{
    "Username": "zascauchy",
    "Password": "kkkna784984"
}
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "User Login is Completed.",
    "response": {
        "username": "zascauchy",
        "userID": "622261e5f61f9a3ee8b25101",
        "online": ""
    }
}
```
*Example 2:*</br>
Login with wrong password 
##### Body
```
{
    "Username": "c2233",
    "Password": "kkkna74"
}
```
##### Response:
```
{
    "code": 404,
    "status": "Not Found",
    "message": "Your Login Password is incorrect.",
    "response": null
}
```

*Example 3:*</br>
Login with unregistered username
##### Body
```
{
    "Username": "842233",
    "Password": "kkkna74"
}
```
##### Response:
```
{
    "code": 404,
    "status": "Not Found",
    "message": "This account does not exist in our system.",
    "response": null
}
```

### 3. User Session Check API
```
POST    http://localhost:8000/userSessionCheck/{userID}
```
#### Examples:
*Example 1:*
test case with logged in userID 622261e5f61f9a3ee8b25101
```
        http://localhost:8000/userSessionCheck/622261e5f61f9a3ee8b25101
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "You are logged in.",
    "response": true
}
```
*Example 2:*
test case with logged out userID 62211e70de6da79873505622
```
http://localhost:8000/userSessionCheck/62211e70de6da79873505622
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "You are logged in.",
    "response": false
}
```
### 4. Send Message for 1v1 Chatting API
Description: 
In this function, there are 3 event: join, message and disconnect. 
```
websocket Request =>   ws://localhost:8000/ws/fromUserID
```
#### Examples:</br>

*Example 1:*</br>
"join" Event Case with toUserID as 62211d78de6da79873505621  </br>
``` 
ws://localhost:8000/ws/62211d78de6da79873505621
```
##### Body:
```
{
    "eventName": "join",
    "eventPayload":
    "62211d78de6da79873505621"
}
```
*Example 2:*</br>
"message" Event Case with toUserID as 62211d78de6da79873505621  </br>
``` 
ws://localhost:8000/ws/62211d78de6da79873505621
```
##### Body:
```
{
    "eventName":    "message",
    "eventPayload":
    {
        "fromUserID": "62211d78de6da79873505621",
        "toUserID": "62211e70de6da79873505622",
        "message": "4 March Test"
    }
}
```
*Example 3:*</br>
"disconnect" Event Case with toUserID as 62211d78de6da79873505621  </br>
``` 
ws://localhost:8000/ws/62211d78de6da79873505621
```
##### Body:
```
{
    "eventName": "disconnect",
    "eventPayload":
    "62211d78de6da79873505621"
}
```
##### Response:
```
{
    "eventName": "chatlist-response",
    "eventPayload": {
        "type": "user-disconnected",
        "chatlist": {
            "username": "",
            "userID": "",
            "online": "N"
        }
    }
}
```

### 4. Get 1v1 Chatting Conversation API
Description: Retrive the messages history between two users
```
GET localhost:8000/getConversation/{toUserID}/{fromUserID}
```
#### Examples:</br>

*Example 1:*</br>
Normal Case with toUserID as 62211e74de6da79873505622 ,fromUserID as 62211e74de6da79873505621 </br>
```
GET localhost:8000/getConversation/62211e74de6da79873505622/62211e74de6da79873505621
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "Username is available.",
    "response": [
        {
            "id": "6222382aa638a22e8350a36b",
            "message": "messsssHey",
            "toUserID": "62211e74de6da79873505622",
            "fromUserID": "62211e74de6da79873505621"
        }
    ]
}
```
*Example 2:*</br>
If there is no messages from between this two users:
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "Username is available.",
    "response": null
}
```

### 5. Room Create API
```
POST    http://localhost:8000/CreateRoom
```
#### Examples:</br>
*Example 1:*</br>
Normal Case
##### Body
```
{
     "Username": "Cauchy",
     "UserID": "6260504c1190410dbe5babac",
     "RoomName": "2889",
     "GenerateRoomPassword":"Yes"
}
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "You have created a chat room",
    "response": {
        "roomNo": "26931255",
        "roomPassword": "I3a4tR(2#.",
        "roomName": "2889"
    }
}
```
*Example 2:*</br>
Error Case
##### Body
```

```
##### Response:
```

```

### 6. Join Group Chatting API
```
POST    http://localhost:8000/JoinRoom
```
#### Examples:
*Example 1:*</br>
Error Case =>
##### Body
```
{
     "Username": "Kexin Zhang",
     "RoomNo": "26931255",
     "RoomPassword":"I3a4tR(2#."
}
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "user joint the group chat",
    "response": "user joint the group chat"
}
```
*Example 2:*</br>
wrong password Case
##### Body
```
{
    "Username": "Kexin Zhang",
    "RoomNo": "26931255",
    "RoomPassword":"I3a4tR(2#"
}
```
##### Response:
```
{
    "code": 400,
    "status": "Bad Request",
    "message": "Kexin Zhang cannot join the group chat room",
    "response": "Password isn't correct"
}
```
*Example 3:*</br>
unlegal user case:
##### Body
```
{
     "Username": "Kexin Zhang",
     "RoomNo": "",
     "RoomPassword":"I3a4tR(2#."
}
```
##### Response:
```
{
    "code": 400,
    "status": "Bad Request",
    "message": "RoomNo can't be empty.",
    "response": null
}
```


### 7. Get Group Chatting Conversation API
```
POST    http://localhost:8000/getRoomChatConversation/{toChatID}/{fromUserID}
```
#### Examples:</br>

*Example 1:*</br>
Normal Case with RoomID as 29439400 ,fromUserID as f625f06024b1ed70f00d2a48e </br>
```
http://localhost:8000/getRoomChatConversation/29439400/f625f06024b1ed70f00d2a48e
```
##### Body
```
{
    "Username": "kkkkk",
    "UserID": "6247348aa456b8b53a5fce90",
    "RoomName": "29",
    "GenerateRoomPassword":"Yes"
}
```
##### Response:
```

```
*Example 2:*</br>
Error Case
##### Body
```

```
##### Response:
```

```

### 8. Send BroadCasting  Message API
```
POST  getBroadcast/{fromUserID}
```
#### Examples:</br>

*Example 1:*</br>
Normal Case with fromUserID as 6247348aa456b8b53a5fce90 </br>
```
http://localhost:8000/getBroadcast/f625f06024b1ed70f00d2a48e
```
##### Body
```
{
    "fromUserID": "6247348aa456b8b53a5fce90",
    "Message":"This is a broadcast Test"
}
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "Username is available.",
    "response": null
}
```

### 9. Send DriftBottle  Message API
```
POST  getDriftBottle/{toUserID}/{fromUserID}
```
#### Examples:</br>

*Example 1:*</br>
##### Body
```
{
    "fromUserID": "6247348aa456b8b53a5fce90",
    "toUserID": "f625f06024b1ed70f00d2a48e",
    "Message":"This is a DriftBottle Test"
}
```
##### Response:
```
{
    "code": 200,
    "status": "OK",
    "message": "Username is available.",
    "response": null
}
```
### 9. Add Friend API
```
websocket request  ws://localhost:8000/ws/{userID}
```
#### Examples:</br>

*Example 1:*</br>
```
ws://localhost:8000/ws/6260504c1190410dbe5babac
```
##### Body
```
{
    "eventName": "add_friends",
    "eventPayload": {
        "fromUserID": "6260504c1190410dbe5babac",
        "toUserID": "626050421190410dbe5baba",
        "message": "I love uuuuuuu",
    }
}
```
##### Response:</br>
friend side:

```
{
    "eventName": "frinedslist-response",
    "eventPayload": {
        "fromUserID": "6260504c1190410dbe5babac",
        "toUserID": "626050421190410dbe5babaa",
        "message": "Cauchy wants to add Cauchy Zhang as a friend",
        "image": ""
    }
}
```
user-side:
```
{
    "eventName": "frinedslist-response",
    "eventPayload": {
        "fromUserID": "626050421190410dbe5babaa",
        "toUserID": "6260504c1190410dbe5babac",
        "message": "You have been friends with Cauchy",
        "image": ""
    }
}
```

### get group Chat history API
```
websocket request  ws://localhost:8000/ws/{userID}
```
#### Examples:</br>

*Example 1:*</br>
new user joined
```
ws://localhost:8000/ws/6260504c1190410dbe5babac
```
##### Body
```
{
    "eventName": "new-user-joined",
    "eventPayload": {
        "fromUserID": "6260504c1190410dbe5babac",
        "toUserID": "626050421190410dbe5baba",
        "message": "I love uuuuuuu",
    }
}
```
##### Response:</br>

```
{
    "eventName": "chatlist-response",
    "eventPayload": {
        "type": "new-user-joined",
        "chatlist": {
            "username": "Cauchy",
            "userID": "6260504c1190410dbe5babac",
            "online": "Y"
        }
    }
}
```
*Example 2:*</br>
user disconnect
##### Response :</br>
```
{
    "eventName": "chatlist-response",
    "eventPayload": {
        "type": "user-disconnected",
        "chatlist": {
            "username": "Cauchy",
            "userID": "6260504c1190410dbe5babac",
            "online": "N"
        }
    }
}
```
*Example 3:*</br>
my-chat-list
##### Response :</br>
```
{
    "eventName": "chatlist-response",
    "eventPayload": {
        "type": "my-chat-list",
        "chatlist": [
            {
                "username": "Kiy123",
                "userID": "625f5dfc87dbe1be871557f0",
                "online": "Y"
            },
            {
                "username": "Kitty123",
                "userID": "625f6119276fadd8352c5e1d",
                "online": "Y"
            },
            {
                "username": "Kitty123",
                "userID": "625f61db09cce2204b13786d",
                "online": "Y"
            },
            {
                "username": "Nick",
                "userID": "62604231708c25d855218974",
                "online": "Y"
            },
            {
                "username": "Cauchy",
                "userID": "6260504c1190410dbe5babac",
                "online": "Y"
            },
            {
                "username": "Kity123",
                "userID": "62608d1b24fc98f890f2cab0",
                "online": "Y"
            },
            {
                "username": "Kity123",
                "userID": "62609f46ce19074375bc2791",
                "online": "Y"
            }
        ]
    }
}
```
