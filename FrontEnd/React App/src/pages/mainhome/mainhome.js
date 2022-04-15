import React from 'react';
import ReactDOM from 'react-dom';
import { Link, withRouter } from 'react-router-dom';
import moment from 'moment'
import { Navbar, Nav, Dropdown, Affix } from 'rsuite';


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

function Mainhome() {
    return (

        <div>
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
                <button className='btn-3d cyan'>Join Meeting</button>
                <button className='btn-3d yellow'>Friends</button>
            </div>

        </div>
    );

}


export default Mainhome;