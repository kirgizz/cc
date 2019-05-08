import React, { Component } from "react";
import {logout} from "../../api/auth-api";
import Cookies from 'js-cookie';

import router from '../../router';
import Article from "./articles-container";



class LogoutContainer extends Component {
    componentWillMount() {
        var ssid = Cookies.get("ssid")
        logout({"ssid": ssid})

    }

    render() {
        return (
            <div/>
        );
    }
}

export default LogoutContainer

