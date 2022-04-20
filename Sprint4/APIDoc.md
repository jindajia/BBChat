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

### 6. Join Group Chatting API
```
POST    http://localhost:8000/JoinRoom
```
#### Examples:
*Example 1:*</br>
Normal Case
##### Body
```
{
     "Username": "Cauchy",
     "RoomNo": "29",
     "RoomPassword":"$@%ky,ub+B" 
}
```
##### Response:
```

```
*Example 2:*</br>
Error Case =>
##### Body
```

```
##### Response:
```

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

