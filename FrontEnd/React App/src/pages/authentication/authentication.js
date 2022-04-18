import React, { useState,useEffect } from 'react';

import './authentication.css';

import Login from './login/login';
import Registration from './registration/registration'
import { withRouter } from 'react-router-dom';
import {
  connectToWebSocket,
  listenToWebSocketEvents,
  eventEmitter
} from './../../services/socket-service';
import {
  getItemInLS,
  removeItemInLS
} from "./../../services/storage-service";
function Authentication(props) {
  const userDetails = getItemInLS('userDetails');
  const [activeTab, setTabType] = useState('login');
  const [loaderStatus, setLoaderStatus] = useState(false);

  const changeTabType = (type) => {
    setTabType(type);
  }

  const getActiveClass = (type) => {
    return type === activeTab ? 'active' : '';
  };

  const displayPageLoader = (shouldDisplay) => {
    setLoaderStatus(shouldDisplay)
  }
  useEffect(() => {

    (async () => {
      if (userDetails === null || userDetails === '') {
        console.log("user not log in");
      } else {
        const webSocketConnection = connectToWebSocket(userDetails.userID);
        if (webSocketConnection.webSocketConnection === null) {
          // setInternalError(webSocketConnection.message);
          props.history.push(`/authentication`);
        } else {
          listenToWebSocketEvents()
          props.history.push(`/`);
        }
      }
    })();

  }, [props, userDetails]);
  return (
    <React.Fragment>
      <div className={`app__loader ${loaderStatus ? 'active': ''}`}>
        <img alt="Loader" src="/loader.gif"/>
      </div>
      <div className='app__authentication-container'>
        <div className='authentication__tab-switcher'>
          <button
            className={`${getActiveClass('login')} authentication__tab-button`}
            onClick={() => changeTabType('login')}
          >
            Login
          </button>
          <button
            className={`${getActiveClass('registration')} authentication__tab-button`}
            onClick={() => changeTabType('registration')}
          >
            Registration
          </button>
        </div>
        <div className='authentication__tab-viewer'>
          {activeTab === 'login' ? <Login displayPageLoader={displayPageLoader}/> : <Registration displayPageLoader={displayPageLoader}/>}
        </div>
      </div>
    </React.Fragment>
  );
}

export default withRouter(Authentication);