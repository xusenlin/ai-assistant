import Axios from "axios";
import { getToken,getCurrentUrl,removeToken } from "@/utils/common.js";
import { showLoadingToast,closeToast,showToast } from "vant";
import { appid } from "@/config/app"
import { baseURL,timeout,requestRetryDelay,requestRetry } from "@/config/request";

const service = Axios.create({
  baseURL,
  headers: {
    Accept: "*/*"
  },
  timeout
});

service.defaults.retry = requestRetry;
service.defaults.retryDelay = requestRetryDelay;


service.interceptors.request.use(
  config => {
    if (config.showLoading) {
      showLoadingToast({
        message: '加载中...',
        forbidClick: true,
      });
    }
    config.headers["Authorization"] = getToken();
    return config;
  },
  error => {
    closeToast()
    Promise.reject(error);
  }
);

service.interceptors.response.use(
  res => {
    closeToast()
    if (res.status !== 200) {
      showToast('数据返回出错');
      return Promise.reject("响应非200！");
    } else {
      if(res.data.code === "401"){
        showLoadingToast({
          message: '登录失效，正在重新登录',
          forbidClick: true,
        });
        removeToken()
        setTimeout(()=>{
          window.location.href = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=${appid}&redirect_uri=${getCurrentUrl()}&response_type=code&scope=snsapi_base&state=123#wechat_redirect`
        },600)
        return
      }
      if(res.data.code === "A0215"){//code失效
        showLoadingToast({
          message: '微信授权失败，正在重新授权登录',
          forbidClick: true,
        });
        removeToken()
        setTimeout(()=>{
          window.location.href = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=${appid}&redirect_uri=${getCurrentUrl()}&response_type=code&scope=snsapi_base&state=123#wechat_redirect`
        },600)
        return
      }
      if (res.data.code != '00000') {
        //统一处理错误
        showToast(res.data.msg);
        return Promise.reject("error");
      }
      return res.data.data;
    }
  },
  error => {
    showToast(error.message);
    return Promise.reject(error);
  }
);

export default service;


export const transformRequest = [function (data) {
  let formData = new FormData()
  for (let key in data) {
    formData.append(key, data[key])
  }
  return formData
}];
