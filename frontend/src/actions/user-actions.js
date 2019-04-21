import * as types from '../actions/actions-types';

export function getProfileSuccess(profile) {
    return {
        type: types.GET_PROFILE_SUCCESS,
        profile
    };
}
