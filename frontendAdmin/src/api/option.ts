import r from "@/utils/request.ts"

export type Option = {
    "ID": number,
    "CreatedAt": string,
    "UpdatedAt": string,
    OptionKey: string
    OptionValue: string
}

export function getOptionByKey(params: { key: string }) {
    return r.request<Option>({
        url: '/v1/option/get',
        method: 'get',
        params,
    })
}

export function setOptionByKey(params: { key: string,val:string }) {
    return r.request<Option>({
        url: '/v1/option/set',
        method: 'get',
        params,
    })
}