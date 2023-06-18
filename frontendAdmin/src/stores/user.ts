import {defineStore} from "pinia"
import storage from "good-storage"
import {userInfoKey,tokenKey} from "@/config/app"



export type User = {
  "ID": number,
  "CreatedAt": string,
  "UpdatedAt": string,
  "Username":string,
  "IsAdmin": boolean,
  "TokenConsumed": number,
  "RemainingDialogueCount": number,
  "Token": string
}//自己完善

export interface UserStore {
  info: User,
//权限之类的
}



export const useUserStore = defineStore("user", {
  state: (): UserStore => {
    return {
      info: storage.get(userInfoKey),
    }
  },
  getters: {
    token: (state) => state.info.Token,
    userName: (state) => state.info.Username,
  },
  actions: {
    updateUserInfo(user: User) {
      this.info = user
      storage.set(userInfoKey, user)
      storage.set(tokenKey,user.Token)
    },
    loginOut(){
      storage.remove(userInfoKey)
      window.location.reload()
    }
  }
})

//下面这些函数是从本地读取，在pinia还没有创建的时候，例如路由钩子
export const getUserInfo = (): User => {
  return storage.get(userInfoKey, {})
}

export const getUserInfoByKey = (k: keyof User): any => {
  let u = getUserInfo()
  return u[k]
}

export const getToken = (): string => {
  let u = getUserInfo()
  return u["Token"]
}
