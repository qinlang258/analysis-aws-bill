package config

import (
	"fmt"
	"regexp"
)

const (
	CsvName       = "./BillingManagementMonthReport202508.csv"
	MysqlUser     = "root"
	MysqlPassword = "ql2252528"
	MysqlHost     = "127.0.0.1"
	MysqlPort     = 3306
	MysqlDbName   = "202508"
)

func GetMonthDay() string {
	// 修改正则表达式以匹配6位数字（假设原始字符串是 ylz-aliyun-bill202501）
	re := regexp.MustCompile(`.*bill(\d{6})`)
	match := re.FindStringSubmatch(MysqlDbName)
	if len(match) < 2 {
		fmt.Println("Invalid dbName format")
		return ""
	}

	dateStr := match[1]  // 提取到 "202501"
	year := dateStr[:4]  // 年份前四位: "2025"
	month := dateStr[4:] // 月份后两位: "01"
	result := fmt.Sprintf("%s-%s", year, month)

	return result
}
