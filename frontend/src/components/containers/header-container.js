import React, { Component } from "react";
import Header from '../views/header';
import {connect} from "react-redux";


class HeaderContainer extends Component {
    render() {
        return (
            <Header {...this}/>
        );
    }
}

const mapStateToProps = function(store) {
    return {
        profile: store.profileState.profile
    };
};

export default connect(mapStateToProps)(HeaderContainer);

