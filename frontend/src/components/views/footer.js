import React from 'react';
import { Link } from 'react-router';

//import Home from '../home';


function Footer(props) {
  return (
    <footer>
    <div className="footer-nav">
      <ul id="footer-menu">
        <li><Link to='/'>Home</Link></li>
        <li><Link to='/publications'>Home</Link></li>
        <li><Link to='/users'>Home</Link></li>
        <li><Link to='/rubrics'>Home</Link></li>
        <li><Link to='/companies'>Home</Link></li>
        <li><Link to='/sandbox'>Home</Link></li>
      </ul>
    </div>
    </footer>
  )
}

export default Footer
