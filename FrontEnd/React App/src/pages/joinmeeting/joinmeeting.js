import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import {  getItemInLS, removeItemInLS, setItemInLS} from "./../../services/storage-service";
import { userJoinRoom } from "./../../services/api-service";


import "./joinmeeting.css"
import {ReactNotifications, Store } from 'react-notifications-component'
import 'react-notifications-component/dist/theme.css'

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
                        <Link to={"/home"}>
                            <li><button className='navbutton'>Chat</button></li>
                        </Link>
                        <li><button className='button-53'>{getUserName(userDetails)}</button></li>
                    </ul>
                </nav>
            </div>
        </div>
    );
}

function JoinMeeting(props) {

    const [roomNumber, updateRoomNumber] = useState("");
    const [roomPassword, updateRoomPassword] = useState("");
    const userDetails = getItemInLS('userDetails');
    const handleroomNumberChange = (event) => {
        updateRoomNumber(event.target.value);
      }
    
    const handleRoomPasswordChange = (event) => {
        updateRoomPassword(event.target.value);
    }
    const joinRoom = async () => {
        // props.displayPageLoader(true);
        const joinRoomMessage = {}
        joinRoomMessage.username = userDetails.username;
        joinRoomMessage.roomNo = roomNumber;
        joinRoomMessage.roomPassword = roomPassword;
        console.log(joinRoomMessage);
        var joinResponse = null;
        try{
            joinResponse = await userJoinRoom(joinRoomMessage);
        } catch(err) {
          console.log(err);
          Store.addNotification({
            title: "Join Room Failed!",
            message: "connection error",
            type: "danger",
            insert: "top",
            container: "top-right",
            animationIn: ["animate__animated", "animate__fadeIn"],
            animationOut: ["animate__animated", "animate__fadeOut"],
            dismiss: {
              duration: 5000,
              onScreen: true
            }
          });
        }
        // props.displayPageLoader(false);
    
        if (joinResponse.code===200){
            setItemInLS('chatRoomNo', roomNumber);
            setItemInLS('chatRoomPass', roomPassword);
            setTimeout(() => {props.history.push(`/groupchat`); }, 3000);
            console.log(joinResponse.response)
            Store.addNotification({
              title: "Join Room Success!",
              message: joinResponse.response,
              type: "success",
              insert: "top",
              container: "top-right",
              animationIn: ["animate__animated", "animate__fadeIn"],
              animationOut: ["animate__animated", "animate__fadeOut"],
              dismiss: {
                duration: 3000,
                onScreen: true
              }
            });
        } else {
          console.log(joinResponse.code)
          Store.addNotification({
            title: "Join Room Failed!",
            message: joinResponse.response,
            type: "danger",
            insert: "top",
            container: "top-right",
            animationIn: ["animate__animated", "animate__fadeIn"],
            animationOut: ["animate__animated", "animate__fadeOut"],
            dismiss: {
              duration: 5000,
              onScreen: true
            }
          });
        }
    };

    return (
        <div>
            <div>
                <Navhome />
            </div>
            <div>
                <ReactNotifications />
                <div className='number-container'>
                    <label className='label-style'>RoomNumber:</label>
                    <input id='roomnumber' className='input-style' value={roomNumber} onChange={handleroomNumberChange} />
                </div>
                <div className='password-container'>
                    <label className='label-style'>Roompassword:</label>
                    <input id='password' className='input-style' value={roomPassword} onChange={handleRoomPasswordChange}/>
                </div>
                <div className='button-container'>
                    <button className='button-73' onClick={joinRoom}>JOIN</button>
                </div>
            </div>
        </div>
    );
}

export default withRouter(JoinMeeting);