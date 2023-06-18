import request from "@/utils/request";


export function wechatSignatureApi(params){
  return request({
    url: "/api/wxuser/getConfigData",
    method: "get",
    params: params
  });
}
