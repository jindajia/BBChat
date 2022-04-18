import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import { Tooltip, Whisper } from 'rsuite';
import { Button, IconButton, ButtonGroup, ButtonToolbar } from 'rsuite';
import { Input, InputGroup, MaskedInput } from 'rsuite';

import "./friend.css"


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

const tooltip1 = (
    <Tooltip className='tip'>
        Chat with a friend
    </Tooltip>
);

const tooltip2 = (
    <Tooltip className='tip'>
        Chat with a random online user
    </Tooltip>
);


function Friend() {
    return (
        <div>
            <div>
                <Navhome />
            </div>
            <div >
                <div className='add-container'>
                    <Whisper trigger="focus" speaker={<Tooltip className='tip'>Friend's Username</Tooltip>}>
                        <Input className='add-input' style={{ width: 280, height: 40 }} placeholder="Default Input" />
                    </Whisper>
                </div>
                <div className='add-button'>
                    <Whisper placement='right' controlId="control-id-hover" trigger="hover" speaker={tooltip1}>
                        <Button className='button-85'>Add Friend</Button>
                    </Whisper>
                </div>

            </div>

            <div className='random-container'>
                <Whisper placement='right' controlId="control-id-hover" trigger="hover" speaker={tooltip2}>
                    <Button className='button-85'>Random Chat</Button>
                </Whisper>

            </div>

        </div>
    );
}

export default withRouter(Friend);