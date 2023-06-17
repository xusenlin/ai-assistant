import {getToken} from "@/utils/user.ts"
import {toFormData} from "@/utils/app.ts"
import axios from "axios";
import { ElNotification,ElLoading }  from "element-plus"
import {baseURL, timeout} from "@/config/request.ts"
import type {AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig} from "axios";

export type ResponseResult<T> = {
    Status: boolean;
    Msg: string;
    Data: T;
};

declare module 'axios' {
    export interface AxiosRequestConfig {
        closeLoading?: boolean,//默认所有请求Loading，可关闭
        token?: string,//默认获取本地token，可针对某个请求写死或置空
        isFormRequest?: boolean,//将请求自动转换为表单请求
        closeInstance?: boolean
    }
}

export type RequestConfig = Omit<AxiosRequestConfig, 'closeInstance' | 'transformRequest'>

let loadingInstance:any
export class Request {
    instance: AxiosInstance;

    constructor(config: AxiosRequestConfig) {
        this.instance = axios.create(config);
        this.instance.interceptors.request.use(
            (config: InternalAxiosRequestConfig) => {
                if ("token" in config) {
                    config.headers.Authorization = config.token
                } else {
                    config.headers.Authorization = getToken()
                }
                if (config.isFormRequest) {
                    config.transformRequest = toFormData
                }
                if (!config.closeLoading) {
                    loadingInstance = ElLoading.service({ fullscreen: true })
                }
                return config;
            },
            (err: any) => {
                // 请求错误，这里可以用全局提示框进行提示
                return Promise.reject(err);
            }
        );

        this.instance.interceptors.response.use(
            (res: AxiosResponse) => {
                // 直接返回res，当然你也可以只返回res.data
                // 系统如果有自定义code也可以在这里处理
                if (loadingInstance){
                    loadingInstance.close()
                }
                if(res.config.closeInstance){
                    return Promise.resolve(res);
                }
                let response = res.data as ResponseResult<any>
                if (!response.Status){
                    ElNotification({
                        showClose: true,
                        message: response.Msg,
                        type: "error",
                    });
                    return Promise.reject(response.Msg);
                }
                return Promise.resolve(response.Data);
            },
            (err: any) => {
                if (loadingInstance){
                    loadingInstance.close()
                }
                ElNotification({
                  showClose: true,
                  message: `${err.message}，请检查网络或联系管理员！`,
                  type: "error",
                });
                return Promise.reject(err.response);
            }
        );
    }

    //未拦截请求，响应原封不动返回
    unhandledRequest<T>(config: RequestConfig): Promise<AxiosResponse<ResponseResult<T>>> {
        return this.instance.request({...config, closeInstance: true});
    }

    //做了拦截处理，自动报错，只返回关心的数据
    request<T>(config: RequestConfig): Promise<T> {
        return this.instance.request(config);
    }
}

export default new Request({baseURL, timeout})



