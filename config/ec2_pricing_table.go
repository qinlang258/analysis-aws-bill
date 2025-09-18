package config

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// EC2PricingRecord EC2实例价格记录结构
type EC2PricingRecord struct {
	AccountID         string  `json:"account_id"`        // AWS账户ID
	Region            string  `json:"region"`            // 地区名称
	RegionCode        string  `json:"region_code"`       // 地区代码
	InstanceType      string  `json:"instance_type"`     // 实例类型 (如: t3.micro, c5.large)
	UsageType         string  `json:"usage_type"`        // 使用类型 (如: APE1-BoxUsage:t3.micro)
	OnDemandHourly    float64 `json:"on_demand_hourly"`  // 按需实例每小时价格
	OnDemandMonthly   float64 `json:"on_demand_monthly"` // 按需实例每月价格 (730小时)
	SP3YearNoUpfront  float64 `json:"sp_3y_no_upfront"`  // Savings Plan 3年无预付每月价格
	SP3YearAllUpfront float64 `json:"sp_3y_all_upfront"` // Savings Plan 3年全预付每月价格
	RI3YearNoUpfront  float64 `json:"ri_3y_no_upfront"`  // Reserved Instance 3年无预付每月价格
	RI3YearAllUpfront float64 `json:"ri_3y_all_upfront"` // Reserved Instance 3年全预付每月价格
}

// EC2PricingTable EC2价格表
type EC2PricingTable struct {
	Records map[string]*EC2PricingRecord // key: accountID-region-instanceType
}

// NewEC2PricingTable 创建新的EC2价格表
func NewEC2PricingTable() *EC2PricingTable {
	return &EC2PricingTable{
		Records: make(map[string]*EC2PricingRecord),
	}
}

// generateKey 生成记录的唯一键
func (table *EC2PricingTable) generateKey(accountID, region, instanceType string) string {
	return fmt.Sprintf("%s-%s-%s", accountID, region, instanceType)
}

// AddRecord 添加或更新价格记录
func (table *EC2PricingTable) AddRecord(record *EC2PricingRecord) {
	key := table.generateKey(record.AccountID, record.Region, record.InstanceType)
	table.Records[key] = record
}

// GetRecord 获取价格记录
func (table *EC2PricingTable) GetRecord(accountID, region, instanceType string) *EC2PricingRecord {
	key := table.generateKey(accountID, region, instanceType)
	return table.Records[key]
}

// LoadOnDemandPricesFromCSV 从CSV文件加载按需价格
func (table *EC2PricingTable) LoadOnDemandPricesFromCSV(csvFilePath string) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("无法打开CSV文件: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("读取CSV文件失败: %v", err)
	}

	// 跳过标题行
	for i, record := range records {
		if i == 0 {
			continue
		}

		// 只处理EC2相关的使用记录
		if len(record) < 18 || record[5] != "Usage" || record[8] != "Amazon Elastic Compute Cloud" {
			continue
		}

		// 只处理BoxUsage类型的记录
		if !strings.Contains(record[9], "BoxUsage:") {
			continue
		}

		// 解析数据
		accountID := record[3]
		region := record[7]
		regionCode := record[6]
		usageType := record[9]

		// 从UsageType中提取实例类型 (如: APE1-BoxUsage:t3.micro -> t3.micro)
		instanceType := ""
		if parts := strings.Split(usageType, ":"); len(parts) == 2 {
			instanceType = parts[1]
		}

		// 解析按需价格
		hourlyRate, err := strconv.ParseFloat(record[15], 64)
		if err != nil {
			continue
		}

		// 计算每月价格 (730小时)
		monthlyRate := hourlyRate * 730

		// 创建或更新记录
		key := table.generateKey(accountID, region, instanceType)
		if existingRecord, exists := table.Records[key]; exists {
			// 更新现有记录
			existingRecord.OnDemandHourly = hourlyRate
			existingRecord.OnDemandMonthly = monthlyRate
		} else {
			// 创建新记录
			newRecord := &EC2PricingRecord{
				AccountID:       accountID,
				Region:          region,
				RegionCode:      regionCode,
				InstanceType:    instanceType,
				UsageType:       usageType,
				OnDemandHourly:  hourlyRate,
				OnDemandMonthly: monthlyRate,
			}
			table.Records[key] = newRecord
		}
	}

	return nil
}

// updatePricesFromMap 从价格映射更新记录价格
func (table *EC2PricingTable) updatePricesFromMap(record *EC2PricingRecord, instancePrices map[string]float64) {
	// 更新SP和RI价格
	if sp3NoUpfront, exists := instancePrices["sp-3-no-month"]; exists {
		record.SP3YearNoUpfront = sp3NoUpfront
	} else if sp3NoUpfront, exists := instancePrices["sp-3-no"]; exists {
		record.SP3YearNoUpfront = sp3NoUpfront
	}

	if sp3AllUpfront, exists := instancePrices["sp-3-all-month"]; exists {
		record.SP3YearAllUpfront = sp3AllUpfront
	} else if sp3AllUpfront, exists := instancePrices["sp-3-all"]; exists {
		record.SP3YearAllUpfront = sp3AllUpfront
	}

	if ri3NoUpfront, exists := instancePrices["ri-3-no-month"]; exists {
		record.RI3YearNoUpfront = ri3NoUpfront
	} else if ri3NoUpfront, exists := instancePrices["ri-3-no"]; exists {
		record.RI3YearNoUpfront = ri3NoUpfront
	}

	if ri3AllUpfront, exists := instancePrices["ri-3-all-month"]; exists {
		record.RI3YearAllUpfront = ri3AllUpfront
	} else if ri3AllUpfront, exists := instancePrices["ri-3-all"]; exists {
		record.RI3YearAllUpfront = ri3AllUpfront
	}
}

// GetRegionInstanceTypePrice 获取全局价格数据（用于调试）
func GetRegionInstanceTypePrice() map[string]map[string]map[string]map[string]float64 {
	return regionInstanceTypePrice
}

// PrintTable 打印价格表
func (table *EC2PricingTable) PrintTable() {
	fmt.Printf("%-15s %-25s %-15s %-12s %-12s %-12s %-12s %-12s %-12s\n",
		"账户ID", "地区", "实例类型", "按需/小时", "按需/月", "SP3年无预付", "SP3年全预付", "RI3年无预付", "RI3年全预付")
	fmt.Println(strings.Repeat("-", 150))

	for _, record := range table.Records {
		fmt.Printf("%-15s %-25s %-15s $%-11.4f $%-11.2f $%-11.2f $%-11.2f $%-11.2f $%-11.2f\n",
			record.AccountID,
			record.Region,
			record.InstanceType,
			record.OnDemandHourly,
			record.OnDemandMonthly,
			record.SP3YearNoUpfront,
			record.SP3YearAllUpfront,
			record.RI3YearNoUpfront,
			record.RI3YearAllUpfront,
		)
	}
}

// GetSavingsComparison 获取节省对比
func (table *EC2PricingTable) GetSavingsComparison(accountID, region, instanceType string) map[string]interface{} {
	record := table.GetRecord(accountID, region, instanceType)
	if record == nil {
		return nil
	}

	onDemandMonthly := record.OnDemandMonthly
	result := map[string]interface{}{
		"instance_info": map[string]string{
			"account_id":    accountID,
			"region":        region,
			"instance_type": instanceType,
		},
		"pricing": map[string]float64{
			"on_demand_monthly": onDemandMonthly,
			"sp_3y_no_upfront":  record.SP3YearNoUpfront,
			"sp_3y_all_upfront": record.SP3YearAllUpfront,
			"ri_3y_no_upfront":  record.RI3YearNoUpfront,
			"ri_3y_all_upfront": record.RI3YearAllUpfront,
		},
		"savings": map[string]interface{}{},
	}

	// 计算节省百分比和金额
	savings := result["savings"].(map[string]interface{})

	if record.SP3YearNoUpfront > 0 {
		savingsAmount := onDemandMonthly - record.SP3YearNoUpfront
		savingsPercent := (savingsAmount / onDemandMonthly) * 100
		savings["sp_3y_no_upfront"] = map[string]float64{
			"amount":  savingsAmount,
			"percent": savingsPercent,
		}
	}

	if record.SP3YearAllUpfront > 0 {
		savingsAmount := onDemandMonthly - record.SP3YearAllUpfront
		savingsPercent := (savingsAmount / onDemandMonthly) * 100
		savings["sp_3y_all_upfront"] = map[string]float64{
			"amount":  savingsAmount,
			"percent": savingsPercent,
		}
	}

	if record.RI3YearNoUpfront > 0 {
		savingsAmount := onDemandMonthly - record.RI3YearNoUpfront
		savingsPercent := (savingsAmount / onDemandMonthly) * 100
		savings["ri_3y_no_upfront"] = map[string]float64{
			"amount":  savingsAmount,
			"percent": savingsPercent,
		}
	}

	if record.RI3YearAllUpfront > 0 {
		savingsAmount := onDemandMonthly - record.RI3YearAllUpfront
		savingsPercent := (savingsAmount / onDemandMonthly) * 100
		savings["ri_3y_all_upfront"] = map[string]float64{
			"amount":  savingsAmount,
			"percent": savingsPercent,
		}
	}

	return result
}
