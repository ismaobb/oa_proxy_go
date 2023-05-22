package repositories

type OrderAccessory struct {
	JobOrder  string `gorm:"primaryKey;column:TZDBH" json:"job_order"`
	Accessory string `gorm:"column:MXMC" json:"accessory"`
	Category  string `gorm:"column:PACKAGEMC" json:"category"`
	Count     int    `gorm:"column:MXNUM" json:"count"`
}

func (OrderAccessory) TableName() string {
	return "T_SCTZD_PACKAGE"
}
