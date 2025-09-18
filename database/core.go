package database

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"analysis-aws-bill/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ContentMysql() *gorm.DB {
	//读取.ini里面的数据库配置

	//dsn := fmt.Sprintf("root:ql2252528@tcp(127.0.0.1:3306)/bill202502?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //打印sql
		SkipDefaultTransaction: true,
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get SQL DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(6532)
	sqlDB.SetMaxIdleConns(10)

	return DB
}

func WriteToMysql(ctx context.Context, fileName string) error {
	// 打开数据库连接
	db := ContentMysql()

	// 只需要在程序开始时进行一次 AutoMigrate
	err := db.AutoMigrate(&DetailedBill{})
	if err != nil {
		return err
	}

	// 打开 CSV 文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open CSV file: %v", err)
		return err
	}
	defer file.Close()

	// 解析 CSV 文件
	reader := csv.NewReader(file)
	records, err := reader.Read()
	if err != nil {
		log.Fatalf("failed to read CSV header: %v", err)
		return err
	}

	// 确认自己需要的字段
	tagNumber := make(map[int]string)
	for index, record := range records {
		switch record {
		case "lineItem/UsageAccountId":
			tagNumber[index] = "usage_account_id"
		case "product/location":
			tagNumber[index] = "location"
		case "product/ProductName":
			tagNumber[index] = "product_name"
		case "UsageQuantity":
			tagNumber[index] = "usage_quantity"
		case "lineItem/UsageType":
			tagNumber[index] = "usage_type"
		case "lineItem/Operation":
			tagNumber[index] = "operation"
		case "AvailabilityZone":
			tagNumber[index] = "availability_zone"
		case "ReservedInstance":
			tagNumber[index] = "reserved_instance"
		case "lineItem/LineItemDescription":
			tagNumber[index] = "item_description"
		case "UsageStartDate":
			tagNumber[index] = "usage_start_date"
		case "UsageEndDate":
			tagNumber[index] = "usage_end_date"
		case "lineItem/BlendedCost":
			tagNumber[index] = "blended_cost"
		case "lineItem/UnblendedCost":
			tagNumber[index] = "un_blended_cost"
		case "lineItem/ResourceId":
			tagNumber[index] = "resource_id"
		case "aws:autoscaling:groupName":
			tagNumber[index] = "group_name"
		case "user:Name":
			tagNumber[index] = "name"
		case "user:department":
			tagNumber[index] = "department"
		case "user:business":
			tagNumber[index] = "business"
		case "aws:elasticmapreduce:instance-group-role":
			tagNumber[index] = "instance_group_role"
		case "aws:elasticmapreduce:job-flow-id":
			tagNumber[index] = "job_flow_id"
		}
	}

	// 存储待插入的 records
	var bills []*DetailedBill
	batchSize := 1000

	for {
		// 逐行读取数据
		record, err := reader.Read()
		if err != nil {
			break // 到达文件末尾时退出
		}

		bill := &DetailedBill{}
		for index, value := range record {
			switch tagNumber[index] {
			case "usage_account_id":
				bill.UsageAccountId = value
			case "location":
				bill.Location = value
			case "product_name":
				bill.ProductName = value
			case "usage_quantity":
				bill.UsageQuantity = value
			case "usage_type":
				bill.UsageType = value
			case "operation":
				bill.Operation = value
			case "availability_zone":
				bill.AvailabilityZone = value
			case "reserved_instance":
				bill.ReservedInstance = value
			case "item_description":
				bill.ItemDescription = value
			case "usage_start_date":
				bill.UsageStartDate = value
			case "usage_end_date":
				bill.UsageEndDate = value
			case "blended_cost":
				bill.BlendedCost, _ = strconv.ParseFloat(value, 64)
			case "un_blended_cost":
				bill.UnBlendedCost, _ = strconv.ParseFloat(value, 64)
			case "resource_id":
				bill.ResourceID = value
			case "group_name":
				bill.GroupName = value
			case "name":
				bill.Name = value
			case "department":
				bill.Department = value
			case "business":
				bill.Business = value
			case "instance_group_role":
				bill.InstanceGroupRole = value
			case "job_flow_id":
				bill.JobFlowID = value
			}
		}

		// 将当前的 bill 添加到批次中
		bills = append(bills, bill)

		// 每读取 1000 条数据插入一次
		if len(bills) >= batchSize {
			if err := insertBatch(db, bills); err != nil {
				return err
			}
			// 清空 bills 数组以便继续存储下一批数据
			bills = nil
		}
	}

	// 如果有剩余的记录没有插入，插入它们
	if len(bills) > 0 {
		if err := insertBatch(db, bills); err != nil {
			return err
		}
	}

	return nil
}

// 批量插入数据
func insertBatch(db *gorm.DB, bills []*DetailedBill) error {
	if err := db.Table("detailed_bill").Create(&bills).Error; err != nil {
		log.Printf("failed to insert batch data: %v", err)
		return err
	}
	return nil
}
