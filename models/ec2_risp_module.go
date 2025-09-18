package models

import (
	"regexp"
	"strconv"

	"analysis-aws-bill/config"
	"analysis-aws-bill/database"
)

func getEc2PriceFromDetailedBill() []database.Ec2Price {
	db := database.ContentMysql()

	detailedList := []database.DetailedBill{}
	ec2List := []database.Ec2Price{}

	//select usage_account_id,location,usage_type,SUM(un_blended_cost) from detailed_bill WHERE product_name = 'Amazon Elastic Compute Cloud' and usage_type like '%BoxUsage%' and un_blended_cost > 0 GROUP BY usage_account_id,location,usage_type

	db.Model(&database.DetailedBill{}).
		Select("usage_account_id,location,usage_type,item_description,SUM(un_blended_cost) as cost").
		Where("product_name = 'Amazon Elastic Compute Cloud' and usage_type like '%BoxUsage%' and un_blended_cost > 0").
		Group("usage_account_id,location,usage_type,item_description").
		Scan(&detailedList)

	for _, values := range detailedList {

		info := config.GetEC2Pricing(values.UsageType)

		if info == nil {
			continue
		}

		priceRegex := regexp.MustCompile(`\$(\d+\.?\d*)`)
		ondemandCostList := priceRegex.FindStringSubmatch(values.ItemDescription)

		var price float64
		var err error
		if len(ondemandCostList) > 0 {
			priceStr := ondemandCostList[1]
			price, err = strconv.ParseFloat(priceStr, 64)
			if err != nil {
				price = 0
			}
		}

		ec2Price := database.Ec2Price{
			UsageAccountId:    values.UsageAccountId,
			Location:          values.Location,
			InstanceType:      info.InstanceType,
			SP3YearAllUpfront: info.SP3YearAllUpfront,
			SP3YearNoUpfront:  info.SP3YearNoUpfront,
			RI3YearAllUpfront: info.RI3YearAllUpfront,
			RI3YearNoUpfront:  info.RI3YearNoUpfront,
			OnDemand:          price,
		}

		ec2List = append(ec2List, ec2Price)

	}

	return ec2List

}

func InsertEc2PriceToDB() {
	ec2PriceList := getEc2PriceFromDetailedBill()

	db := database.ContentMysql()

	db.Create(&ec2PriceList)

}
