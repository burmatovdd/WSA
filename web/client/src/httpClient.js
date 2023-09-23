// import *as storage from "./storage";
import axios from "axios";
import * as storage from "./storage.js";

function postRequest(method,url,data){

    return axios.post(url, JSON.stringify(data), {headers:{
            Authorization: storage.getToken()
        }});
}

function getRequest(method, url){
    return axios.get(url)
}

export function Post(url,sendData){
    return postRequest('POST', url,sendData);
}

export function Get(url){
    return getRequest('GET', url)
}
