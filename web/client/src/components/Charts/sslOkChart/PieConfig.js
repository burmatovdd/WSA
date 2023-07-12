import * as httpClient from "../../../httpClient.js";

export async function getData() {
    let sendUrl = "http://localhost:8080/api/stats";
    let postInfo = httpClient.Get(sendUrl)

    let okCertificates, noOkCertificates

    postInfo.then(response => {
        let resp = JSON.parse(response.data.body)
        okCertificates =  resp.okCertificates
        noOkCertificates = resp.allCertificates - okCertificates
    })
    await new Promise((resolve, reject) => setTimeout(resolve, 1000));
    return {okCertificates,noOkCertificates}
}
