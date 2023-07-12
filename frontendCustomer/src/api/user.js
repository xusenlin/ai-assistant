import r from "../utils/request.js"


export function getMineInfo()  {
  return r.request({
    closeLoading:true,
    url: '/api/user/getUserInfo',
    method: 'get',
  })
}


export function updatePwd(password)  {
  return r.request({
    url: '/api/user/updatePwd',
    method: 'post',
    data:{password}
  })
}
