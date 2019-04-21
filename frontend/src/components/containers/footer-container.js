import React, { Component } from "react";
import Footer from '../views/footer';
import {connect} from "react-redux";

//import store from "../../store"
//import profileReducer from "../../reducers"

class FooterContainer extends Component {
    constructor(props) {
        super(props);

    }

    render() {
        //console.log("render", this.props)
        return (

            <div></div>
            //<Footer {...this}/>
        );
    }
}


const mapStateToProps = function(store) {
    return {
        profile: store.profileState.profile
    };
};

export default connect(mapStateToProps)(FooterContainer);

