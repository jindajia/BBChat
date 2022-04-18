import React from 'react';
import './App.css';


import {
  BrowserRouter as Router,
  Route,
  Switch
} from "react-router-dom";

import Mainhome from './pages/mainhome/mainhome';
import Authentication from './pages/authentication/authentication';
import Home from './pages/home/home';
import FourOFour from './pages/four-o-four/four-o-four';
import Groupchat from './pages/groupchat/groupchat'
import CreateRoom from './pages/createRoom/createRoom'
import Friend from './pages/friend/friend';
import JoinMeeting from './pages/joinmeeting/joinmeeting';

function App() {
  return (
     <Router>
        <Switch>
          <Route path="/" exact component={Mainhome} />
          <Route path="/authentication" component={Authentication} />
          <Route path="/home/" component={Home} />
          <Route path="/createroom/" component={CreateRoom} />
          <Route path="/groupchat/" component={Groupchat} />
          <Route path="/friend/" component={Friend} />
          <Route path="/joinmeeting/" component={JoinMeeting} />
          <Route component={FourOFour} />
        </Switch>
      </Router>
  );
}

export default App;
