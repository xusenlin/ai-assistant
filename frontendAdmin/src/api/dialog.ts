import  r  from "@/utils/request.ts"


export function dialogList(params:any)  {
  return r.request({
    url: '/v1/record/dialog/list',
    method: 'get',
    params,
  })
}
