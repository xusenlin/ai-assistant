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

export function userList(data:any)  {
  return r.request<User[]>({
    url: '/v1/users/list',
    method: 'get',
    data,
  })
}

export function destroy(params:{id:number})  {
  return r.request({
    url: '/v1/user/destroy',
    method: 'get',
    params
  })
}
