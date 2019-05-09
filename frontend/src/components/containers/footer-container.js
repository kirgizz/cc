import React, { Component } from "react";
import {connect} from "react-redux";

class FooterContainer extends Component {
    render() {
        return (
            <div><p>This is footer</p>
            </div>
        );
    }
}

const mapStateToProps = function(store) {
    return {
        profile: store.profileState.profile
    };
};

export default connect(mapStateToProps)(FooterContainer);

