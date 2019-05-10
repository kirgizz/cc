import React, { Component } from "react";
import Header from '../views/header';

class HeaderContainer extends Component {
    render() {
        return (
            <Header {...this}/>
        );
    }
}

export default HeaderContainer

