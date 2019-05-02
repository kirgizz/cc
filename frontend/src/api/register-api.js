import axios from 'axios';
import qs from 'qs';
import store from "../store";
import {getProfileSuccess} from "../actions/user-actions";



export function registerUser(credentials) {
    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        //why is qs??
        data: qs.stringify(credentials),
        url: 'http://c-c.ru/api/register',
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
