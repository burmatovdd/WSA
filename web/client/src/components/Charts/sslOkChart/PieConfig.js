import * as httpClient from "../../../httpClient.js";

export async function getData() {
    let sendUrl = "http://localhost:8080/api/statistic";
    let postInfo = httpClient.Get(sendUrl)

    return httpClient.Get(sendUrl).then(response =>{
        let okCertificates, noOkCertificates
        let resp = JSON.parse(response.data.body)
        okCertificates =  resp.okCertificates
        noOkCertificates = resp.allCertificates - okCertificates
        return {okCertificates,noOkCertificates}
    })
}
