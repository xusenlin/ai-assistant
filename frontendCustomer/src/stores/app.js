import {defineStore} from "pinia"
import storage from "good-storage"
import {storageKeyToken} from "../config/app.js";



export const useAppStore = defineStore("app", {
  state:() =>{
    return {
      showLogin: !storage.get(storageKeyToken,false)
    }
  },
  actions: {

  }
})
