package database

// DetailedBill represents the structure of the `detailed_bill` table in the database
type Ec2Price struct {
	ID                int     `gorm:"primaryKey;autoIncrement;column:id"`
	UsageAccountId    string  `gorm:"type:varchar(255);column:usage_account_id;default:null"`
	Location          string  `gorm:"type:varchar(255);column:location;default:null"`
	InstanceType      string  `gorm:"type:varchar(255);column:instance_type;default:null"`
	SP3YearAllUpfront float64 `gorm:"type:decimal(10,5);column:sp3_year_all_upfront;default:null"`
	SP3YearNoUpfront  float64 `gorm:"type:decimal(10,5);column:sp3_year_no_upfront;default:null"`
	RI3YearAllUpfront float64 `gorm:"type:decimal(10,5);column:ri3_year_all_upfront;default:null"`
	RI3YearNoUpfront  float64 `gorm:"type:decimal(10,5);column:ri3_year_no_upfront;default:null"`
	OnDemand          float64 `gorm:"type:decimal(10,5);column:on_demand;default:null"`
}

// TableName specifies the table name for the DetailedBill model
func (Ec2Price) TableName() string {
	return "ec2_price"
}
