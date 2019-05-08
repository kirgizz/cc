//import React, { Component } from "react";
import React from "react";
import { Button, FormGroup, FormControl, ControlLabel } from "react-bootstrap";
import "./styles/login.css";

function Login(props) {
    return (
    <div>
        <div className="authForm">
            <form onSubmit={props.handleSubmit}>
                <FormGroup controlId="email" bsSize="large">
                    <ControlLabel>Email</ControlLabel>
                    <FormControl
                        autoFocus
                        value={props.state.email}
                        onChange={props.handleChange}
                    />
                </FormGroup>
                <FormGroup controlId="password" bsSize="large">
                    <ControlLabel>Password</ControlLabel>
                    <FormControl
                        value={props.state.password}
                        onChange={props.handleChange}
                        type="password"
                    />
                </FormGroup>
                <Button
                    block
                    bsSize="large"
                    disabled={!props.validateForm()}
                    type="submit"
                >
                    Login
                </Button>
            </form>
        </div>

        <div>
            {
                props.state.elemIsVisible?
                    <div>
                        wrong credentials
                    </div>
                    :null
            }
        </div>
     </div>

    );

}

export default Login
