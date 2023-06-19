import request from "@/utils/request.js"


export const login = (params) => {
  return request({
    url: "/v1/login",
    method: "post",
    data: params
  });
}
