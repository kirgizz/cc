import * as types from '../actions/actions-types';
//import _ from 'lodash';

const initialState = {
    rubrics: "",
};

const rubricsReducer = function(state = initialState, action) {

    switch(action.type) {

        case types.GET_RUBRICS_SUCCESS:
            return Object.assign({}, state, { rubrics: action.rubrics });


        default:

    }

    return state;

}

export default rubricsReducer;
