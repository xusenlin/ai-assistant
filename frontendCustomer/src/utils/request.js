import {getUserToken} from "../stores/user.js"
import {toFormData} from "./app.js"
import axios from "axios";
import { ElNotification,ElLoading }  from "element-plus"
import {baseURL, timeout} from "../config/req.js"


let loadingInstance = null
export class Request {
    instance;

    constructor(config) {
        this.instance = axios.create(config);
        this.instance.interceptors.request.use(
            (config) => {
                if ("token" in config) {
                    config.headers.Authorization = config.token
                } else {
                    config.headers.Authorization = getUserToken()
                }
                if (config.isFormRequest) {
                    config.transformRequest = toFormData
                }
                if (!config.closeLoading) {
                    loadingInstance = ElLoading.service({ fullscreen: true })
                }
                return config;
            },
            (err) => {
                // 请求错误，这里可以用全局提示框进行提示
                return Promise.reject(err);
            }
        );

        this.instance.interceptors.response.use(
            (res) => {
                // 直接返回res，当然你也可以只返回res.data
                // 系统如果有自定义code也可以在这里处理
                if (loadingInstance){
                    loadingInstance.close()
                }
                if(res.config.closeInstance){
                    return Promise.resolve(res);
                }
                let response = res.data
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
            (err) => {
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
    unhandledRequest(config) {
        return this.instance.request({...config, closeInstance: true});
    }

    //做了拦截处理，自动报错，只返回关心的数据
    request(config) {
        return this.instance.request(config);
    }
}

export default new Request({baseURL, timeout})



