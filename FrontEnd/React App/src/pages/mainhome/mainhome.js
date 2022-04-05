import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter } from 'react-router-dom';

import './mainhome.css'

function Mainhome() {
    return (
        <div class="content-3d">
            <button className='btn-3d green'>Hot Topic</button>
            <button className='btn-3d purple'>Create MeetingRoom</button>
            <button className='btn-3d cyan'>Join MeetingRoom</button>
            <button className='btn-3d yellow'>Friends</button>
        </div>

    );

}


export default Mainhome;