package repositories

type AccessoryColumn struct {
	PkgID     string `gorm:"primaryKey;column:PACKAGEID" json:"pkg_id"`
	Accessory string `gorm:"column:MXMC" json:"accessory"`
	Category  string `gorm:"column:PACKAGEMC" json:"category"`
}

func (AccessoryColumn) TableName() string {
	return "T_CODES_PACKAGE"
}
