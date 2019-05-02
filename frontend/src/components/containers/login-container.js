import React, { Component } from "react";
import Login from '../views/login';
import {checkCredentials} from "../../api/auth-api";
import {connect} from "react-redux";

import router from '../../router';

//import store from "../../store"
//import profileReducer from "../../reducers"

class LoginContainer extends Component {
    constructor(props) {
        super(props);

        this.state = {
            email: '',
            password: '',
            elemIsVisible: false,
            loginSuccess: true
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.showElement = this.showElement.bind(this);

    }

    handleChange = event => {
        this.setState({
            [event.target.id]: event.target.value
        });
    }


    handleSubmit(event) {
        event.preventDefault();
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
                    this.loginSuccess = true
                    router.props.history.push("/")
                    //history.push("/")
                    //history.push("/")
                } else {
                    console.log(this)
                    this.showElement()
                }
            }.bind(this)
        )
        event.preventDefault();

        //console.log(store.getState())

    }
    render() {
        //console.log("render", this.props)
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

