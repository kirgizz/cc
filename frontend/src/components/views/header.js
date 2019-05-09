import React from 'react';
import { Link } from 'react-router';
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

      </ul>
    </div>
    </header>
  );
}

export default Header
