import React from 'react';
import { Link } from 'react-router';
import { Button } from "react-bootstrap";
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
        <li><Link to='/feedback'>feedback</Link></li>
        <li><Link to='/users'>users</Link></li>
        <li><Link to='/rubrics'>rubrics</Link></li>
        <li><Link to='/companies'>companies</Link></li>
        <li><Link to='/sandbox'>sandbox</Link></li>


          <div className="authMenu">
              {

                  props.state.isAuth ? (
                      <div>
                          <li id="logout"><Link to='/logout'>logout</Link></li>
                          <li id="logout"><Link to='/add'>write</Link></li>
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
