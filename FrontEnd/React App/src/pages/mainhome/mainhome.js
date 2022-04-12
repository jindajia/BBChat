import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter } from 'react-router-dom';
import moment from 'moment'

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

            <div className='tp-cl-container'>
                <div className='clock-container'>
                    <Clock />
                </div>
                <div className='tp-container'>

                </div>
            </div>

            <div class="content-3d">
                <button className='btn-3d green'>Hot Topic</button>
                <button className='btn-3d purple'>New Meeting</button>
                <button className='btn-3d cyan'>Join Meeting</button>
                <button className='btn-3d yellow'>Friends</button>
            </div>

        </div>
    );

}


export default Mainhome;