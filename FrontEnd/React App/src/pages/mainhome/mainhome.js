import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter } from 'react-router-dom';


function Main() {
    return (
        <div>
            <div>
                <button>Hot Topic</button>
            </div>
            <div>
                <button>Create MeetingRoom</button>
            </div>
            <div>
                <button>Join MeetingRoom</button>
            </div>
            <div>
                <button>Friends</button>
            </div>
        </div>

    );

}

// ReactDOM.render(
//     <Main />,
//     document.getElementById('root')
// );

export default withRouter(Main);