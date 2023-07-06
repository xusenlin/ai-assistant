package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golangunit/localtime"
	"go-admin/global"
	"net/http"
	"time"
)

type OpenaiKey struct {
	BaseModel
	Value           string  `binding:"required,min=20,max=88"`
	ExceptionReason string  `gorm:"type:varchar(1000);comment:异常原因"`
	TotalAmount     float64 `gorm:"default:0.00"`
	UsedAmount      float64 `gorm:"default:0.00"`
	ExpireTime      *localtime.LocalTime
	Status          int `gorm:"default:1"`
}

func (o *OpenaiKey) Destroy(id string) (err error) {
	return global.DB.Where("id = ?", id).Delete(&o).Error
}

func (o *OpenaiKey) Save() (err error) {
	return global.DB.Save(&o).Error
}

func (o *OpenaiKey) UpdateAmount() error {
	if o.Value == "" {
		return errors.New("key 无效")
	}

	type SubscriptionData struct {
		AccessUntil      int     `json:"access_until"`
		HardLimitUsd     float64 `json:"hard_limit_usd"`
		HasPaymentMethod bool    `json:"has_payment_method"`
	}

	type UsageData struct {
		TotalUsage float64 `json:"total_usage"`
	}
	// 计算起始日期和结束日期
	now := time.Now()
	startDate := now.Add(-90 * 24 * time.Hour)
	endDate := now.Add(24 * time.Hour)

	// 设置API请求URL和请求头
	urlSubscription := "https://api.openai.com/v1/dashboard/billing/subscription"                                                                      // 查是否订阅
	urlUsage := fmt.Sprintf("https://api.openai.com/v1/dashboard/billing/usage?start_date=%s&end_date=%s", formatDate(startDate), formatDate(endDate)) // 查使用量
	headers := map[string]string{
		"Authorization": "Bearer " + o.Value,
		"Content-Type":  "application/json",
	}

	// 创建HTTP客户端
	client := &http.Client{}

	// 获取API限额
	req, err := http.NewRequest("GET", urlSubscription, nil)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("openai响应非200，可能你的账户已被封禁，请登录OpenAI进行查看。")
	}

	var subscriptionData SubscriptionData
	err = json.NewDecoder(resp.Body).Decode(&subscriptionData)
	if err != nil {
		return err
	}

	// 判断是否过期
	timestampNow := time.Now().Unix()
	timestampExpire := int64(subscriptionData.AccessUntil)
	if timestampNow > timestampExpire {
		return fmt.Errorf("您的账户额度已过期, 请登录OpenAI进行查看。")
	}

	totalAmount := subscriptionData.HardLimitUsd
	isSubscribed := subscriptionData.HasPaymentMethod

	// 获取已使用量
	req, err = http.NewRequest("GET", urlUsage, nil)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var usageData UsageData
	err = json.NewDecoder(resp.Body).Decode(&usageData)
	if err != nil {
		return err
	}

	totalUsage := usageData.TotalUsage / 100

	// 如果用户绑卡，额度每月会刷新
	if isSubscribed {
		// 获取当前月的第一天日期
		day := now.Day()                                                                                                                                  // 本月过去的天数
		startDate = now.Add(-time.Hour * 24 * time.Duration(day-1))                                                                                       // 本月第一天
		urlUsage = fmt.Sprintf("https://api.openai.com/v1/dashboard/billing/usage?start_date=%s&end_date=%s", formatDate(startDate), formatDate(endDate)) // 查使用量
		req, err = http.NewRequest("GET", urlUsage, nil)
		if err != nil {
			return err
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&usageData)
		if err != nil {
			return err
		}

		totalUsage = usageData.TotalUsage / 100
	}
	//timestampExpire
	o.ExpireTime = &localtime.LocalTime{Time: time.Unix(timestampExpire, 0)}
	o.TotalAmount = totalAmount
	o.UsedAmount = totalUsage

	return global.DB.Save(&o).Error
}
func formatDate(date time.Time) string {
	year, month, day := date.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
