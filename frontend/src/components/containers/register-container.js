import React, { Component } from "react";
import Register from '../views/register';
import {connect} from "react-redux";
import {registerUser} from "../../api/register-api";
import router from "../../router";
import AddContainer from "./add-container";


class RegisterContainer extends Component {
    constructor(props) {
        super(props);

        this.state = {
            email: '',
            password: '',
            confirmPassword: '',
            nickname:'',
            elemIsVisible: false,
            loginSuccess: true
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange = event => {
        this.setState({
            [event.target.id]: event.target.value
        });
    }

    handleSubmit(event) {
        event.preventDefault();
    }
    validatePassword = () => this.state.password == this.state.confirmPassword;

    validateForm = () => this.state.email.length > 0 && this.state.password.length > 0 && this.state.nickname.length > 0;

    handleSubmit = event => {

        event.preventDefault();
        if (this.validatePassword()) {

            registerUser({"email":this.state.email, "password": this.state.password}).then(
                function(result) {
                    if (result.status === 200) {
                        console.log("user registered")
                    } else {
                        console.log(this)
                    }
                }.bind(this))
        } else {
            this.setState({
                elemIsVisible: true
            })
        }
    }
    render() {
        return (
            <Register {...this}/>
        );
    }
}


export default RegisterContainer
//const mapStateToProps = function(store) {
//    return {
//        profile: store.profileState.profile
//    };
//};

//export default connect(mapStateToProps)(RegisterContainer);

