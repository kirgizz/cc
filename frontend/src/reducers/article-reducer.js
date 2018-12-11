import * as types from '../actions/actions-types';
import _ from 'lodash';

const initialState = {
  articles: [],
};

const articleReducer = function(state = initialState, action) {

  switch(action.type) {

    case types.GET_ARTICLES_SUCCESS:
      return Object.assign({}, state, { articles: action.articles });

  }

  return state;

}

export default articleReducer;
