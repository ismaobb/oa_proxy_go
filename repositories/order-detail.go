package repositories

type OrderDetail struct {
	JobOrder        string  `gorm:"primaryKey;column:TZDBH;type:varchar(16);comment:'通知单号'" json:"job_order"`
	Index           int     `gorm:"column:PH;type:int;comment:'编号'" json:"index"`
	Type            string  `gorm:"column:XH;type:varchar(10);comment:'型号'" json:"type"`
	Thickness       string  `gorm:"column:GGDEEP;type:varchar(6);comment:'厚度'" json:"thickness"`
	Width           int     `gorm:"column:GGWIDTH;type:int;comment:'宽度'" json:"width"`
	Length          int     `gorm:"column:GGLENGTH;type:int;comment:'长度'" json:"length"`
	Corona          string  `gorm:"column:TREATMENT;type:varchar(6);comment:'电晕'" json:"corona"`
	Weight          float64 `gorm:"column:WEIGHT;type:decimal;comment:'重量'" json:"weight"`
	Rolls           int     `gorm:"column:ROLL;type:int;comment:'总卷数'" json:"rolls"`
	SchemeRolls     float64 `gorm:"column:RemainderRollCountPre;type:decimal;comment:'剩余卷数'" json:"scheme_rolls"`
	ActualLength    float64 `gorm:"column:SJLENGTH;type:decimal;comment:'实际米数'" json:"actual_length"`
	ActualThickness float64 `gorm:"column:SJDEEP;type:decimal;comment:'实际厚度'" json:"actual_thickness"`
}

func (OrderDetail) TableName() string {
	return "T_SCTZD_WXMX"
}
