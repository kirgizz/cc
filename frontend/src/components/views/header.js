import React from 'react';
import { Link } from 'react-router';

import Home from '../home';

function Header(props) {
  return (
    <header>
    <div classname="topnav">
      <ul id="header-menu">
        <li><Link to='/'>Home</Link></li>
        <li><Link to='/publications'>Home</Link></li>
        <li><Link to='/users'>Home</Link></li>
        <li><Link to='/rubrics'>Home</Link></li>
        <li><Link to='/companies'>Home</Link></li>
        <li><Link to='/sandbox'>Home</Link></li>
        <div classname="right">
          <li id="login"><Link to='/login'>Home</Link></li>
          <li id="register"><Link to='/register'>Home</Link></li>
        </div>
      </ul>
    </div>
    </header>
  );
}

export default Header
