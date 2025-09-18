package config

import (
	"fmt"
	"strings"
)

// EC2PricingInfo EC2实例价格信息
type EC2PricingInfo struct {
	InstanceType      string  `json:"instance_type"`     // 实例类型 (如: APS1-BoxUsage:c6a.xlarge)
	OnDemandMonthly   float64 `json:"on_demand_monthly"` // 按需实例每月价格
	SP3YearNoUpfront  float64 `json:"sp_3y_no_upfront"`  // Savings Plan 3年无预付每月价格
	SP3YearAllUpfront float64 `json:"sp_3y_all_upfront"` // Savings Plan 3年全预付每月价格
	RI3YearNoUpfront  float64 `json:"ri_3y_no_upfront"`  // Reserved Instance 3年无预付每月价格
	RI3YearAllUpfront float64 `json:"ri_3y_all_upfront"` // Reserved Instance 3年全预付每月价格
	Region            string  `json:"region"`            // 地区
	AccountID         string  `json:"account_id"`        // 账户ID
}

// EC2PricingMap 全局价格映射表
var EC2PricingMap map[string]*EC2PricingInfo

//   初始化EC2价格数据
func  main() {
	EC2PricingMap = make(map[string]*EC2PricingInfo)

	// 账户 011528268154 - 新加坡地区
	EC2PricingMap["APS1-BoxUsage:c6a.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c6a.xlarge",
		SP3YearNoUpfront:  66.70,
		SP3YearAllUpfront: 2178.87/36,
		RI3YearNoUpfront:  58.41,
		RI3YearAllUpfront: 1830.00/36,
	}

	// 账户 011528268154 - 东京地区
	EC2PricingMap["APN1-BoxUsage:c6a.large"] = &EC2PricingInfo{
		InstanceType:      "APN1-BoxUsage:c6a.large",
		SP3YearNoUpfront:  73.20,
		SP3YearAllUpfront: 66.42,
		RI3YearNoUpfront:  64.07,
		RI3YearAllUpfront: 55.78,
	}

	EC2PricingMap["APN1-BoxUsage:c6a.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APN1-BoxUsage:c6a.xlarge",
		SP3YearNoUpfront:  146.40,
		SP3YearAllUpfront: 132.85,
		RI3YearNoUpfront:  128.14,
		RI3YearAllUpfront: 111.56,
	}

	EC2PricingMap["APN1-BoxUsage:t2.micro"] = &EC2PricingInfo{
		InstanceType:      "APN1-BoxUsage:t2.micro",
		SP3YearNoUpfront:  5.69,
		SP3YearAllUpfront: 5.69,
		RI3YearNoUpfront:  5.48,
		RI3YearAllUpfront: 4.78,
	}

	// 账户 017820696133 - 新加坡地区
	EC2PricingMap["APS1-BoxUsage:c5.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c5.large",
		SP3YearNoUpfront:  37.23,
		SP3YearAllUpfront: 33.58,
		RI3YearNoUpfront:  30.66,
		RI3YearAllUpfront: 26.89,
	}

	EC2PricingMap["APS1-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c5.xlarge",
		SP3YearNoUpfront:  74.46,
		SP3YearAllUpfront: 67.16,
		RI3YearNoUpfront:  61.32,
		RI3YearAllUpfront: 53.78,
	}

	EC2PricingMap["APS1-BoxUsage:c6a.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c6a.large",
		SP3YearNoUpfront:  33.35,
		SP3YearAllUpfront: 30.26,
		RI3YearNoUpfront:  29.21,
		RI3YearAllUpfront: 25.42,
	}

	EC2PricingMap["APS1-BoxUsage:c6a.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c6a.xlarge",
		SP3YearNoUpfront:  66.70,
		SP3YearAllUpfront: 60.52,
		RI3YearNoUpfront:  58.42,
		RI3YearAllUpfront: 50.83,
	}

	EC2PricingMap["APS1-BoxUsage:c6a.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c6a.2xlarge",
		SP3YearNoUpfront:  133.40,
		SP3YearAllUpfront: 121.03,
		RI3YearNoUpfront:  116.84,
		RI3YearAllUpfront: 101.67,
	}

	EC2PricingMap["APS1-BoxUsage:c6a.4xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:c6a.4xlarge",
		SP3YearNoUpfront:  266.80,
		SP3YearAllUpfront: 242.07,
		RI3YearNoUpfront:  233.68,
		RI3YearAllUpfront: 203.33,
	}

	EC2PricingMap["APS1-BoxUsage:t2.micro"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t2.micro",
		SP3YearNoUpfront:  5.18,
		SP3YearAllUpfront: 4.75,
		RI3YearNoUpfront:  4.53,
		RI3YearAllUpfront: 3.92,
	}

	EC2PricingMap["APS1-BoxUsage:r6i.16xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:r6i.16xlarge",
		SP3YearNoUpfront:  1852.20,
		SP3YearAllUpfront: 60505.23/36,
		RI3YearNoUpfront:  1610.61,
		RI3YearAllUpfront: 50466.00/36,
	}

	// 账户 017820696133 - 法兰克福地区
	EC2PricingMap["EUC1-BoxUsage:c6a.large"] = &EC2PricingInfo{
		InstanceType:      "EUC1-BoxUsage:c6a.large",
		SP3YearNoUpfront:  33.07,
		SP3YearAllUpfront: 30.00,
		RI3YearNoUpfront:  28.91,
		RI3YearAllUpfront: 25.17,
	}

	// 账户 017820696133 - 美国东部
	EC2PricingMap["BoxUsage:c6a.large"] = &EC2PricingInfo{
		InstanceType:      "BoxUsage:c6a.large",
		SP3YearNoUpfront:  26.97,
		SP3YearAllUpfront: 24.47,
		RI3YearNoUpfront:  24.59,
		RI3YearAllUpfront: 21.39,
	}

	EC2PricingMap["BoxUsage:t2.micro"] = &EC2PricingInfo{
		InstanceType:      "BoxUsage:t2.micro",
		SP3YearNoUpfront:  4.23,
		SP3YearAllUpfront: 3.80,
		RI3YearNoUpfront:  3.65,
		RI3YearAllUpfront: 3.19,
	}

	// 账户 047719631469 - 香港地区
	EC2PricingMap["APE1-BoxUsage:c6i.8xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c6i.8xlarge",
		SP3YearNoUpfront:  652.67,
		SP3YearAllUpfront: 21320.44/36,
		RI3YearNoUpfront:  572.19,
		RI3YearAllUpfront: 17929.00/36,
	}

	EC2PricingMap["APE1-BoxUsage:c6i.12xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c6i.12xlarge",
		SP3YearNoUpfront:  979.00,
		SP3YearAllUpfront: 31980.66/36,
		RI3YearNoUpfront:  1251.67,
		RI3YearAllUpfront: 26893.00/36,
	}	


	EC2PricingMap["APE1-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c5.xlarge",
		SP3YearNoUpfront:  81.76,
		SP3YearAllUpfront: 73.73,
		RI3YearNoUpfront:  67.89,
		RI3YearAllUpfront: 59.28,
	}

	EC2PricingMap["APE1-BoxUsage:t3.medium"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.medium",
		SP3YearNoUpfront:  20.81,
		SP3YearAllUpfront: 678.02/36,
		RI3YearNoUpfront:  18.03,
		RI3YearAllUpfront: 566.00/36,
	}

	EC2PricingMap["APN1-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APN1-BoxUsage:c5.xlarge",
		SP3YearNoUpfront:  91.98,
		SP3YearAllUpfront: 83.22,
		RI3YearNoUpfront:  67.17,
		RI3YearAllUpfront: 58.72,
	}

	EC2PricingMap["APN1-BoxUsage:t3.medium"] = &EC2PricingInfo{
		InstanceType:      "APN1-BoxUsage:t3.medium",
		SP3YearNoUpfront:  22.56,
		SP3YearAllUpfront: 20.44,
		RI3YearNoUpfront:  17.16,
		RI3YearAllUpfront: 14.94,
	}

	EC2PricingMap["APS1-BoxUsage:t3.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.xlarge",
		SP3YearNoUpfront:  75.19,
		SP3YearAllUpfront: 68.18,
		RI3YearNoUpfront:  65.34,
		RI3YearAllUpfront: 56.83,
	}

	EC2PricingMap["APE1-BoxUsage:t3.micro"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.micro",
		SP3YearNoUpfront:  5.18,
		SP3YearAllUpfront: 170.82/36,
		RI3YearNoUpfront:  4.53,
		RI3YearAllUpfront: 141.00/36,
	}

	EC2PricingMap["APE1-BoxUsage:t3.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.xlarge",
		SP3YearNoUpfront:  83.15,
		SP3YearAllUpfront: 2717.35/36,
		RI3YearNoUpfront:  72.20,
		RI3YearAllUpfront: 2263.00/36,
	}	

	EC2PricingMap["APE1-BoxUsage:t3.small"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.small",
		SP3YearNoUpfront:  10.37,
		SP3YearAllUpfront: 339.01/36,
		RI3YearNoUpfront:  9.05,
		RI3YearAllUpfront: 283.00/36,
	}		

	EC2PricingMap["APE1-BoxUsage:t3.large"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.large",
		SP3YearNoUpfront:  41.61,
		SP3YearAllUpfront: 1358.68/36,
		RI3YearNoUpfront:  36.14,
		RI3YearAllUpfront: 1132.00/36,
	}		

	EC2PricingMap["APE1-BoxUsage:t3.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:t3.2xlarge",
		SP3YearNoUpfront:  166.29,
		SP3YearAllUpfront: 5432.08/36,
		RI3YearNoUpfront:  144.47,
		RI3YearAllUpfront: 4526.00/36,
	}		

	EC2PricingMap["APE1-BoxUsage:r5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:r5.xlarge",
		SP3YearNoUpfront:  127.02,
		SP3YearAllUpfront: 4152.24/36,
		RI3YearNoUpfront:  105.12,
		RI3YearAllUpfront: 3300.00/36,
	}		

	EC2PricingMap["APE1-BoxUsage:r5.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:r5.xlarge",
		SP3YearNoUpfront:  254.04,
		SP3YearAllUpfront: 8304.48/36,
		RI3YearNoUpfront:  210.97,
		RI3YearAllUpfront: 6601.00/36,
	}	

	EC2PricingMap["APE1-BoxUsage:m5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:m5.xlarge",
		SP3YearNoUpfront:  97.82,
		SP3YearAllUpfront: 3179.88/36,
		RI3YearNoUpfront:  83.22,
		RI3YearAllUpfront: 2609.00/36,
	}	

	EC2PricingMap["APE1-BoxUsage:m5.large"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:m5.xlarge",
		SP3YearNoUpfront: 97.82,
		SP3YearAllUpfront: 3179.88/36,
		RI3YearNoUpfront: 83.22,
		RI3YearAllUpfront: 2609.00/36,
	}	

	EC2PricingMap["APE1-BoxUsage:m5.4xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:m5.4xlarge",
		SP3YearNoUpfront: 391.28,
		SP3YearAllUpfront: 12772.08/36,
		RI3YearNoUpfront: 332.88,
		RI3YearAllUpfront: 10435.00/36,
	}	
	
	EC2PricingMap["APE1-BoxUsage:m5.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:m5.2xlarge",
		SP3YearNoUpfront: 195.64,
		SP3YearAllUpfront: 6386.04/36,
		RI3YearNoUpfront: 166.44,
		RI3YearAllUpfront: 5217.00/36,
	}	
	
	EC2PricingMap["APE1-BoxUsage:c5d.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c5d.2xlarge",
		SP3YearNoUpfront: 184.69,
		SP3YearAllUpfront: 6018.12/36,
		RI3YearNoUpfront: 155.49,
		RI3YearAllUpfront: 4862.00/36,
	}		


	EC2PricingMap["APE1-BoxUsage:c5a.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c5a.xlarge",
		SP3YearNoUpfront: 73.73,
		SP3YearAllUpfront: 2391.48/36,
		RI3YearNoUpfront: 61.32,
		RI3YearAllUpfront: 1917.00/36,
	}		

	EC2PricingMap["APE1-BoxUsage:c5.large"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c5.large",
		SP3YearNoUpfront: 40.88,
		SP3YearAllUpfront: 1340.28/36,
		RI3YearNoUpfront: 34.31,
		RI3YearAllUpfront: 1067.00/36,
	}	
	
	EC2PricingMap["APE1-BoxUsage:c5.4xlarge"] = &EC2PricingInfo{
		InstanceType:      "APE1-BoxUsage:c5.4xlarge",
		SP3YearNoUpfront: 326.31,
		SP3YearAllUpfront: 10669.68/36,
		RI3YearNoUpfront: 272.29,
		RI3YearAllUpfront: 8537.00/36,
	}	

	EC2PricingMap["USE2-BoxUsage:r7i.xlarge"] = &EC2PricingInfo{
		InstanceType:      "USE2-BoxUsage:r7i.xlarge",
		SP3YearNoUpfront: 100.76,
		SP3YearAllUpfront: 3291.57/36,
		RI3YearNoUpfront: 87.61,
		RI3YearAllUpfront: 2745.00/36,
	}	


	EC2PricingMap["USE2-BoxUsage:m7i.xlarge"] = &EC2PricingInfo{
		InstanceType:      "USE2-BoxUsage:m7i.xlarge",
		SP3YearNoUpfront: 74.64,
		SP3YearAllUpfront: 2438.00/36,
		RI3YearNoUpfront: 66.76,
		RI3YearAllUpfront: 2092.00/36,
	}	
	
	EC2PricingMap["MEC1-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "MEC1-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 84.68,
		SP3YearAllUpfront: 2785.68/36,
		RI3YearNoUpfront: 66.43,
		RI3YearAllUpfront: 2087.00/36,
	}		

	EC2PricingMap["EUW2-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "EUW2-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 81.03,
		SP3YearAllUpfront: 2654.28/36,
		RI3YearNoUpfront: 63.51,
		RI3YearAllUpfront: 1996.00/36,
	}	

	EC2PricingMap["EUC1-BoxUsage:t2.medium"] = &EC2PricingInfo{
		InstanceType:      "EUC1-BoxUsage:t2.medium",
		SP3YearNoUpfront: 23.00,
		SP3YearAllUpfront: 751.61/36,
		RI3YearNoUpfront: 20.15,
		RI3YearAllUpfront: 638.00/36,
	}		

	EC2PricingMap["EUC1-BoxUsage:r6i.16xlarge"] = &EC2PricingInfo{
		InstanceType:      "EUC1-BoxUsage:r6i.16xlarge",
		SP3YearNoUpfront: 1852.20,
		SP3YearAllUpfront: 60505.23/36,
		RI3YearNoUpfront: 1610.61,
		RI3YearAllUpfront: 50466.00/36,
	}		

	EC2PricingMap["APS6-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS6-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 83.22,
		SP3YearAllUpfront: 2733.12/36,
		RI3YearNoUpfront: 70.08,
		RI3YearAllUpfront: 2194.00/36,
	}		

	EC2PricingMap["APS4-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS4-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 73.73,
		SP3YearAllUpfront: 2417.76/36,
		RI3YearNoUpfront: 62.05,
		RI3YearAllUpfront: 1937.00/36,
	}	
	
	EC2PricingMap["APS3-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS3-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 64.24,
		SP3YearAllUpfront: 2102.40/36,
		RI3YearNoUpfront: 53.29,
		RI3YearAllUpfront: 1680.00/36,
	}	

	EC2PricingMap["APS2-BoxUsage:c5.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS2-BoxUsage:c5.xlarge",
		SP3YearNoUpfront: 83.22,
		SP3YearAllUpfront: 2733.12/36,
		RI3YearNoUpfront: 70.08,
		RI3YearAllUpfront: 2194.00/36,
	}	
	
	EC2PricingMap["APS1-BoxUsage:t3a.medium"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3a.medium",
		SP3YearNoUpfront: 16.79,
		SP3YearAllUpfront: 549.25/36,
		RI3YearNoUpfront: 14.60,
		RI3YearAllUpfront: 457.00/36,
	}	

	EC2PricingMap["APS1-BoxUsage:t3.small"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.small",
		SP3YearNoUpfront: 9.42,
		SP3YearAllUpfront: 307.48/36,
		RI3YearNoUpfront: 8.18,
		RI3YearAllUpfront: 256.00/36,
	}		


	EC2PricingMap["APS1-BoxUsage:t3.micro"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.micro",
		SP3YearNoUpfront: 4.67,
		SP3YearAllUpfront: 152.42/36,
		RI3YearNoUpfront: 4.09,
		RI3YearAllUpfront: 128.00/36,
	}	

	EC2PricingMap["APS1-BoxUsage:t3.medium"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.medium",
		SP3YearNoUpfront: 18.76,
		SP3YearAllUpfront: 614.95/36,
		RI3YearNoUpfront: 16.35,
		RI3YearAllUpfront: 512.00/36,
	}
	
	EC2PricingMap["APS1-BoxUsage:t3.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.large",
		SP3YearNoUpfront: 37.60,
		SP3YearAllUpfront: 1227.28/36,
		RI3YearNoUpfront: 32.63,
		RI3YearAllUpfront: 1023.00/36,
	}		

	EC2PricingMap["APS1-BoxUsage:t3.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.large",
		SP3YearNoUpfront: 10.37,
		SP3YearAllUpfront: 339.01/36,
		RI3YearNoUpfront: 9.05,
		RI3YearAllUpfront: 283.00/36,
	}	
	
	EC2PricingMap["APS1-BoxUsage:t3.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t3.large",
		SP3YearNoUpfront: 10.37,
		SP3YearAllUpfront: 339.01/36,
		RI3YearNoUpfront: 9.05,
		RI3YearAllUpfront: 283.00/36,
	}	

	EC2PricingMap["APS1-BoxUsage:t2.medium"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t2.medium",
		SP3YearNoUpfront: 20.81,
		SP3YearAllUpfront: 678.02/36,
		RI3YearNoUpfront: 18.03,
		RI3YearAllUpfront: 566.00/36,
	}	

	EC2PricingMap["APS1-BoxUsage:t2.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:t2.large",
		SP3YearNoUpfront: 41.54,
		SP3YearAllUpfront: 1356.05/36,
		RI3YearNoUpfront: 36.06,
		RI3YearAllUpfront: 1131.00/36,
	}		


	EC2PricingMap["APS1-BoxUsage:r6i.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:r6i.xlarge",
		SP3YearNoUpfront: 115.76,
		SP3YearAllUpfront: 3781.69/36,
		RI3YearNoUpfront: 100.66,
		RI3YearAllUpfront: 3154.00/36,
	}	

	EC2PricingMap["APS1-BoxUsage:r6a.xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:r6a.xlarge",
		SP3YearNoUpfront: 104.19,
		SP3YearAllUpfront: 3403.52/36,
		RI3YearNoUpfront: 90.59,
		RI3YearAllUpfront: 2839.00/36,
	}	
	
	EC2PricingMap["APS1-BoxUsage:r6a.2xlarge"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:r6a.2xlarge",
		SP3YearNoUpfront: 208.37,
		SP3YearAllUpfront: 6806.78/36,
		RI3YearNoUpfront: 181.19,
		RI3YearAllUpfront: 5677.00/36,
	}	
	
	EC2PricingMap["APS1-BoxUsage:r5a.large"] = &EC2PricingInfo{
		InstanceType:      "APS1-BoxUsage:r5a.large",
		SP3YearNoUpfront: 51.83,
		SP3YearAllUpfront: 1708.20/36,
		RI3YearNoUpfront: 43.07,
		RI3YearAllUpfront: 1344.00/36,
	}		
}

// GetEC2Pricing 根据实例类型获取价格信息
func GetEC2Pricing(instanceType string) *EC2PricingInfo {
	if EC2PricingMap == nil {
		//  ()
	}
	return EC2PricingMap[instanceType]
}

// GetAllEC2Pricing 获取所有价格信息
func GetAllEC2Pricing() map[string]*EC2PricingInfo {
	if EC2PricingMap == nil {
		 ()
	}
	return EC2PricingMap
}

// SearchEC2Pricing 搜索包含指定关键词的实例类型
func SearchEC2Pricing(keyword string) []*EC2PricingInfo {
	if EC2PricingMap == nil {
		 ()
	}

	var results []*EC2PricingInfo
	for instanceType, pricing := range EC2PricingMap {
		if strings.Contains(strings.ToLower(instanceType), strings.ToLower(keyword)) {
			results = append(results, pricing)
		}
	}
	return results
}



// PrintPricingInfo 打印价格信息
func (p *EC2PricingInfo) PrintPricingInfo() {
	fmt.Printf("实例类型: %s\n", p.InstanceType)
	fmt.Printf("地区: %s\n", p.Region)
	fmt.Printf("账户: %s\n", p.AccountID)
	fmt.Printf("按需价格: $%.2f/月\n", p.OnDemandMonthly)
	fmt.Printf("SP 3年无预付: $%.2f/月\n", p.SP3YearNoUpfront)
	fmt.Printf("SP 3年全预付: $%.2f/月\n", p.SP3YearAllUpfront)
	fmt.Printf("RI 3年无预付: $%.2f/月\n", p.RI3YearNoUpfront)
	fmt.Printf("RI 3年全预付: $%.2f/月\n", p.RI3YearAllUpfront)

	bestPrice, bestOption, savingsPercent := p.GetBestPricing()
	if savingsPercent > 0 {
		fmt.Printf("最佳方案: %s - $%.2f/月 (节省%.1f%%)\n", bestOption, bestPrice, savingsPercent)
		fmt.Printf("年度节省: $%.2f\n", (p.OnDemandMonthly-bestPrice)*12)
	} else {
		fmt.Printf("最佳方案: 按需 - $%.2f/月\n", bestPrice)
	}
}
