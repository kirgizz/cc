import * as types from '../actions/actions-types';

export function getRubricsSuccess(rubrics) {
    return {
        type: types.GET_RUBRICS_SUCCESS,
        rubrics
    };
}
