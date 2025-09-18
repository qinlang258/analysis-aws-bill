package main

import (
	"encoding/json"
	"fmt"
	"log"

	"analysis-aws-bill/config"
)

func main() {
	// 创建EC2价格表
	pricingTable := config.NewEC2PricingTable()

	// 从CSV文件加载按需价格
	csvFilePath := "./BillingManagementMonthReport202508.csv"
	err := pricingTable.LoadOnDemandPricesFromCSV(csvFilePath)
	if err != nil {
		log.Fatalf("加载CSV文件失败: %v", err)
	}

	// 加载Savings Plans和Reserved Instance价格
	pricingTable.LoadSavingsPlansAndRIPrices()

	// 打印价格表
	fmt.Println("=== EC2实例价格对比表 ===")
	pricingTable.PrintTable()

	// 获取特定实例的节省对比
	fmt.Println("\n=== 节省对比分析 ===")

	// 示例1: 香港地区的t3.micro实例
	comparison1 := pricingTable.GetSavingsComparison("605134445666", "Asia Pacific (Hong Kong)", "t3.micro")
	if comparison1 != nil {
		fmt.Println("\n香港地区 t3.micro 实例价格对比:")
		printSavingsComparison(comparison1)
	}

	// 示例2: 新加坡地区的c5.large实例
	comparison2 := pricingTable.GetSavingsComparison("017820696133", "Asia Pacific (Singapore)", "c5.large")
	if comparison2 != nil {
		fmt.Println("\n新加坡地区 c5.large 实例价格对比:")
		printSavingsComparison(comparison2)
	}

	// 生成JSON报告
	fmt.Println("\n=== 生成JSON价格报告 ===")
	generateJSONReport(pricingTable)
}

// printSavingsComparison 打印节省对比信息
func printSavingsComparison(comparison map[string]interface{}) {
	instanceInfo := comparison["instance_info"].(map[string]string)
	pricing := comparison["pricing"].(map[string]float64)
	savings := comparison["savings"].(map[string]interface{})

	fmt.Printf("实例信息: %s - %s - %s\n",
		instanceInfo["account_id"],
		instanceInfo["region"],
		instanceInfo["instance_type"])

	fmt.Printf("价格对比:\n")
	fmt.Printf("  按需实例 (每月):     $%.2f\n", pricing["on_demand_monthly"])
	fmt.Printf("  SP 3年无预付 (每月): $%.2f\n", pricing["sp_3y_no_upfront"])
	fmt.Printf("  SP 3年全预付 (每月): $%.2f\n", pricing["sp_3y_all_upfront"])
	fmt.Printf("  RI 3年无预付 (每月): $%.2f\n", pricing["ri_3y_no_upfront"])
	fmt.Printf("  RI 3年全预付 (每月): $%.2f\n", pricing["ri_3y_all_upfront"])

	fmt.Printf("节省分析:\n")
	for savingType, savingData := range savings {
		if data, ok := savingData.(map[string]float64); ok {
			fmt.Printf("  %s: 节省 $%.2f (%.1f%%)\n",
				getSavingTypeName(savingType),
				data["amount"],
				data["percent"])
		}
	}
}

// getSavingTypeName 获取节省类型的中文名称
func getSavingTypeName(savingType string) string {
	switch savingType {
	case "sp_3y_no_upfront":
		return "SP 3年无预付"
	case "sp_3y_all_upfront":
		return "SP 3年全预付"
	case "ri_3y_no_upfront":
		return "RI 3年无预付"
	case "ri_3y_all_upfront":
		return "RI 3年全预付"
	default:
		return savingType
	}
}

// generateJSONReport 生成JSON格式的价格报告
func generateJSONReport(pricingTable *config.EC2PricingTable) {
	report := make(map[string]interface{})
	report["title"] = "EC2实例价格对比报告"
	report["records"] = pricingTable.Records

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Printf("生成JSON报告失败: %v", err)
		return
	}

	fmt.Printf("JSON报告已生成 (前500字符):\n%s...\n", string(jsonData[:min(500, len(jsonData))]))
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
