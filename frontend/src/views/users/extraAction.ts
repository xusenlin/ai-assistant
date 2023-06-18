import {ElMessage, ElMessageBox} from 'element-plus'
import { destroy ,updatePassword,updateStatus,updateRemainingDialogueCount} from "@/api/user.ts"



export default function (refreshTable: () => void) {


    const deleteRow = (row: any) => {
        ElMessageBox.confirm(
            '确认删除此用户吗？',
            '确认操作',
            {
                confirmButtonText: '确认',
                cancelButtonText: '关闭',
                type: 'warning',
            }
        )
            .then(() => {
                let {ID} = row;
                destroy({id: ID}).then(() => {
                    ElMessage({
                        type: 'success',
                        message: '删除成功',
                    })
                    refreshTable()
                }).catch(() => {})
            }).catch(() => {})
    }

    const setPassword = (id:number) => {
        ElMessageBox.prompt('请输入新的密码', '提示', {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
        }).then(({value}) => {
            updatePassword({id,password:value}).then(() => {
                ElMessage({
                    type: 'success',
                    message: `设置成功`,
                })
            })
        }).catch(() => {})
    }

    const setRemainingDialogueCount = (id:number) => {
        ElMessageBox.prompt('设置用户对话次数', '提示', {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
        }).then(({value}) => {
            updateRemainingDialogueCount({id,count:value}).then(() => {
                ElMessage({
                    type: 'success',
                    message: `设置成功`,
                })
                refreshTable()
            })
        }).catch(() => {})
    }
    const setStatus = (row:any) => {
        const { Status,ID } = row


        ElMessageBox.confirm(
            `确认${Status==0?"启用":"禁用"}此用户吗？`,
            '确认操作',
            {
                confirmButtonText: '确认',
                cancelButtonText: '关闭',
                type: 'warning',
            }
        )
            .then(() => {
                updateStatus({id:ID,status:Status == 0?1:0}).then(() => {
                    ElMessage({
                        type: 'success',
                        message: '操作成功',
                    })
                    refreshTable()
                }).catch(() => {})
            }).catch(() => {})
    }

    return {
        deleteRow,setPassword,setStatus,setRemainingDialogueCount
    }

}
