import React, { useState, useEffect, useRef } from 'react';
import {
  eventEmitter,
  sendWebSocketPayload
} from '../../../services/socket-service';
import { getConversationBetweenGroups } from '../../../services/api-service';

import './roomconversation.css'

const alignMessages = (userDetails, toUserID) => {
  const { userID } = userDetails;
  return userID === toUserID;
}

const scrollMessageContainer = (messageContainer) => {
  if (messageContainer.current !== null) {
    try {
      setTimeout(() => {
        messageContainer.current.scrollTop = messageContainer.current.scrollHeight;
      }, 100);
    } catch (error) {
      console.warn(error);
    }
  }
}

const getMessageUI = (messageContainer, userDetails, roomconversations) => {
  return (
    <ul ref={messageContainer} className='message-thread-container'>
      {roomconversations.map((roomconversation, index) => (
        <li
          className={`message ${
            alignMessages(userDetails, roomconversation.fromUserID) ? 'align-right' : ''
          }`}
          key={index}
        >
          {roomconversation.message}
        </li>
      ))}
    </ul>
  );
}

const getInitiateroomconversationUI = (userDetails) =>{
  if (userDetails !== null) {
    return (
      <div className="message-thread-container start-chatting-banner">
        <p className="heading">
          You haven 't chatted with {userDetails.username} in a while,
          <span className="sub-heading"> Say Hi.</span>
        </p>			
      </div>
    )
  }    
}

function RoomConversation(props) {
  const roomNumber = props.roomNumber;
  const userDetails = props.userDetails;

  const messageContainer = useRef(null);
  const [roomconversation, updateroomconversation] = useState([]);
  const [messageLoading, updateMessageLoading] = useState(true);

  useEffect(() => {
    if (userDetails && roomNumber) {
      (async () => {
        const roomconversationsResponse = await getConversationBetweenGroups(roomNumber, userDetails.userID);
        updateMessageLoading(false)
        if (roomconversationsResponse.response) {
          updateroomconversation(roomconversationsResponse.response);
        } else if (roomconversationsResponse.response === null) {
          updateroomconversation([]);
        }
        scrollMessageContainer(messageContainer);
      })();
    }
  }, [userDetails, roomNumber])

  useEffect(() => {
    const newMessageSubscription = (messagePayload) => {
      console.log(messagePayload);
      updateroomconversation([...roomconversation, messagePayload]);
      scrollMessageContainer(messageContainer);
    };

    eventEmitter.on('chatmessage-response', newMessageSubscription);

    return () => {
      eventEmitter.removeListener('chatmessage-response', newMessageSubscription);
    };
  });

  const sendMessage = (event) => {
    if (event.key === 'Enter') {
      const message = event.target.value;

      if (message === '' || message === undefined || message === null) {
        alert(`Message can't be empty.`);
        event.target.value = '';
      } else if (userDetails.userID === '') {
        this.router.navigate(['/']);
      } else if (roomNumber === undefined) {
        alert(`Select a group to chat.`);
      } else {
        event.target.value = '';

        const messagePayload = {
          fromUserID: userDetails.userID,
          message: message.trim(),
          toUserID: roomNumber,
        };
        // console.log("senwebsocketpayload ", messagePayload)
        sendWebSocketPayload("room-chat", messagePayload);
        // updateroomconversation([...roomconversation, messagePayload]);
        // scrollMessageContainer(messageContainer);
      }
    }
  }

  if (messageLoading) {
    return (
      <div
        className="message-overlay"
      >
        <h3>
          {roomNumber !== null
            ? 'Loading Messages'
            : ' Select a Room to chat.'}
        </h3>
      </div>
    )
  }

  return (
    <div className='app__conversion-container'>
      {getMessageUI(messageContainer, userDetails, roomconversation)}
      <div className='app__text-container'>
        <textarea
          placeholder={`${
            roomNumber !== null ? '' : 'Select a user and'
          } Type your message here`}
          className='text-type'
          onKeyPress={sendMessage}
        ></textarea>
      </div>
    </div>
  );
}

export default RoomConversation;