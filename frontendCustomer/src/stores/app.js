import {defineStore} from "pinia"
import storage from "good-storage"
import {ElMessage } from "element-plus"
import {storageKeyDialog, storageKeyToken} from "../config/app.js";
import {v4 as uuid} from "uuid"


export const useAppStore = defineStore("app", {
  state:() =>{
    return {
      showLogin: !storage.get(storageKeyToken,false),
      dialogIndex:0,
      dialog:storage.get(storageKeyDialog,[{title:"默认对话",id:uuid(),content:[]}])
    }
  },
  actions: {
    addDialog(){
      this.dialog.push({
        id:uuid(),
        title:`新的对话`,
        content:[]
      })
      this.dialogIndex = this.dialog.length-1
      storage.set(storageKeyDialog,this.dialog)
    },
    delDialog(index){
      if(this.dialog.length===1){
        ElMessage.warning("最少保留一个对话")
        return
      }

      if(index<=this.dialogIndex){
        this.dialogIndex--
      }
      this.dialog.splice(index,1)
      storage.set(storageKeyDialog,this.dialog)
    },
    resetDialog(index){
      this.dialog[index].content = []
      storage.set(storageKeyDialog,this.dialog)
    },
    pushContent(id,dialog=[]){//响应和初始化不能使用index
      for (let i =0;i<this.dialog.length;i++){
        let r = this.dialog[i]
        if(r.id ===id){
          if(this.dialog[i].content.length === 0){
            let q = dialog[0].content
            this.dialog[i].title = q.substring(0,q.length>=15?15:q.length)
          }
          this.dialog[i].content.push(...dialog)
          storage.set(storageKeyDialog,this.dialog)
          return
        }
      }
    },
    pushSetAnswer(id,val){
      for (let i =0;i<this.dialog.length;i++){
        let r = this.dialog[i]
        if(r.id ===id){
          let len = this.dialog[i].content.length
          let c = this.dialog[i].content[len-1]
          this.dialog[i].content[len-1].content = c.content + val
          storage.set(storageKeyDialog,this.dialog)
          return
        }
      }
    }
  }
})
