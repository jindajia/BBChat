import React, { useState, useEffect } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { Tooltip, Whisper } from 'rsuite';
import { Button, IconButton, ButtonGroup, ButtonToolbar } from 'rsuite';
import { Input, InputGroup, MaskedInput } from 'rsuite';
import { userSessionCheckHTTPRequest } from "./../../services/api-service";
import {
    connectToWebSocket,
    listenToWebSocketEvents,
    eventEmitter
  } from './../../services/socket-service';
import {
    getItemInLS,
    removeItemInLS
  } from "./../../services/storage-service";

import "./friend.css"


const getUserName = (userDetails) => {
    if (userDetails && userDetails.username) {
        return userDetails.username;
    }
    return '___';
};

function Navhome() {

    const userDetails = getItemInLS('userDetails');
    return (
        <div className='container'>
            <div className='logo'></div>
            <div className='nav'>
                <nav>
                    <ul>
                        <Link to={"/mainhome"}>
                            <li><button className='navbutton'>Home</button></li>
                        </Link>
                        <Link to={"/groupchat"}>
                            <li><button className='navbutton'>Chat</button></li>
                        </Link>
                        <li><button className='button-53'>{getUserName(userDetails)}</button></li>
                    </ul>
                </nav>
            </div>
        </div>
    );
}

const tooltip1 = (
    <Tooltip className='tip'>
        Chat with a friend
    </Tooltip>
);

const tooltip2 = (
    <Tooltip className='tip'>
        Chat with a random online user
    </Tooltip>
);


function Friend(props) {
    const userDetails = getItemInLS('userDetails');
    const [internalError, setInternalError] = useState(null);
    const [friendList, setFriendList] = useState([]);

    const friendListSubscription = (socketPayload) => {
        let newFriendList = friendList;
    
    };
    
    useEffect(() => {
        eventEmitter.on('frinedslist-response', friendListSubscription);
        return () => {
            eventEmitter.removeListener('frinedslist-response', friendListSubscription);
        };
    });
    useEffect(() => {

        (async () => {
          if (userDetails === null || userDetails === '') {
            console.log("userDetails is null");
            // props.history.push(`/authentication`);
          } else {
            const webSocketConnection = connectToWebSocket(userDetails.userID);
            if (webSocketConnection.webSocketConnection === null) {
              setInternalError(webSocketConnection.message);
            } else {
              listenToWebSocketEvents()
            }
          }
        })();
    
      }, [props, userDetails]);
    return (
        <div>
            <div>
                <Navhome />
            </div>
            <div >
                <div className='add-container'>
                    <Whisper trigger="focus" speaker={<Tooltip className='tip'>Friend's Username</Tooltip>}>
                        <Input className='add-input' style={{ width: 280, height: 40 }} placeholder="Default Input" />
                    </Whisper>
                </div>
                <div className='add-button'>
                    <Whisper placement='right' controlId="control-id-hover" trigger="hover" speaker={tooltip1}>
                        <Button className='button-85'>Add Friend</Button>
                    </Whisper>
                </div>

            </div>

            <div className='random-container'>
                <Whisper placement='right' controlId="control-id-hover" trigger="hover" speaker={tooltip2}>
                    <Button className='button-85'>Random Chat</Button>
                </Whisper>

            </div>

        </div>
    );
}

export default withRouter(Friend);