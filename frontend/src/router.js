import React from 'react';
import { Router, Route, browserHistory } from 'react-router';

// Layouts
import MainLayout from './components/layouts/main-layout';

import About from './components/views/about';
import feedback from './components/views/feedback';

export default (

  <Router history={browserHistory}>
    <Route component={MainLayout}>
      <Route path="/" component={About} />
      <Route path="/about" component={About}/>
      <Route path="/feedback" component={feedback}/>
    </Route>
  </Router>
);


//      <Route path="/" component={ArticleContainer} />
