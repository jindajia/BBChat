import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import moment from 'moment'
import {
    connectToWebSocket,
    listenToWebSocketEvents,
    emitLogoutEvent,
} from './../../services/socket-service';
import {
    getItemInLS,
    removeItemInLS
} from "./../../services/storage-service"; import { Popover, Dropdown, ButtonToolbar, IconButton, Divider } from 'rsuite';
import ArrowDownIcon from '@rsuite/icons/ArrowDown';
import { Button, ButtonGroup, Whisper } from 'rsuite';
import { Modal } from 'rsuite';



import './mainhome.css'

function FormattedDate(props) {
    return (
        <div>
            <h2 className='year'>{moment().format('YYYY-MM-DD')}</h2>
            <h1 className='hour'>{moment().format('HH:mm:ss')}</h1>
        </div>
    );
}

class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.state = { date: new Date() };
    }

    componentDidMount() {
        this.timeID = setInterval(() => this.tick(), 1000);
    }

    componentWillUnmount() {
        clearInterval(this.timeID);
    }

    tick() {
        this.setState({
            date: new Date(),
        });
    }

    render() {
        return (
            <div>
                <FormattedDate date={this.state.date} />
            </div>
        );
    }
}

class Hottopic extends React.Component {
    state = {
        topic: [
            { id: 1, content: 'Afghanistan' },
            { id: 2, content: 'AMC Stock' },
            { id: 3, content: 'COVID Vaccine' },
            { id: 4, content: 'Dogecoin' },
            { id: 5, content: 'GME Stock' },
            { id: 6, content: 'Stimulus Check' },
            { id: 7, content: 'Georgia Senate Race' },
            { id: 8, content: 'Hurricane Ida' },
            { id: 9, content: 'COVID' },
            { id: 10, content: 'Ethereum Price' }
        ]
    }

    render() {
        return (
            <div>
                <ul>
                    {this.state.topic.map(item => (
                        <li key={item.id}>
                            <h3 className='hp'>{item.id} : {item.content}</h3>
                        </li>
                    ))}
                </ul>
            </div>
        );
    }
}

const getUserName = (userDetails) => {
    if (userDetails && userDetails.username) {
        return userDetails.username;
    }
    return '___';
};

const logoutUser = (props, userDetails) => {
    if (userDetails !== null && userDetails.userID) {
        removeItemInLS('userDetails');
        emitLogoutEvent(userDetails.userID);
        props.history.push('/');
    }
};


const AddFriendRequest = () => {
    const [open, setOpen] = React.useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);
    return (
        <div>
            <ButtonToolbar>
                <Button className='logout' href='#' onClick={handleOpen}> Add Friend Request</Button>
            </ButtonToolbar>

            <Modal className='modal-container' open={open} onClose={handleClose}>
                <Modal.Header>
                    <Modal.Title className='modaltitle'>Add Friend Request</Modal.Title>
                </Modal.Header>
                <Modal.Body className='modalbody'>
                   <p>JD wants to add you as friend</p>
                </Modal.Body>
                <Modal.Footer className='modalfooter'>
                    <Button className='button-24' onClick={handleClose} appearance="primary">
                        Ok
                    </Button>
                    <Button className='button-12' onClick={handleClose} appearance="subtle">
                        Cancel
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
};


const renderMenu = ({ left, top, className }, ref, props, userDetails) => {

    return (
        <div >
            <Popover ref={ref} className={className} style={{ left, top }} full>
                <Dropdown.Menu className='dropdown-container'>
                    <Dropdown.Item>
                        <p>Signed in as</p>
                        <strong id='username'>{getUserName(userDetails)}</strong>
                    </Dropdown.Item>
                    <Divider className='divider' />
                    <Dropdown.Item>
                        <AddFriendRequest />
                    </Dropdown.Item>
                    <Divider className='divider' />
                    <Dropdown.Item>
                        <button id='logout' className='logout' href='#' onClick={() => {
                            logoutUser(props, userDetails);
                        }} >Sign out
                        </button>
                    </Dropdown.Item>
                </Dropdown.Menu>
            </Popover>
        </div>
    );
};

function Navhome(props) {
    const mainhomeprops = props.mainhomeprops;
    const userDetails = getItemInLS('userDetails');
    useEffect(() => {

        (async () => {
          if (userDetails === null || userDetails === '') {
            console.log("user not log in");
            mainhomeprops.history.push(`/`);
        }
        })();

    }, [props, userDetails]);
    return (
        <div className='container'>
            <div className='logo'></div>
            <div className='nav'>
                <nav>
                    <ul>
                        <Link to={"/mainhome"}>
                            <li><button className='navbutton'>Home</button></li>
                        </Link>
                        <Link to={"/home"}>
                            <li><button className='navbutton'>Chat</button></li>
                        </Link>
                        <li><button className='button-53'>{getUserName(userDetails)}</button></li>
                        <li className='buttongroup'>
                            <Whisper trigger="click" speaker={renderMenu({}, null, mainhomeprops, userDetails)}>
                                <IconButton id='icon' className='icon-container' icon={<ArrowDownIcon />} />
                            </Whisper>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>
    );
}

function Mainhome(props) {
    return (

        <div>
            <div>
                <Navhome mainhomeprops={props} />
            </div>


            <div className='tp-cl-container'>
                <div className='clock-container'>
                    <Clock />
                </div>
                <div className='tp-title'>
                    <h1>Today's Hot Topic</h1>
                </div>
                <div className='tp-container-out'>
                    <div className='tp-container-in'>
                        <Hottopic />
                    </div>
                </div>
            </div>

            <div class="content-3d">
                <button className='btn-3d green'>Hot Topic</button>
                <Link to={"/createRoom"}>
                    <button className='btn-3d purple'>New Meeting</button>
                </Link>
                <Link to={"/joinmeeting"}>
                    <button className='btn-3d cyan'>Join Meeting</button>
                </Link>
                <Link to={"/friend"}>
                    <button className='btn-3d yellow'>Friends</button>
                </Link>
            </div>

        </div>
    );

}

export default withRouter(Mainhome);