import request from "@/utils/request.js"
import { getUserInfo,getToken } from "@/utils/common.js"
import { baseURL } from "@/config/request.js"

export const QueryReport = (params) => {
  return request({
    url: "/api/report/userQueryReport",
    method: "post",
    data: params
  });
}
export const getSamplesByUser = (params) => {
  let userId = getUserInfo("loginId")
  return request({
    url: `/api/sample/getSamplesByUser/${userId}`,
    method: "get",
    params: params
  });
}

export const bindOrUnBindSampleUser = (params) => {
  let userId = getUserInfo("loginId")
  return request({
    url: `/api/sample/bindOrUnBindSampleUser`,
    method: "post",
    data: {userId,...params}
  });
}




export const DownloadReport = (url) => {
  return  `${baseURL}/api/report/download?fileName=${url}&Authorization=${getToken()}`
}

export const getPdfViewUrl = (params) => {
  return request({
    url: `/api/report/getUrl`,
    method: "get",
    params: params
  });
}


export const sendTelMsg = (params) => {
  return request({
    url: "/api/wxuser/sendTelMsg",
    method: "post",
    data: params
  });
}

export const wxUserLogin = (params) => {
  return request({
    url: "/api/wxuser/wxUserLogin",
    method: "post",
    data: params
  });
}
