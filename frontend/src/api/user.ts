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
