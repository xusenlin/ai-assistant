import {defineStore} from "pinia"
import storage from "good-storage"
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
  },
  actions: {
    updateUserInfo(user) {
      this.info = user
      storage.set(storageKeyUser, user)
      storage.set(storageKeyToken,user.Token)
    },
    decrementDialogueCount(){
      this.info.RemainingDialogueCount = this.info.RemainingDialogueCount - 1
      storage.set(storageKeyUser, this.info)
    },
    loginOut(){
      storage.remove(storageKeyUser)
      storage.remove(storageKeyToken)
      window.location.reload()
    }
  }
})

export const getUserInfo = ()=> {
  return storage.get(storageKeyUser, {})
}


export const getUserToken = ()=> {
  return storage.get(storageKeyToken, {})
}
