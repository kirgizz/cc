import React, { Component } from 'react';

function Header(props) {
  return (
    <header>
    <div classname="topnav">
      <ul id="header-menu">
        <li><a href="publications.html">publications</a></li>
        <li><a href="users.html">users</a></li>
        <li><a href="rubrics.html">reubrics</a></li>
        <li><a href="companies.html">companies</a></li>
        <li><a href="sandbox.html">sandbox</a></li>
          <div classname="right">
            <li><a id="login "href="login.html">publications</a></li>
            <li><a id="register" href="register.html">register</a></li>
          </div>
      </ul>
    </div>
    </header>
  )

}

class App extends Component {
  render() {
    return (
			<Header />
    );
  }
}

export default App;
