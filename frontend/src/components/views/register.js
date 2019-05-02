//import React, { Component } from "react";
import React from "react";
import { Button, FormGroup, FormControl, ControlLabel } from "react-bootstrap";
import "./styles/login.css";

function Register(props) {
    return (
    <div>
        <div className="RegisterForm">
            <form onSubmit={props.handleSubmit}>
                <FormGroup controlId="email" bsSize="large">
                    <ControlLabel>Email</ControlLabel>
                    <FormControl
                        autoFocus
                        value={props.state.email}
                        onChange={props.handleChange}
                    />
                </FormGroup>

                <FormGroup controlId="nickname" bsSize="large">
                    <ControlLabel>Nickname</ControlLabel>
                    <FormControl
                        autoFocus
                        value={props.state.nickname}
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

                <FormGroup controlId="confirmPassword" bsSize="large">
                    <ControlLabel>Confir password</ControlLabel>
                    <FormControl
                        value={props.state.confirmPassword}
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
                    Register
                </Button>
            </form>
            <div>
                {
                    props.state.elemIsVisible?
                        <div>
                            Passwords dont confirm
                        </div>
                        :null
                }
            </div>
        </div>
     </div>

    );

}

export default Register
