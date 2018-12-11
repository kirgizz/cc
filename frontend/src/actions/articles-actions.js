import * as types from '../actions/actions-types';

export function getArticlesSuccess(articles) {
  return {
    type: types.GET_ARTICLES_SUCCESS,
    articles
  };
}
