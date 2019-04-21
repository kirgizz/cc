import React, { Component } from "react";
import Header from '../views/header';
import {connect} from "react-redux";
import Cookies from 'js-cookie';
import * as authApi from "../../api/auth-api";
import {checkCredentials} from "../../api/auth-api";
import router from "../../router";

//import router from '../../router';

//import store from "../../store"
//import profileReducer from "../../reducers"

class HeaderContainer extends Component {
    constructor(props) {
        super(props);

        this.state = {
            isAuth: false,
        };

    }

    componentWillMount () {
        var ssid = Cookies.get("ssid")
        if (ssid ) {
            authApi.checkCookie({"ssid": ssid}).then(
                function(result) {
                    if (result.status === 200) {
                        this.setState({
                            isAuth: true
                        })
                    }
                }.bind(this)
            )
        }
    }

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

