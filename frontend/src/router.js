import React from 'react';
import { Router, Route, browserHistory, IndexRoute } from 'react-router';

// Layouts
import MainLayout from './components/layouts/main-layout';

// Pages
import Home from './components/home';
import ArticleContainer from './components/containers/articles-container';

export default (
  <Router history={browserHistory}>
    <Route component={MainLayout}>
      <Route path="/" component={ArticleContainer} />
      <Route path="/home" component={Home}/>
    </Route>
  </Router>
);