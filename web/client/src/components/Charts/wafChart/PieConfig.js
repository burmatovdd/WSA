import * as httpClient from "../../../httpClient.js";

 export async function getData() {
    let sendUrl = "http://localhost:8080/api/stats";

    return  httpClient.Get(sendUrl).then( response => {
        let withWaf, noWaf
        let resp = JSON.parse(response.data.body)
        withWaf = resp.withWaf
        noWaf = resp.allServers - withWaf
        return {withWaf,noWaf}
    })
}
