import  r  from "@/utils/request.ts"


export function sensitiveWordsList(params:any)  {
  return r.request({
    url: '/v1/sensitiveWords/list',
    method: 'get',
    params,
  })
}
export function migrate()  {
  return r.request({
    url: '/v1/sensitiveWords/migrate',
    method: 'get',
  })
}
export function destroy(params:{id:number})  {
  return r.request({
    url: '/v1/sensitiveWords/destroy',
    method: 'get',
    params
  })
}
export function add(word:string)  {
  return r.request({
    url: '/v1/sensitiveWords/add',
    method: 'get',
    params:{word}
  })
}
