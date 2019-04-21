import axios from 'axios';
import qs from 'qs';
import store from "../store";
import {getProfileSuccess} from "../actions/user-actions";



export function checkCredentials(credentials) {
    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        //why is qs??
        data: qs.stringify(credentials),
        url: 'http://c-c.ru/api/login',
        withCredentials: true
    };
    return axios(options)
        .then(res => {
            //console.log(res)
            store.dispatch(getProfileSuccess(res.data));
            return res;
        })
        .catch(error => {
            return error;
        });
}

export function checkCookie(cookie) {
    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        data: cookie,
        url: 'http://c-c.ru/api/checkSession',
    };
    return axios(options)
        .then(res => {
            return res;
        })
        .catch(error => {
            return error;
        });
}
