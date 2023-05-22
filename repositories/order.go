package repositories

type Order struct {
	JobOrder      string  `gorm:"column:TZDBH;type:varchar(16);comment:'通知单号'" json:"job_order"`
	Customer      string  `gorm:"column:KHDM;type:varchar(20);comment:'客户代码'" json:"customer"`
	Salesman      string  `gorm:"column:HTYWY;type:varchar(20);comment:'业务员'" json:"salesman"`
	Memo          string  `gorm:"column:BZ;type:varchar(100);comment:'备注'" json:"memo"`
	WrapperMemo   string  `gorm:"column:BZFBZ;type:varchar(100);comment:'包装格式备注'" json:"wrapper_memo"`
	Wrapper       string  `gorm:"column:ZXWCB;comment:'包装材料'" json:"wrapper"`
	PlateType     string  `gorm:"column:ZXCBLX;comment:'铲板类型'" json:"plate_type"`
	PlateSize     string  `gorm:"column:ZXCBZL;comment:'铲板尺寸'" json:"plate_size"`
	PlateWrapper  string  `gorm:"column:ZXCBFS;comment:'铲板包装方式'" json:"plate_wrapper"`
	RollPlacement string  `gorm:"column:ZXCBFMJ;comment:'膜卷摆放'" json:"roll_placement"`
	Other         string  `gorm:"column:ZXQT;comment:'其他'" json:"other"`
	Packing       string  `gorm:"column:ZXYQ;comment:'装箱要求'" json:"packing"`
	DealDate      string  `gorm:"column:SXRQ;comment:'生效日期'" json:"deal_date"`
	TradeMode     string  `gorm:"column:XSFS;comment:'销售方式'" json:"trade_mode"`
	DeliveryTime  string  `gorm:"column:JHRQ;comment:'发货日期'" json:"delivery_time"`
	SchemeDate    string  `gorm:"column:SCBD_SXRQ;comment:'排单生效日期'" json:"scheme_date"`
	ProduceDate   string  `gorm:"column:SCSXRQ;comment:'生产生效日期'" json:"produce_date"`
	Density       float64 `gorm:"column:JSFSMD;comment:'密度'" json:"density"`
	State         string  `gorm:"column:SCBDZT;comment:'数据拉取标识'" json:"state"`
	Factory       string  `gorm:"column:GCBH;comment:'所属工厂'" json:"factory"`
	Quality       string  `gorm:"column:ZLQT;comment:'质量标准';type:varchar(256)" json:"quality"`
	Level         string  `gorm:"column:ZLDJYQ;comment:'客户等级';type:varchar(128)" json:"level"`
}

func (Order) TableName() string {
	return "T_SCTZD_WXZB"
}
