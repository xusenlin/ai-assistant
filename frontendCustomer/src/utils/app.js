

/**
 * 下载或者保存一个Blob
 * @param blob
 * @param fileName
 * @param isOpen
 * 接口返回数据流时，如果是pdf可以设置isOpen直接新窗口打开
 * export function exportReport(params: { fileCode:string }) {
 *   return request({
 *     responseType:"blob",
 *     closeResponseInterceptors:true,
 *     url: '/customer-service/open/api/report/getReport',
 *     method: 'get',
 *     params
 *   })
 * }
 */
export function saveBlob(blob,fileName,isOpen = false){

  let url = window.URL.createObjectURL(blob);
  if(isOpen){
    window.open(url)
  }else {
    let a = document.createElement("a");
    document.body.appendChild(a);
    a.setAttribute("display","none")
    a.href = url;
    a.download = fileName;
    a.click();
    a.remove();
    window.URL.revokeObjectURL(url);
  }
}


export const toFormData = (data) =>{
  const formData = new FormData()
  for (const key in data) {
    formData.append(key, data[key])
  }
  return formData
}
