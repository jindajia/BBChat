import React, {useState, useEffect} from 'react';
import { Content } from 'rsuite';

import {
  eventEmitter
} from "./../../services/socket-service";


function Randomfriend(props) {
  const [selectedUser, updateSelectedUser] = useState(null);
  const [chatList, setChatList] = useState([]);

  const chatListSubscription = (socketPayload) => {
    let newChatList = chatList;

    if (socketPayload.type === 'new-user-joined') {
      const incomingChatList = socketPayload.chatlist;
      if (incomingChatList) {
        const loggedInUserIndex = newChatList.findIndex(
          (obj) => obj.userID === incomingChatList.userID
        );
        if (loggedInUserIndex >= 0) {
          newChatList[loggedInUserIndex].online = 'Y';
        } else {
          newChatList = newChatList.filter(
            (obj) => obj.userID !== incomingChatList.userID
          );
          /* Adding new online user into chat list array */
          newChatList = [...newChatList, ...[incomingChatList]];
        }
      }

    } else if (socketPayload.type === 'user-disconnected') {
      const outGoingUser = socketPayload.chatlist;
      const loggedOutUserIndex = newChatList.findIndex(
        (obj) => obj.userID === outGoingUser.userID
      );
      if (loggedOutUserIndex >= 0) {
        newChatList[loggedOutUserIndex].online = 'N';
      }
    } else if (socketPayload.type ==='my-chat-list'){
      newChatList = socketPayload.chatlist;
    }

    // slice is used to create aa new instance of an array
    setChatList(newChatList.slice());
  };

  useEffect(() => {
    eventEmitter.on('chatlist-response', chatListSubscription);
    return () => {
      eventEmitter.removeListener('chatlist-response', chatListSubscription);
    };
  });
  
  const setSelectedUser = (user) => {
    if (user) {
      updateSelectedUser(user);
      props.updateSelectedUser(user); 
    }
  };

  if(chatList && chatList.length === 0) {
    return (
      <div className="alert">
        {chatList.length === 0 ? 'Loading your chat list.' : 'No User Available to chat.'}
      </div>
    );
  }

  const randomfriendlist = ['Raindrop'];
  
  return (
    <div className='app__chatlist-container'>
      <div className='user__chat-list'>
        {randomfriendlist.map((user) => (
          <div onClick={() => setSelectedUser(user)}>
            {user}
            <span className={user.online ==='online'}></span>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Randomfriend;