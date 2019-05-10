import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, browserHistory } from 'react-router';
import MainLayout from './components/layouts/main-layout';
import About from './components/views/about';
import Feedback from './components/views/feedback';

import './styles/base.css'
import './styles/about.css'
import './styles/header.css'
import './styles/footer.css'
import './styles/feedback.css'


require('es6-promise').polyfill();

const Routes = () => (

    <Router history={browserHistory}>
        <Route component={MainLayout}>
            <Route path="/" component={About} />
            <Route path="/feedback" component={Feedback}/>
            <Route path="/about" component={About}/>
        </Route>
    </Router>
);


ReactDOM.render((
    <Routes />
), document.getElementById('app'))
