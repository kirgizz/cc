import React, { Component } from "react";
import Register from '../views/register';
import {connect} from "react-redux";
import {registerUser} from "../../api/register-api";
import router from "../../router";

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
                        //this.loginSuccess = true
                        //router.props.history.push("/")
                        console.log("user registered")
                    } else {
                        console.log(this)
                        //this.showElement()
                    }
                }.bind(this))
        } else {
            this.setState({
                elemIsVisible: true
            })
        }
    }
    render() {
        //console.log("render", this.props)
        return (
            <Register {...this}/>
        );
    }
}

const mapStateToProps = function(store) {
    return {
        profile: store.profileState.profile
    };
};

export default connect(mapStateToProps)(RegisterContainer);

