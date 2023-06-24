import r from "@/utils/request.ts"

export type OpenaiKey = {
    "ID": number,
    "CreatedAt": string,
    "UpdatedAt": string,
    Value: string,
    isCardBound: boolean,
    ExpirationTime: number,
    Status: number,
}


export function getOpenaiKey() {
    return r.request<OpenaiKey[]>({
        url: '/v1/openai/keys',
        method: 'get',
    })
}

export function destroy(params:{id:number}) {
    return r.request({
        url: '/v1/openai/key/destroy',
        method: 'get',
        params,
    })
}
export function add(data:{Value:string}) {
    return r.request({
        url: '/v1/openai/key/add',
        method: 'post',
        data,
    })
}
export function ping(params:{key:string}) {
    return r.request<string>({
        url: '/v1/openai/key/ping',
        method: 'get',
        params,
    })
}
