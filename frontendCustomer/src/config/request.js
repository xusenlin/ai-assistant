import {isProEnv } from "./env"
export const timeout = 3000

export const requestRetry = 4

export const requestRetryDelay = 800

export const baseURL = isProEnv ? window.location.origin:process.env.VUE_APP_REQ_URL


