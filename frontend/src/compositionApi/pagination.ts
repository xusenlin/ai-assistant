import { ref,nextTick} from 'vue'
import { resetArgs } from "@/utils/app"


export default function () {
    const tableData = ref([])
    const paginationRef = ref<{ Refresh:()=>void, QueryParams:()=>void } | null>(null);
    const searchParams = ref<any>({})

    const setTableData = (r:any) :void => {
        if(Array.isArray(r.List) && r.List.length!==0){
            tableData.value = r.List||[]
        }
    }

    const refreshTable = ():void => {
        paginationRef.value?.Refresh()
    }
    const resetParams = async () => {
        searchParams.value = resetArgs(searchParams.value)
        await nextTick()
        refreshTable()
    }

    return {
        searchParams,
        tableData,
        paginationRef,
        setTableData,
        refreshTable,
        resetParams
    }
}
