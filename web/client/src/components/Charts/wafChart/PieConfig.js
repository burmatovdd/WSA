import * as httpClient from "../../../httpClient.js";

 export async function getData() {
    let sendUrl = "http://localhost:8080/api/stats";
    let postInfo = httpClient.Get(sendUrl)

     let withWaf, noWaf

     postInfo.then(response => {
        let resp = JSON.parse(response.data.body)
        withWaf = resp.withWaf
        noWaf = resp.allServers - withWaf
    })
     await new Promise((resolve, reject) => setTimeout(resolve, 1000));
    return {withWaf,noWaf}
}
