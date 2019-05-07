package models

// Device 仪表model
type Device struct {
	BaseModel
	Name        string `gorm:"column:name;not null;unique;type:varchar(47);index" json:"name"`
	Description string `gorm:"column:description;default:'';type:text(200)" json:"description"`
	GatherType  string `gorm:"column:gather_type;not null;type:varchar(64)" json:"gather_type"`
	AreaID      uint   `gorm:"column:gather_type;not null;" json:"area_id"`
	GatherWay   string `gorm:"column:gather_way;not null;type:varchar(64)" json:"gather_way"`
	GatewayID   uint   `gorm:"column:gather_type;not null;" json:"gateway_id"`
	DeviceID    string `gorm:"column:device_id;not null;type:varchar(36)" json:"device_id"`
	DeviceIntId int32  `gorm:"column:device_int_id;not null;default:0" json:"device_int_id"`
	//	# 最大负载（告警负载）（不于小零的整数，为0表示不告警）
	MaxLoad                  float32 `gorm:"column:max_load;not null;default:0;" json:"max_load"`
	WaitLoad                 bool    `gorm:"column:wait_load;not null;default:false;not null;" json:"wait_load"`
	IsHeadMeter              bool    `gorm:"column:is_head_meter;not null;default:false" json:"is_head_meter"`
	Type                     string  `gorm:"column:type;not null;type:varchar(64)" json:"type"`
	ProductId                string  `gorm:"column:product_id;not null;default:'';type:varchar(64)" json:"product_id"`
	IsShow                   bool    `gorm:"column:is_show;not null;default:true" json:"is_show"`
	ManualReadmeterThreshold float32 `gorm:"column:manual_readmeter_threshold;not null;default:0;" json:"manual_readmeter_threshold"`
	BusinessRulesId          int32   `gorm:"column:business_rules_id;not null;default:0" json:"business_rules_id"`
	ModbusIndex              int32   `gorm:"column:modbus_index;default:0" json:"modbus_index"`
	RxcelLocation            string  `gorm:"column:excel_location;default:'';unique;type:varchar(32)" json:"excel_location"`
	WorkLoad                 int32   `gorm:"column:work_load;not null;default:0" json:"work_load"`
	LoadValue                float32 `gorm:"column:load_value;not null;default:0" json:"load_value"`
	Magnification            float32 `gorm:"column:magnification;not null;default:1" json:"magnification"`
	Blocked                  int32   `gorm:"column:sort_num;not null;default:0" json:"sort_num" valid:"Range(0,1)"`
	SortNum                  int32   `gorm:"column:blocked;default:0" json:"blocked"`
	NormalLoad               float32 `gorm:"column:normal_load;default:0" json:"normal_load"`
}
