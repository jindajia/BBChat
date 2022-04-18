import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';

import "./joinmeeting.css"

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

function JoinMeeting() {
    return (
        <div>
            <div>
                <Navhome />
            </div>
            <div>
                <div className='number-container'>
                    <label className='label-style'>RoomNumber:</label>
                    <input className='input-style' />
                </div>
                <div className='password-container'>
                    <label className='label-style'>Roompassword:</label>
                    <input className='input-style' />
                </div>
                <div className='button-container'>
                    <button className='button-73'>JOIN</button>
                </div>
            </div>
        </div>
    );
}

export default withRouter(JoinMeeting);