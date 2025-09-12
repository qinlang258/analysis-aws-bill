package database

// DetailedBill represents the structure of the `detailed_bill` table in the database
type DetailedBill struct {
	ID                int     `gorm:"primaryKey;autoIncrement;column:id"`
	UsageAccountId    string  `gorm:"type:varchar(255);column:usage_account_id;default:null"`
	Location          string  `gorm:"type:varchar(255);column:location;default:null"`
	ProductName       string  `gorm:"type:varchar(255);column:product_name;default:null"`
	UsageType         string  `gorm:"type:varchar(255);column:usage_type;default:null"`
	Operation         string  `gorm:"type:varchar(255);column:operation;default:null"`
	AvailabilityZone  string  `gorm:"type:varchar(255);column:availability_zone;default:null"`
	ReservedInstance  string  `gorm:"type:varchar(255);column:reserved_instance;default:null"`
	ItemDescription   string  `gorm:"type:varchar(255);column:item_description;default:null"`
	UsageStartDate    string  `gorm:"type:varchar(255);column:usage_start_date;default:null"`
	UsageEndDate      string  `gorm:"type:varchar(255);column:usage_end_date;default:null"`
	UsageQuantity     string  `gorm:"type:varchar(255);column:usage_quantity;default:null"`
	BlendedCost       float64 `gorm:"type:decimal(10,5);column:blended_cost;default:null"`
	UnBlendedCost     float64 `gorm:"type:decimal(10,5);column:un_blended_cost;default:null"`
	ResourceID        string  `gorm:"type:varchar(255);column:resource_id;default:null"`
	GroupName         string  `gorm:"type:varchar(255);column:group_name;default:null"`
	InstanceGroupRole string  `gorm:"type:varchar(255);column:instance_group_role;default:null"`
	JobFlowID         string  `gorm:"type:varchar(255);column:job_flow_id;default:null"`
	Name              string  `gorm:"type:varchar(255);column:name;default:null"`
	Business          string  `gorm:"type:varchar(255);column:business;default:null"`
	Department        string  `gorm:"type:varchar(255);column:department;default:null"`
}

// TableName specifies the table name for the DetailedBill model
func (DetailedBill) TableName() string {
	return "detailed_bill"
}
