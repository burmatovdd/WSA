import *as storage from "./storage";
import axios from "axios";

function request(method,url,data){

   return axios.post(url, JSON.stringify(data));
}

export function Post(url,sendData){
    return request('POST', url,sendData);
}
