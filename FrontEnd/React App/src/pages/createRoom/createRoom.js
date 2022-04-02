import React, { useState } from 'react';
import ReactDOM from 'react-dom';

import './createRoom.css'


function CreateRoom(props) {

  const [roomnumber, updateRoomNumber] = useState(null);
  const [roomPassword, updateRoomPassword] = useState(null);

  const handleroomNumberChange = (event) => {
    updateRoomNumber(event.target.value)
  }

  const handleRoomPasswordChange = (event) => {
    updateRoomPassword(event.target.value)
  }


  return (
    <div className='create_container'>
      <div className='create_1'>
        <label>RoomNumber:</label>
      </div>
      <div className='create_2'>
        <input onChange={handleRoomPasswordChange} />
      </div>
      <div className='create_1'>
        <label>Roompassword:</label>
      </div>
      <div className='create_2'>
        <input onChange={handleRoomPasswordChange} />
      </div>
      <div className='create_button'>
        <button>CreateRoom</button>
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



ReactDOM.render(
  <CreateRoom />,
  document.getElementById('root')
);

