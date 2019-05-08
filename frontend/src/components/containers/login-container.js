import React, { Component } from "react";
import Login from '../views/login';
import {checkCredentials} from "../../api/auth-api";
import {connect} from "react-redux";

import router from '../../router';

class LoginContainer extends Component {
    constructor(props) {
        super(props);

        this.state = {
            email: '',
            password: '',
            elemIsVisible: false,
            loginSuccess: false
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange = event => {
        this.setState({
            [event.target.id]: event.target.value
        });
    }

    validateForm = () => this.state.email.length > 0 && this.state.password.length > 0;

    showElement() {
        this.setState({
            elemIsVisible: true
        })
    }

    handleSubmit = event => {

        checkCredentials({"email":this.state.email, "password": this.state.password}).then(
            function(result) {
                if (result.status === 200) {
                    router.props.history.push("/")
                    window.location.reload()
                } else {
                    console.log()
                    this.showElement("")
                }
            }.bind(this)
        )
        event.preventDefault();
    }
    render() {
        return (
            <Login {...this}/>
        );
    }
}


const mapStateToProps = function(store) {
    return {
        profile: store.profileState.profile
    };
};

export default connect(mapStateToProps)(LoginContainer);

