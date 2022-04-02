import React from 'react';
import './App.css';


import {
  BrowserRouter as Router,
  Route,
  Switch
} from "react-router-dom";

import Authentication from './pages/authentication/authentication';
import Home from './pages/home/home';
import FourOFour from './pages/four-o-four/four-o-four';
import Groupchat from './pages/groupchat/groupchat'
import CreateRoom from './pages/createRoom/createRoom'
function App() {
  return (
     <Router>
        <Switch>
          <Route path="/" exact component={Authentication} />
          <Route path="/home/" component={Home} />
          <Route path="/createroom/" component={CreateRoom} />
          <Route path="/groupchat/" component={Groupchat} />
          <Route component={FourOFour} />
        </Switch>
      </Router>
  );
}

export default App;
