import axios from 'axios';
import store from '../store';
import { getArticlesSuccess } from '../actions/articles-actions';
import {getProfileSuccess} from "../actions/user-actions";
import qs from 'qs';


//TODO think about correction urls
export function getPublications() {
  return axios.get('http://c-c.ru/api/getPublications')
    .then(response => {
      //store.dispatch(getArticlesSuccess(response.data));
      return response;
    })
      .catch(error => {
          return error;
      });;
}

export function getRubrics() {
    return axios.get('http://c-c.ru/api/rubrics')
        .then(response => {
            //store.dispatch(getArticlesSuccess(response.data));

            response.data.map(a => {
                //TODO is len returned rubrics is empty
                a['value'] = a['name']
                a['label'] = a['name']
            })
            return response;
        })
        .catch(error => {
            console.log(error)
            return error;
        });
}

export function savePublication(publication) {
    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        data: JSON.stringify(publication),
        url: 'http://c-c.ru/api/savePublication',
       // withCredentials: true
    };
    return axios(options)
        .then(res => {
            //console.log(res)
            //store.dispatch(getProfileSuccess(res.data));
            return res;
        })
        .catch(error => {
            return error;
        });
}

