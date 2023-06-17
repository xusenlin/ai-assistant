import {ElMessage, ElMessageBox} from 'element-plus'
import {migrate, destroy, add} from "@/api/sensitiveWords.ts"


export default function (refreshTable: () => void) {


    const deleteRow = (row: any) => {
        ElMessageBox.confirm(
            '确认删除此记录吗？',
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
    const clickMigrate = () => {
        migrate().then(() => {
            ElMessage({
                type: 'success',
                message: '迁移成功',
            })
        })
    }

    const addSensitiveWords = () => {
        ElMessageBox.prompt('请输入你的敏感词', '提示', {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
        }).then(({value}) => {
            add(value).then(() => {
                ElMessage({
                    type: 'success',
                    message: `添加成功`,
                })
                refreshTable()
            })

        })
            .catch(() => {

            })
    }

    return {
        clickMigrate,
        deleteRow,
        addSensitiveWords
    }

}
