package config

// 全局变量存储价格数据
var regionInstanceTypePrice map[string]map[string]map[string]map[string]float64

// ensureMapExists 确保嵌套map存在，如果不存在则创建
func ensureMapExists(accountID, region, instanceType string) {
	if regionInstanceTypePrice == nil {
		regionInstanceTypePrice = make(map[string]map[string]map[string]map[string]float64)
	}
	if regionInstanceTypePrice[accountID] == nil {
		regionInstanceTypePrice[accountID] = make(map[string]map[string]map[string]float64)
	}
	if regionInstanceTypePrice[accountID][region] == nil {
		regionInstanceTypePrice[accountID][region] = make(map[string]map[string]float64)
	}
	if regionInstanceTypePrice[accountID][region][instanceType] == nil {
		regionInstanceTypePrice[accountID][region][instanceType] = make(map[string]float64)
	}
}

// setPricing 设置实例价格的辅助函数
func setPricing(accountID, region, instanceType string, prices map[string]float64) {
	ensureMapExists(accountID, region, instanceType)
	for priceType, price := range prices {
		regionInstanceTypePrice[accountID][region][instanceType][priceType] = price
	}
}

// MakeEc2InstanceTypeCost 创建EC2实例类型成本配置
// 参数 region: 地区名称
// 该函数初始化不同AWS账户在不同地区的EC2实例类型价格数据
