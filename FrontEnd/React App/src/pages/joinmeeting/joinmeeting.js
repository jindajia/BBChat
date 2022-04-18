import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import {getItemInLS} from "./../../services/storage-service";


import "./joinmeeting.css"

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