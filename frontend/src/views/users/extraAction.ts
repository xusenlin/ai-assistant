import {ElMessage, ElMessageBox} from 'element-plus'
import { destroy } from "@/api/user.ts"


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
                }).catch(() => {
                })

            })
            .catch(() => {
            })
    }

    return {
        deleteRow,
    }

}
