import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';

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


function Friend() {
    return (
    <div>
        <div>
            <Navhome />
        </div>
    </div>
    );
}

export default withRouter(Friend);