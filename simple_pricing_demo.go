package main

import (
	"fmt"
	"sort"

	"analysis-aws-bill/config"
)

func main() {
	fmt.Println("=== EC2实例价格查询系统 ===")
	fmt.Println()

	// 示例1: 查询特定实例类型
	fmt.Println("1. 查询特定实例类型价格:")
	instanceTypes := []string{
		"APS1-BoxUsage:c6a.xlarge",
		"APS1-BoxUsage:c5.large",
		"APE1-BoxUsage:t3.micro",
		"APN1-BoxUsage:c5.xlarge",
	}

	for _, instanceType := range instanceTypes {
		pricing := config.GetEC2Pricing(instanceType)
		if pricing != nil {
			fmt.Printf("\n--- %s ---\n", instanceType)
			pricing.PrintPricingInfo()
		} else {
			fmt.Printf("\n未找到 %s 的价格信息\n", instanceType)
		}
	}

	// 示例2: 搜索包含关键词的实例
	fmt.Println("\n\n2. 搜索包含 'c5' 的实例:")
	c5Instances := config.SearchEC2Pricing("c5")
	for _, pricing := range c5Instances {
		bestPrice, bestOption, savingsPercent := pricing.GetBestPricing()
		fmt.Printf("%-25s %s $%.2f -> $%.2f (%s, 节省%.1f%%)\n",
			pricing.InstanceType,
			getShortRegion(pricing.Region),
			pricing.OnDemandMonthly,
			bestPrice,
			bestOption,
			savingsPercent)
	}

	// 示例3: 生成完整价格对比表
	fmt.Println("\n\n3. 完整价格对比表:")
	generatePricingTable()

	// 示例4: 成本优化建议
	fmt.Println("\n\n4. 成本优化建议:")
	generateOptimizationReport()
}

func getShortRegion(region string) string {
	switch region {
	case "Asia Pacific (Hong Kong)":
		return "香港"
	case "Asia Pacific (Singapore)":
		return "新加坡"
	case "Asia Pacific (Tokyo)":
		return "东京"
	case "EU (Frankfurt)":
		return "法兰克福"
	case "US East (N. Virginia)":
		return "弗吉尼亚"
	default:
		return region
	}
}

func generatePricingTable() {
	allPricing := config.GetAllEC2Pricing()

	fmt.Printf("%-30s %-10s %-10s %-10s %-10s %-10s %-10s %-20s\n",
		"实例类型", "地区", "按需/月", "SP无预付", "SP全预付", "RI无预付", "RI全预付", "最佳方案")
	fmt.Println("-----------------------------------------------------------------------------------------------------------")

	// 按实例类型排序
	var sortedKeys []string
	for key := range allPricing {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		pricing := allPricing[key]
		_, bestOption, savingsPercent := pricing.GetBestPricing()

		bestInfo := fmt.Sprintf("%s(%.1f%%)", bestOption, savingsPercent)
		if savingsPercent == 0 {
			bestInfo = "按需最优"
		}

		fmt.Printf("%-30s %-10s $%-9.2f $%-9.2f $%-9.2f $%-9.2f $%-9.2f %-20s\n",
			pricing.InstanceType,
			getShortRegion(pricing.Region),
			pricing.OnDemandMonthly,
			pricing.SP3YearNoUpfront,
			pricing.SP3YearAllUpfront,
			pricing.RI3YearNoUpfront,
			pricing.RI3YearAllUpfront,
			bestInfo,
		)
	}
}

func generateOptimizationReport() {
	allPricing := config.GetAllEC2Pricing()

	type OptimizationItem struct {
		InstanceType   string
		Region         string
		OnDemandCost   float64
		OptimizedCost  float64
		SavingsAmount  float64
		SavingsPercent float64
		BestOption     string
	}

	var items []OptimizationItem
	totalOnDemand := 0.0
	totalOptimized := 0.0

	for _, pricing := range allPricing {
		bestPrice, bestOption, savingsPercent := pricing.GetBestPricing()
		savingsAmount := pricing.OnDemandMonthly - bestPrice

		items = append(items, OptimizationItem{
			InstanceType:   pricing.InstanceType,
			Region:         pricing.Region,
			OnDemandCost:   pricing.OnDemandMonthly,
			OptimizedCost:  bestPrice,
			SavingsAmount:  savingsAmount,
			SavingsPercent: savingsPercent,
			BestOption:     bestOption,
		})

		totalOnDemand += pricing.OnDemandMonthly
		totalOptimized += bestPrice
	}

	// 按节省金额排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].SavingsAmount > items[j].SavingsAmount
	})

	fmt.Println("按节省金额排序的优化建议 (前10项):")
	fmt.Printf("%-30s %-10s %-12s %-15s %-15s\n", "实例类型", "地区", "月节省金额", "节省百分比", "推荐方案")
	fmt.Println("---------------------------------------------------------------------------------")

	for i, item := range items {
		if i >= 10 { // 只显示前10项
			break
		}
		if item.SavingsAmount > 0 {
			fmt.Printf("%-30s %-10s $%-11.2f %-14.1f%% %-15s\n",
				item.InstanceType,
				getShortRegion(item.Region),
				item.SavingsAmount,
				item.SavingsPercent,
				item.BestOption,
			)
		}
	}

	totalSavings := totalOnDemand - totalOptimized
	totalSavingsPercent := totalSavings / totalOnDemand * 100

	fmt.Println("\n总体优化效果:")
	fmt.Printf("当前按需总费用: $%.2f/月\n", totalOnDemand)
	fmt.Printf("优化后总费用:   $%.2f/月\n", totalOptimized)
	fmt.Printf("总节省金额:     $%.2f/月 (%.1f%%)\n", totalSavings, totalSavingsPercent)
	fmt.Printf("年度节省:       $%.2f\n", totalSavings*12)

	fmt.Println("\n优化建议:")
	fmt.Println("1. 优先优化高费用实例以获得更大的绝对节省")
	fmt.Println("2. 对于稳定工作负载，选择3年全预付RI获得最大折扣")
	fmt.Println("3. 对于不确定工作负载，选择3年无预付SP保持灵活性")
	fmt.Println("4. 定期评估实际使用情况，调整预留实例配置")
}
