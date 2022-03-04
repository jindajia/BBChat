sprint2 Finished Task:
Add new features and integrated with frontend.
Finished websocket related problems and can use websocket to  send and deliver messages now.

By 4 March, the following features/functions has already been fixed:

1. new user sign up
new users can sign up with their emails and will be assigned a distinct userID.
user information will be stored in database/user

2.new user sign in
Forntend using post method sending the input to backend.
Comparing the input with the data stored in our database, if it matches, the user could successfully get access into the chatting room.
The webpage jumps and the users statu info update to 'Y' in the database

3.users send messages to others.
There are two kind of message sending.
One is send message to a certain user, the other is sending messages in the chatting room and the messages will be broadcast to everyone in the chatting room.
the content of messages and other informations related to this conversation will be stored in database/message.
 
4.user logout
Frontend sends disconnect message to the server.  
User status in the database uodates to 'N'

For the unit test part, we use postman to test each of our main feature functions. The integrated test will be done by our group members who takes charge of the frontend side.
I store the test cases in the backEndUnitTestCases.txt. And I will upload a video to show how they works.