let storage = localStorage;
export function get(key){
    return storage[key]
}

export function getToken(){
    return get("token") || "";
}

export function set(key,value){
    storage[key] = value;
}
export function remove(key){
    storage.removeItem(key)
}
