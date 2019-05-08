import axios from 'axios';
import store from "../store";
import {getProfileSuccess} from "../actions/user-actions";
import Cookies from 'js-cookie';



export function checkCredentials(credentials) {
    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        //why is qs??
        data: JSON.stringify(credentials),
        url: 'http://c-c.ru/api/login',
        withCredentials: true
    };
    return axios(options)
        .then(res => {
            store.dispatch(getProfileSuccess(res.data));
            console.log(Cookies.get("ssid"))
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
        data: JSON.stringify(cookie),
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


export function logout(cookie) {
    const options = {
        method: 'POST',
        url: 'http://c-c.ru/api/logout',
        data: JSON.stringify(cookie),
        // withCredentials: true
    };
    return axios(options)
        .then(res => {
            //store.dispatch(getProfileSuccess(res.data));
            return res;
        })
        .catch(error => {
            return error;
        });
}
