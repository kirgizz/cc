import axios from 'axios';
import store from '../store';
import { getArticlesSuccess } from '../actions/articles-actions';

/**
 * Get all users
 */

export function getArticles() {
  return axios.get('http://localhost:8080/api/getArticles')
    .then(response => {
      store.dispatch(getArticlesSuccess(response.data));
      return response;
    });
}
