import  r  from "@/utils/request.ts"
import {User} from "@/stores/user.ts";



export function register(data:any)  {
  return r.request({
    url: '/v1/register',
    method: 'post',
    data,
  })
}

export function login(data:any)  {
  return r.request<User>({
    url: '/v1/login',
    method: 'post',
    data,
  })
}

export function userList(params:any)  {
  return r.request<User[]>({
    url: '/v1/users/list',
    method: 'get',
    params,
  })
}

export function destroy(params:{id:number})  {
  return r.request({
    url: '/v1/user/destroy',
    method: 'get',
    params
  })
}


//
export function updateStatus(params:{id:number,status:number})  {
  return r.request({
    url: '/v1/user/updateStatus',
    method: 'get',
    params
  })
}
export function updatePassword(params:{id:number,password:string})  {
  return r.request({
    url: '/v1/user/updatePassword',
    method: 'get',
    params
  })
}
export function updateRemainingDialogueCount(params:{id:number,count:number})  {
  return r.request({
    url: '/v1/user/updateRemainingDialogueCount',
    method: 'get',
    params
  })
}


export function batchAddUser(data:any)  {
  return r.request({
    url: '/v1/user/batchAddUser',
    method: 'post',
    data,
  })
}