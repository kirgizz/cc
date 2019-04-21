import axios from 'axios';
import store from '../store';
import { getArticlesSuccess } from '../actions/articles-actions';

/**
 * Get all users
 */

export function getArticles() {
  return axios.get('http://c-c.ru/api/getArticles')
    .then(response => {
      store.dispatch(getArticlesSuccess(response.data));
      //return response;
    });
}
