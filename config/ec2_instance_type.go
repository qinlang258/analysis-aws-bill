package config

func MakeEc2InstanceTypeCost(region string) {
	//这里是算的每个月 730小时
	regionInstanceTypePrice := make(map[string]map[string]map[string]map[string]float64, 0)

	switch "lineItem/UsageAccountId" {

	case "011528268154":
		//新加坡
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["sp-3-all-month"] = 2178.87 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["sp-3-no-month"] = 66.70
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["ri-3-all-month"] = 1830.00 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["ri-3-no-month"] = 58.41

		//日本
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"] = map[string]float64{}
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["sp-3-all-month"] = 2391.22 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["sp-3-no-month"] = 73.20
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["ri-3-all-month"] = 2008.00 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["ri-3-no-month"] = 64.07

		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["sp-3-all-month"] = 2391.22 * 2 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["sp-3-no-month"] = 73.20 * 2
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["ri-3-all-month"] = 2008.00 * 2 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["ri-3-no-month"] = 64.07 * 2

		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"] = map[string]float64{}
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["sp-3-all-month"] = 204.98 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["sp-3-no-month"] = 5.69
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["ri-3-all-month"] = 172.00 / 36
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["ri-3-no-month"] = 5.48
		regionInstanceTypePrice["011528268154"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["OnDemand-day"] = 0.0152

	case "017820696133":
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.large"]["sp-3-all-month"] = 1208.88 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.large"]["sp-3-no-month"] = 37.23
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.large"]["ri-3-all-month"] = 968.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.large"]["ri-3-no-month"] = 30.66

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.xlarge"]["sp-3-all-month"] = 1208.88 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.xlarge"]["sp-3-no-month"] = 37.23 * 2
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.xlarge"]["ri-3-all-month"] = 968.00 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c5.xlarge"]["ri-3-no-month"] = 30.66 * 2

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.large"]["sp-3-all-month"] = 1089.31 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.large"]["sp-3-no-month"] = 33.35
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.large"]["ri-3-all-month"] = 915.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.large"]["ri-3-no-month"] = 29.21

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["sp-3-all-month"] = 1089.31 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["sp-3-no-month"] = 33.35 * 2
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["ri-3-all-month"] = 915.00 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.xlarge"]["ri-3-no-month"] = 29.21 * 2

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.2xlarge"]["sp-3-all-month"] = 1089.31 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.2xlarge"]["sp-3-no-month"] = 33.35 * 4
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.2xlarge"]["ri-3-all-month"] = 915.00 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.2xlarge"]["ri-3-no-month"] = 29.21 * 4

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.4xlarge"]["sp-3-all-month"] = 1089.31 * 8 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.4xlarge"]["sp-3-no-month"] = 33.35 * 8
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.4xlarge"]["ri-3-all-month"] = 915.00 * 8 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:c6a.4xlarge"]["ri-3-no-month"] = 29.21 * 8

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:i3en.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:i3en.xlarge"]["sp-3-all-month"] = 6491.16 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:i3en.xlarge"]["sp-3-no-month"] = 198.56
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:i3en.xlarge"]["ri-3-all-month"] = 5396.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:i3en.xlarge"]["ri-3-no-month"] = 172.28

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:m6a.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:m6a.large"]["sp-3-all-month"] = 1306.12 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:m6a.large"]["sp-3-no-month"] = 39.98
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:m6a.large"]["ri-3-all-month"] = 1121.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:m6a.large"]["ri-3-no-month"] = 35.76

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.xlarge"]["sp-3-all-month"] = 3403.52 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.xlarge"]["sp-3-no-month"] = 104.19
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.xlarge"]["ri-3-all-month"] = 2839.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.xlarge"]["ri-3-no-month"] = 90.59

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.2xlarge"]["sp-3-all-month"] = 3403.52 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.2xlarge"]["sp-3-no-month"] = 104.19 * 2
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.2xlarge"]["ri-3-all-month"] = 2839.00 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6a.2xlarge"]["ri-3-no-month"] = 90.59 * 2

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.16xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.16xlarge"]["sp-3-all-month"] = 60505.23 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.16xlarge"]["sp-3-no-month"] = 1860.00
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.16xlarge"]["ri-3-all-month"] = 50466.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.16xlarge"]["ri-3-no-month"] = 1610.61

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["sp-3-all-month"] = 170.82 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["sp-3-no-month"] = 5.18
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["ri-3-all-month"] = 141.00 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["ri-3-no-month"] = 4.53

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["sp-3-all-month"] = 170.82 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["sp-3-no-month"] = 5.18 * 4
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["ri-3-all-month"] = 141.00 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["ri-3-no-month"] = 4.53 * 4

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.large"]["sp-3-all-month"] = 170.82 * 8 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.large"]["sp-3-no-month"] = 5.18 * 8
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.large"]["ri-3-all-month"] = 141.00 * 8 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.large"]["ri-3-no-month"] = 4.53 * 8

		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:c6a.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:c6a.large"]["sp-3-all"] = 1080.11 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:c6a.large"]["sp-3-no"] = 33.07
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:c6a.large"]["ri-3-all"] = 906.00 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:c6a.large"]["ri-3-no"] = 28.91

		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:r6i.16xlarge"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:r6i.16xlarge"]["sp-3-all"] = 60505.23 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:r6i.16xlarge"]["sp-3-no"] = 1852.20
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:r6i.16xlarge"]["ri-3-all"] = 50466.00 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:r6i.16xlarge"]["ri-3-no"] = 1610.61

		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:t2.medium"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:t2.medium"]["sp-3-all"] = 751.61 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:t2.medium"]["sp-3-no"] = 23.12
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:t2.medium"]["ri-3-all"] = 638.00 / 36
		regionInstanceTypePrice["017820696133"]["EU (Frankfurt)"]["EUC1-BoxUsage:t2.medium"]["ri-3-no"] = 20.15

		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"] = make(map[string]map[string]float64)
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:c6a.large"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:c6a.large"]["sp-3-all"] = 880.91 / 36
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:c6a.large"]["sp-3-no"] = 26.97
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:c6a.large"]["ri-3-all"] = 770.00 / 36
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:c6a.large"]["ri-3-no"] = 24.59

		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:t2.micro"] = map[string]float64{}
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:t2.micro"]["sp-3-all"] = 136.66 / 36
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:t2.micro"]["sp-3-no"] = 4.23
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:t2.micro"]["ri-3-all"] = 115.00 / 36
		regionInstanceTypePrice["017820696133"]["US East (N. Virginia)"]["BoxUsage:t2.micro"]["ri-3-no"] = 3.65

	case "047719631469":
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.76
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2134.00 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.89

		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-all"] = 678.02 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-no"] = 20.81
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-all"] = 566.00 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-no"] = 18.03

		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"] = map[string]float64{}
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["sp-3-all"] = 678.02 * 2 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["sp-3-no"] = 20.81 * 2
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["ri-3-all"] = 566.00 * 2 / 36
		regionInstanceTypePrice["047719631469"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["ri-3-no"] = 18.03 * 2

	case "084828581472":
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Singapore)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["sp-3-all-month"] = 170.82 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["sp-3-no-month"] = 5.18 * 2
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["ri-3-all-month"] = 141.00 * 2 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["ri-3-no-month"] = 4.53 * 2

		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["sp-3-all-month"] = 170.82 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["sp-3-no-month"] = 5.18 * 4
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["ri-3-all-month"] = 141.00 * 4 / 36
		regionInstanceTypePrice["017820696133"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.medium"]["ri-3-no-month"] = 4.53 * 4

		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.76
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2134.00 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.89

		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5a.xlarge"]["sp-3-all"] = 2391.48 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5a.xlarge"]["sp-3-no"] = 73.73
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5a.xlarge"]["ri-3-all"] = 1917.00 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5a.xlarge"]["ri-3-no"] = 61.32

		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-all"] = 678.02 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-no"] = 20.81
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-all"] = 566.00 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-no"] = 18.03

		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"] = map[string]float64{}
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-all"] = 170.82 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-no"] = 5.18
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-all"] = 141.00 / 36
		regionInstanceTypePrice["084828581472"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-no"] = 4.53

	case "116981788283":
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.micro"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.micro"]["sp-3-all-month"] = 170.82 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.micro"]["sp-3-no-month"] = 5.18
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.micro"]["ri-3-all-month"] = 141.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.micro"]["ri-3-no-month"] = 4.53

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APS1-BoxUsage:t2.small"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["sp-3-all-month"] = 170.82 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["sp-3-no-month"] = 5.18 * 2
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["ri-3-all-month"] = 141.00 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.small"]["ri-3-no-month"] = 4.53 * 2

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.large"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.large"]["sp-3-all"] = 1708.20 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.large"]["sp-3-no"] = 52.56
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.large"]["ri-3-all"] = 1225.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.large"]["ri-3-no"] = 39.42

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["sp-3-all-month"] = 204.98 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["sp-3-no-month"] = 5.69
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["ri-3-all-month"] = 172.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["ri-3-no-month"] = 5.48
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.micro"]["OnDemand-day"] = 0.0152

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.xlarge"]["sp-3-all"] = 2948.62 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.xlarge"]["sp-3-no"] = 90.23
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.xlarge"]["ri-3-all"] = 2150.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.xlarge"]["ri-3-no"] = 68.62

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-all"] = 170.82 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-no"] = 5.18
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-all"] = 141.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-no"] = 4.53

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["sp-3-all"] = 170.82 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["sp-3-no"] = 5.18 * 2
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["ri-3-all"] = 141.00 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["ri-3-no"] = 4.53 * 2

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-all"] = 170.82 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-no"] = 5.18 * 4
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-all"] = 141.00 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-no"] = 4.53 * 4

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.xlarge"]["sp-3-all"] = 170.82 * 8 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.xlarge"]["sp-3-no"] = 5.18 * 8
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.xlarge"]["ri-3-all"] = 141.00 * 8 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.xlarge"]["ri-3-no"] = 4.53 * 8

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.large"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.large"]["sp-3-all"] = 1340.28 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.large"]["sp-3-no"] = 40.88
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.large"]["ri-3-all"] = 1067.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.large"]["ri-3-no"] = 34.31

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 1340.28 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 40.88 * 2
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 1067.00 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 34.31 * 2

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.4xlarge"]["sp-3-all"] = 1340.28 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.4xlarge"]["sp-3-no"] = 40.88 * 4
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.4xlarge"]["ri-3-all"] = 1067.00 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.4xlarge"]["ri-3-no"] = 34.31 * 4

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["sp-3-all"] = 31980.66 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["sp-3-no"] = 979.00
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["ri-3-all"] = 26893.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["ri-3-no"] = 858.28

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.xlarge"]["sp-3-all"] = 3179.88 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.xlarge"]["sp-3-no"] = 97.82
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.xlarge"]["ri-3-all"] = 2609.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.xlarge"]["ri-3-no"] = 83.22

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["sp-3-all"] = 3179.88 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["sp-3-no"] = 97.82 * 2
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["ri-3-all"] = 2609.00 * 2 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["ri-3-no"] = 83.22 * 2

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["sp-3-all"] = 3179.88 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["sp-3-no"] = 97.82 * 4
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["ri-3-all"] = 2609.00 * 4 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["ri-3-no"] = 83.22 * 4

		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.xlarge"]["sp-3-all"] = 4152.24 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.xlarge"]["sp-3-no"] = 127.02
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.xlarge"]["ri-3-all"] = 3300.00 / 36
		regionInstanceTypePrice["116981788283"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.xlarge"]["ri-3-no"] = 105.12

	case "180294175546":
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["sp-3-all"] = 170.82 * 2 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["sp-3-no"] = 5.18 * 2
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["ri-3-all"] = 141.00 * 2 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.small"]["ri-3-no"] = 4.53 * 2

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-all"] = 170.82 * 4 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-no"] = 5.18 * 4
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-all"] = 141.00 * 4 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-no"] = 4.53 * 4

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["sp-3-all"] = 170.82 * 8 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["sp-3-no"] = 5.18 * 8
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["ri-3-all"] = 141.00 * 8 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.large"]["ri-3-no"] = 4.53 * 8

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.large"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.large"]["sp-3-all"] = 1603.08 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.large"]["sp-3-no"] = 48.91
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.large"]["ri-3-all"] = 1304.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.large"]["ri-3-no"] = 41.61

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["sp-3-all"] = 6386.04 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["sp-3-no"] = 195.64
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["ri-3-all"] = 5217.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.2xlarge"]["ri-3-no"] = 166.44

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["sp-3-all"] = 12772.08 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["sp-3-no"] = 391.28
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["ri-3-all"] = 10434.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:m5.4xlarge"]["ri-3-no"] = 332.88

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.2xlarge"]["sp-3-all"] = 8304.48 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.2xlarge"]["sp-3-no"] = 254.04
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.2xlarge"]["ri-3-all"] = 6601.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:r5.2xlarge"]["ri-3-no"] = 210.97

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.76
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2134.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.89

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.8xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.8xlarge"]["sp-3-all"] = 31980.66 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.8xlarge"]["sp-3-no"] = 652.67
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.8xlarge"]["ri-3-all"] = 17929.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.8xlarge"]["ri-3-no"] = 572.19

		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"] = map[string]float64{}
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["sp-3-all"] = 21320.44 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["sp-3-no"] = 979.00
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["ri-3-all"] = 17929.00 / 36
		regionInstanceTypePrice["180294175546"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c6i.12xlarge"]["ri-3-no"] = 858.28

	case "324037306079":
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:m7i.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:m7i.xlarge"]["sp-3-all"] = 2438.00 / 36
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:m7i.xlarge"]["sp-3-no"] = 74.64
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:m7i.xlarge"]["ri-3-all"] = 2092.00 / 36
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:m7i.xlarge"]["ri-3-no"] = 66.76

		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:r7i.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:r7i.xlarge"]["sp-3-all"] = 3291.57 / 36
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:r7i.xlarge"]["sp-3-no"] = 100.76
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:r7i.xlarge"]["ri-3-all"] = 2745.00 / 36
		regionInstanceTypePrice["324037306079"]["US East (Ohio)"]["USE2-BoxUsage:r7i.xlarge"]["ri-3-no"] = 87.61

		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"] = map[string]float64{}
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-all"] = 170.82 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-no"] = 5.18
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-all"] = 141.00 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-no"] = 4.53

		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.2xlarge"]["sp-3-all"] = 5432.08 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.2xlarge"]["sp-3-no"] = 166.29
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.2xlarge"]["ri-3-all"] = 4526.00 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.2xlarge"]["ri-3-no"] = 144.47

		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5d.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5d.2xlarge"]["sp-3-all"] = 6018.12 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5d.2xlarge"]["sp-3-no"] = 184.69
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5d.2xlarge"]["ri-3-all"] = 4862.00 / 36
		regionInstanceTypePrice["324037306079"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5d.2xlarge"]["ri-3-no"] = 155.49

	case "535002890947":
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"] = make(map[string]map[string]float64, 0)
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.medium"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.medium"]["sp-3-all"] = 822.56 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.medium"]["sp-3-no"] = 25.19
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.medium"]["ri-3-all"] = 687.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t2.medium"]["ri-3-no"] = 21.90

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["sp-3-all"] = 735.84 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["sp-3-no"] = 22.56
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["ri-3-all"] = 538.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["ri-3-no"] = 17.16

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.large"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.large"]["sp-3-all"] = 1497.96 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.large"]["sp-3-no"] = 45.99
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.large"]["ri-3-all"] = 1057.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.large"]["ri-3-no"] = 33.58

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2995.92 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.xlarge"]["sp-3-no"] = 91.98
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2114.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.17

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.2xlarge"]["sp-3-all"] = 6018.12 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.2xlarge"]["sp-3-no"] = 184.69
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.2xlarge"]["ri-3-all"] = 4229.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.2xlarge"]["ri-3-no"] = 135.05

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.4xlarge"]["sp-3-all"] = 12036.24 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.4xlarge"]["sp-3-no"] = 369.38
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.4xlarge"]["ri-3-all"] = 8458.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5.4xlarge"]["ri-3-no"] = 270.10

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.large"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.large"]["sp-3-all"] = 1366.56 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.large"]["sp-3-no"] = 41.61
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.large"]["ri-3-all"] = 949.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.large"]["ri-3-no"] = 29.93

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.xlarge"]["sp-3-all"] = 2706.84 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.xlarge"]["sp-3-no"] = 83.22
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.xlarge"]["ri-3-all"] = 1897.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.xlarge"]["ri-3-no"] = 60.59

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.2xlarge"]["sp-3-all"] = 5439.96 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.2xlarge"]["sp-3-no"] = 166.44
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.2xlarge"]["ri-3-all"] = 3794.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.2xlarge"]["ri-3-no"] = 121.18

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.4xlarge"]["sp-3-all"] = 10879.92 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.4xlarge"]["sp-3-no"] = 332.88
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.4xlarge"]["ri-3-all"] = 7589.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c5a.4xlarge"]["ri-3-no"] = 242.36

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["sp-3-all"] = 1195.48 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["sp-3-no"] = 36.60
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["ri-3-all"] = 1004.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.large"]["ri-3-no"] = 32.04

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["sp-3-all"] = 2391.22 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["sp-3-no"] = 73.20
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["ri-3-all"] = 2008.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.xlarge"]["ri-3-no"] = 64.07

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.2xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.2xlarge"]["sp-3-all"] = 4782.44 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.2xlarge"]["sp-3-no"] = 146.39
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.2xlarge"]["ri-3-all"] = 4015.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.2xlarge"]["ri-3-no"] = 128.14

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.4xlarge"]["sp-3-all"] = 9564.88 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.4xlarge"]["sp-3-no"] = 292.78
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.4xlarge"]["ri-3-all"] = 8030.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:c6a.4xlarge"]["ri-3-no"] = 256.28

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.xlarge"]["sp-3-all"] = 3442.68 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.xlarge"]["sp-3-no"] = 105.12
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.xlarge"]["ri-3-all"] = 2451.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m5.xlarge"]["ri-3-no"] = 78.11

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.large"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.large"]["sp-3-all"] = 1382.59 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.large"]["sp-3-no"] = 42.33
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.large"]["ri-3-all"] = 1156.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.large"]["ri-3-no"] = 36.89

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.xlarge"]["sp-3-all"] = 2765.18 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.xlarge"]["sp-3-no"] = 84.65
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.xlarge"]["ri-3-all"] = 2312.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.xlarge"]["ri-3-no"] = 72.8

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.4xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.4xlarge"]["sp-3-all"] = 11060.99 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.4xlarge"]["sp-3-no"] = 338.60
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.4xlarge"]["ri-3-all"] = 9249.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.4xlarge"]["ri-3-no"] = 295.18

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.8xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.8xlarge"]["sp-3-all"] = 22121.98 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.8xlarge"]["sp-3-no"] = 677.20
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.8xlarge"]["ri-3-all"] = 18498.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:m6a.8xlarge"]["ri-3-no"] = 591.36

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Jakarta)"]["APS4-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Jakarta)"]["APS4-BoxUsage:c5.xlarge"]["sp-3-all"] = 2417.76 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Jakarta)"]["APS4-BoxUsage:c5.xlarge"]["sp-3-no"] = 62.05
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Jakarta)"]["APS4-BoxUsage:c5.xlarge"]["ri-3-all"] = 1937.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Jakarta)"]["APS4-BoxUsage:c5.xlarge"]["ri-3-no"] = 62.05

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Melbourne)"]["APS6-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Melbourne)"]["APS6-BoxUsage:c5.xlarge"]["sp-3-all"] = 2733.12 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Melbourne)"]["APS6-BoxUsage:c5.xlarge"]["sp-3-no"] = 83.22
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Melbourne)"]["APS6-BoxUsage:c5.xlarge"]["ri-3-all"] = 2194.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Melbourne)"]["APS6-BoxUsage:c5.xlarge"]["ri-3-no"] = 70.08

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Mumbai)"]["APS6-BAPS3-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Mumbai)"]["APS6-BAPS3-BoxUsage:c5.xlarge"]["sp-3-all"] = 2102.40 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Mumbai)"]["APS6-BAPS3-BoxUsage:c5.xlarge"]["sp-3-no"] = 64.24
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Mumbai)"]["APS6-BAPS3-BoxUsage:c5.xlarge"]["ri-3-all"] = 1680.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Mumbai)"]["APS6-BAPS3-BoxUsage:c5.xlarge"]["ri-3-no"] = 53.29

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Osaka)"]["APN3-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Osaka)"]["APN3-BoxUsage:c5.xlarge"]["sp-3-all"] = 3022.20 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Osaka)"]["APN3-BoxUsage:c5.xlarge"]["sp-3-no"] = 91.98
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Osaka)"]["APN3-BoxUsage:c5.xlarge"]["ri-3-all"] = 2115.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Osaka)"]["APN3-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.16

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Seoul)"]["APN2-BoxUsage:c5a.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Seoul)"]["APN2-BoxUsage:c5a.xlarge"]["sp-3-all"] = 2312.64 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Seoul)"]["APN2-BoxUsage:c5a.xlarge"]["sp-3-no"] = 70.81
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Seoul)"]["APN2-BoxUsage:c5a.xlarge"]["ri-3-all"] = 1700.00
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Seoul)"]["APN2-BoxUsage:c5a.xlarge"]["ri-3-no"] = 54.02

		regionInstanceTypePrice["535002890947"]["Asia Pacific (Sydney)"]["APS2-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Sydney)"]["APS2-BoxUsage:c5.xlarge"]["sp-3-all"] = 2733.12 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Sydney)"]["APS2-BoxUsage:c5.xlarge"]["sp-3-no"] = 83.22
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Sydney)"]["APS2-BoxUsage:c5.xlarge"]["ri-3-all"] = 2194.00 / 36
		regionInstanceTypePrice["535002890947"]["Asia Pacific (Sydney)"]["APS2-BoxUsage:c5.xlarge"]["ri-3-no"] = 70.08

		regionInstanceTypePrice["535002890947"]["EU (London)"]["EUW2-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["EU (London)"]["EUW2-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["535002890947"]["EU (London)"]["EUW2-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.03
		regionInstanceTypePrice["535002890947"]["EU (London)"]["EUW2-BoxUsage:c5.xlarge"]["ri-3-all"] = 1996.00 / 36
		regionInstanceTypePrice["535002890947"]["EU (London)"]["EUW2-BoxUsage:c5.xlarge"]["ri-3-no"] = 63.51

		regionInstanceTypePrice["535002890947"]["Middle East (UAE)"]["MEC1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["535002890947"]["Middle East (UAE)"]["MEC1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2785.68 / 36
		regionInstanceTypePrice["535002890947"]["Middle East (UAE)"]["MEC1-BoxUsage:c5.xlarge"]["sp-3-no"] = 84.68
		regionInstanceTypePrice["535002890947"]["Middle East (UAE)"]["MEC1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2087.00 / 36
		regionInstanceTypePrice["535002890947"]["Middle East (UAE)"]["MEC1-BoxUsage:c5.xlarge"]["ri-3-no"] = 66.43

	case "605134445666":
		regionInstanceTypePrice["605134445666"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"] = map[string]float64{}
		regionInstanceTypePrice["605134445666"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-all"] = 170.82 / 36
		regionInstanceTypePrice["605134445666"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["sp-3-no"] = 5.18
		regionInstanceTypePrice["605134445666"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-all"] = 141.00 / 36
		regionInstanceTypePrice["605134445666"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.micro"]["ri-3-no"] = 4.53

	case "682033488271":
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["sp-3-all"] = 735.84 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["sp-3-no"] = 22.56
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["ri-3-all"] = 538.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:t3.medium"]["ri-3-no"] = 17.16

		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:r6i.large"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:r6i.large"]["sp-3-all"] = 1890.85 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:r6i.large"]["sp-3-no"] = 57.88
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:r6i.large"]["ri-3-all"] = 1577.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Tokyo)"]["APN1-BoxUsage:r6i.large"]["ri-3-no"] = 50.33

		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["sp-3-all-month"] = 170.82 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["sp-3-no-month"] = 5.18
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["ri-3-all-month"] = 141.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t2.micro"]["ri-3-no-month"] = 4.53

		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.small"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.small"]["sp-3-all"] = 307.48 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.small"]["sp-3-no"] = 9.42
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.small"]["ri-3-all"] = 256.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.small"]["ri-3-no"] = 8.18

		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.xlarge"]["sp-3-all"] = 2454.55 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.xlarge"]["sp-3-no"] = 75.19
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.xlarge"]["ri-3-all"] = 2046.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:t3.xlarge"]["ri-3-no"] = 65.34

		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.xlarge"]["sp-3-all"] = 3781.69 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.xlarge"]["sp-3-no"] = 115.76
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.xlarge"]["ri-3-all"] = 3154.00 / 36
		regionInstanceTypePrice["682033488271"]["Asia Pacific (Singapore)"]["APS1-BoxUsage:r6i.xlarge"]["ri-3-no"] = 100.66

	case "794038255928":
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.76
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2134.00 / 36
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.89

		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"] = map[string]float64{}
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-all"] = 678.02 / 36
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["sp-3-no"] = 20.81
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-all"] = 566.00 / 36
		regionInstanceTypePrice["794038255928"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:t3.medium"]["ri-3-no"] = 18.03

	case "842675990460":
		regionInstanceTypePrice["842675990460"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"] = map[string]float64{}
		regionInstanceTypePrice["842675990460"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-all"] = 2654.28 / 36
		regionInstanceTypePrice["842675990460"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["sp-3-no"] = 81.76
		regionInstanceTypePrice["842675990460"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-all"] = 2134.00 / 36
		regionInstanceTypePrice["842675990460"]["Asia Pacific (Hong Kong)"]["APE1-BoxUsage:c5.xlarge"]["ri-3-no"] = 67.89
	}

}
