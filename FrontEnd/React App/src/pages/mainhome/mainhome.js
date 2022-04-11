import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter } from 'react-router-dom';

import './mainhome.css'

function Mainhome() {
    return (

        <div>
            <div className='container'>
                <div className='logo'></div>
                <div className='nav'>
                    <nav>
                        <ul>
                            <li><button className='navbutton'>Home</button></li>
                            <li><button className='navbutton'>Chat</button></li>
                            <li><button className='navbutton'>Register</button></li>
                            <li><button className='navbutton'>Login</button></li>
                        </ul>
                    </nav>
                </div>
            </div>

            <div class="content-3d">
                <button className='btn-3d green'>Hot Topic</button>
                <button className='btn-3d purple'>New Meeting</button>
                <button className='btn-3d cyan'>Join</button>
                <button className='btn-3d yellow'>Friends</button>
            </div>

        </div>
    );

}


export default Mainhome;