import React from 'react';
import { Link } from 'react-router';
//import Home from '../home';
import './styles/header.css'

function Header(props) {
  return (
    <header>
      <div className="topnav">
          <div className="logo">
              <div className="Logo-img"> </div>
          </div>

      <ul id="header-menu">
        <li><Link to='/about'>about</Link></li>
        <li><Link to='/publications'>publications</Link></li>
        <li><Link to='/users'>users</Link></li>
        <li><Link to='/rubrics'>rubrics</Link></li>
        <li><Link to='/companies'>companies</Link></li>
        <li><Link to='/sandbox'>sandbox</Link></li>


          <div className="authMenu">
              {

                  props.state.isAuth ? (
                      <div>
                          <li id="logout"><Link to='/logout'>logout</Link></li>
                          <li id="hello" > hello user</li>
                      </div>
                      ) : (
                      <div>
                          <li id="login"><Link to='/login'>Login</Link></li>
                          <li id="register"><Link to='/register'>Register</Link></li>
                      </div>
                      )
              }
          </div>

      </ul>
    </div>
    </header>
  );
}

export default Header
