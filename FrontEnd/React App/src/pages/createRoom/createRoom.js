import React, { useState, useEffect } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { userCreateRoom } from "./../../services/api-service";
import {
  getItemInLS,
  setItemInLS
} from "./../../services/storage-service";

import './createRoom.css'
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

function CreateRoom(props) {

  const [roomName, updateRoomName] = useState("");
  const [roomPassword, updateRoomPassword] = useState("");
  const [createRoomMessage, updatecreateRoomMessage] = useState({});
  const [checked, setChecked] = React.useState(false);

  const userDetails = getItemInLS('userDetails');

  const handleroomNumberChange = (event) => {
    updateRoomName(event.target.value);
  }

  const handleRoomPasswordChange = (event) => {
    updateRoomPassword(event.target.value);

  }

  const handleChange = () => {
    setChecked(!checked);
  };

  useEffect(() => {
    if (checked) {
      updateRoomPassword('');
    }
  },[checked]);

  const createRoom = async () => {
    // props.displayPageLoader(true);
    createRoomMessage.username = userDetails.username;
    createRoomMessage.userID = userDetails.userID;
    createRoomMessage.roomNo = "1000";
    if (roomPassword===null||roomPassword===""){
      createRoomMessage.generateRoomPassword = "Yes";
    } else {
      createRoomMessage.generateRoomPassword = "No";
    }
    createRoomMessage.roomPassword = roomPassword;
    createRoomMessage.roomName = roomName;
    console.log(createRoomMessage);
    var roomDetail = null;
    try{
      roomDetail = await userCreateRoom(createRoomMessage);
    } catch(err) {
      console.log(err);
      Store.addNotification({
        title: "Create Room Failed!",
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

    if (roomDetail===null){
      console.log("Error");
    } else if (roomDetail.code === 200) {
      setItemInLS('roomDetail', roomDetail.response)
      console.log(roomDetail.response)
      Store.addNotification({
        title: "Create Room Success!",
        message: "Room: \"".concat(roomDetail.response.roomName, "\" create success!", " RoomNumber: ", roomDetail.response.roomNo ," Passowrd: ", roomDetail.response.roomPassword),
        type: "success",
        insert: "top",
        container: "top-right",
        animationIn: ["animate__animated", "animate__fadeIn"],
        animationOut: ["animate__animated", "animate__fadeOut"],
        dismiss: {
          duration: 5000,
          onScreen: true
        }
      });
    } else {
      // setErrorMessage(roomDetail.message);
      console.log(roomDetail.error)
    }
  };

  return (

    <div>
      <div>
        <Navhome />
      </div>

      <div>
        <ReactNotifications />
        <div className='create-room'>
          <label className='label-style'>RoomName:</label>
          <input id="roomname" className='input-style' value={roomName} onChange={handleroomNumberChange} />
        </div>
        <div className='create-password'>
          <label className='label-style'>Roompassword:</label>
          <input id="roompassword" className='input-style' value={roomPassword} onChange={handleRoomPasswordChange} disabled={checked} />
          <label className='checkboxlabel-style'>
            <input id='randompassword' className='checkbox-style' type="checkbox" checked={checked} onChange={handleChange} />
            Random Password
          </label>
        </div>

        <div className='create_button'>
          <button className='button-49' onClick={createRoom}>CreateRoom</button>
        </div>
      </div>
    </div>
  );
}


export default withRouter(CreateRoom);