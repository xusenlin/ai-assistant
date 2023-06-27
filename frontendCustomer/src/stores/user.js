import {defineStore} from "pinia"
import storage from "good-storage"
import {getMineInfo} from "../api/user.js"
import {storageKeyToken, storageKeyUser} from "../config/app.js"


export const useUserStore = defineStore("user", {
  state:() =>{
    return {
      info: storage.get(storageKeyUser,null),
    }
  },
  getters: {
    token: (state) => state.info.Token,
    userName: (state) => state.info.Username,
    remainingDialogueCount: (state) => state.info.RemainingDialogueCount,
    tokenConsumed: (state) => state.info.TokenConsumed,
  },
  actions: {
    updateUserInfo(user) {
      this.info = user
      storage.set(storageKeyUser, user)
      storage.set(storageKeyToken,user.Token)
    },
    loginOut(){
      storage.remove(storageKeyUser)
      storage.remove(storageKeyToken)
      window.location.reload()
    },
    async updateUserInfoByApi(){
      try {
        this.info = await getMineInfo()
        storage.set(storageKeyUser, this.info)
      } catch (error) {
        return error
      }
    },
  }
})

export const getUserInfo = ()=> {
  return storage.get(storageKeyUser, {})
}


export const getUserToken = ()=> {
  return storage.get(storageKeyToken, "")
}
