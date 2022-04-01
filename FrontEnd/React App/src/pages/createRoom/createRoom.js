import React, { useState } from 'react';
import ReactDOM from 'react-dom';


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
    <div>
      <div>
        <label>Roomnumber:</label>
        <input onChange={handleroomNumberChange} />
      </div>
      <div>
        <label>Roompassword:</label>
        <input onChange={handleRoomPasswordChange} />
      </div>
      <div>
        <button>CreateRoom</button>
      </div>
    </div>
  );
}



ReactDOM.render(
  <CreateRoom />,
  document.getElementById('root')
);

