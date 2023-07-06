package main

import (
	"bufio"
	"fmt"
	"go-admin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

//2023-07-04 10:08:49之前
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("此操作会将Dialog表创建时间在2023-07-04 10:08:49之前的数据的创建时间全部添加8个小时，是否同意？(Y/N): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if strings.EqualFold(answer, "Y") {
		RunFixTime()
	} else {
		fmt.Println("您不同意。已操作取消。")
	}
}

func RunFixTime() {
	DB, err := gorm.Open(sqlite.Open("../record.db"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	var records []models.Dialog

	targetTime := time.Date(2023, 7, 4, 10, 8, 49, 0, time.UTC)
	if err := DB.Where("created_at < ?", targetTime).Find(&records).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("总共查询到符合时间的数据%v条，正在调整时间", len(records)))
	for _, r := range records {
		r.CreatedAt.Time = r.CreatedAt.Time.Add(8 * time.Hour)
		if err := DB.Save(r).Error; err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("ID=%v的数据已更新完成", r.ID))
	}
}
