import React from 'react';
import { Router, Route, browserHistory } from 'react-router';

// Layouts
import MainLayout from './components/layouts/main-layout';

// Pages
import Home from './components/home';
import ArticleContainer from './components/containers/articles-container';
import LoginContainer from './components/containers/login-container';

export default (

  <Router history={browserHistory}>
    <Route path="/Login" component={LoginContainer}/>
    <Route component={MainLayout}>
      <Route path="/" component={ArticleContainer} />
      <Route path="/home" component={Home}/>
    </Route>
  </Router>
);
