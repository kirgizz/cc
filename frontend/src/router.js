import React from 'react';
import { Router, Route, browserHistory } from 'react-router';

// Layouts
import MainLayout from './components/layouts/main-layout';

// Pages
import Home from './components/views/home';
import About from './components/views/about';
import ArticleContainer from './components/containers/articles-container';
import LoginContainer from './components/containers/login-container';
import RegisterContainer from './components/containers/register-container';

export default (

  <Router history={browserHistory}>
    <Route component={MainLayout}>
        <Route path="/Login" component={LoginContainer}/>
        <Route path="/Register" component={RegisterContainer}/>
      <Route path="/" component={ArticleContainer} />
      <Route path="/about" component={About}/>
    </Route>
  </Router>
);
