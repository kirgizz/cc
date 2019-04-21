import * as types from '../actions/actions-types';
//import _ from 'lodash';

const initialState = {
    profile: "",
};

const profileReducer = function(state = initialState, action) {

    switch(action.type) {

        case types.GET_PROFILE_SUCCESS:
            return Object.assign({}, state, { profile: action.profile });


        default:

    }

    return state;

}

export default profileReducer;
