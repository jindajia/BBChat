import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import { userCreateRoom } from "./../../services/api-service";
import {
  getItemInLS,
  removeItemInLS,
  setItemInLS
} from "./../../services/storage-service";

import './createRoom.css'
import {ReactNotifications, Store } from 'react-notifications-component'
import 'react-notifications-component/dist/theme.css'


function Navhome() {

  return (
    <div className='container'>
      <div className='logo'></div>
      <div className='nav'>
        <nav>
          <ul>
            <Link to={"/#"}>
              <li><button className='navbutton'>Home</button></li>
            </Link>
            <Link to={"/groupchat"}>
              <li><button className='navbutton'>Chat</button></li>
            </Link>
            <Link to={"/authentication/registraion"}>
              <li><button className='navbutton'>Register</button></li>
            </Link>
            <Link to={"/authentication/login"}>
              <li><button className='navbutton'>Login</button></li>
            </Link>
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
    if (roomPassword==null||roomPassword==""){
      createRoomMessage.generateRoomPassword = "Yes";
    } else {
      createRoomMessage.generateRoomPassword = "No";
    }
    createRoomMessage.roomPassword = roomPassword;
    createRoomMessage.roomName = roomName;
    console.log(createRoomMessage);
    const roomDetail = null;
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
          <input className='input-style' value={roomName} onChange={handleroomNumberChange} />
        </div>
        <div className='create-password'>
          <label className='label-style'>Roompassword:</label>
          <input className='input-style' value={roomPassword} onChange={handleRoomPasswordChange} disabled={checked} />
          <label className='checkboxlabel-style'>
            <input className='checkbox-style' type="checkbox" checked={checked} onChange={handleChange} />
            Random Password
          </label>
        </div>

        <div className='create_button'>
          <button className='button-style2' onClick={createRoom}>CreateRoom</button>
        </div>
      </div>
    </div>
  );
}

// const docStyle = document.documentElement.style
// const aElem = document.querySelector('button')
// const boundingClientRect = aElem.getBoundingClientRect()

// aElem.onmousemove = function(e) {

// 	const x = e.clientX - boundingClientRect.left
// 	const y = e.clientY - boundingClientRect.top

// 	const xc = boundingClientRect.width/2
// 	const yc = boundingClientRect.height/2

// 	const dx = x - xc
// 	const dy = y - yc

// 	docStyle.setProperty('--rx', `${ dy/-1 }deg`)
// 	docStyle.setProperty('--ry', `${ dx/10 }deg`)

// }

// aElem.onmouseleave = function(e) {

// 	docStyle.setProperty('--ty', '0')
// 	docStyle.setProperty('--rx', '0')
// 	docStyle.setProperty('--ry', '0')

// }

// aElem.onmousedown = function(e) {

// 	docStyle.setProperty('--tz', '-25px')

// }

// document.body.onmouseup = function(e) {

// 	docStyle.setProperty('--tz', '-12px')

// }



// ReactDOM.render(
//   <CreateRoom />,
//   document.getElementById('root')
// );

export default withRouter(CreateRoom);