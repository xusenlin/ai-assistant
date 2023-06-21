import { isProduction } from "./app.js"


export const  baseURL = isProduction ? window.location.origin:import.meta.env.VITE_BASE_URL

