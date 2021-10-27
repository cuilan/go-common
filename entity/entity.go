package entity

// 数据库: weattech_dl

func (DlService) TableName() string {
	return "t_dl_service"
}

// DlService 下载服务项表
type DlService struct {
	Id         int64  `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
	Name       string `gorm:"name" json:"name"`                    //服务名称
	Desc       string `gorm:"column:desc" json:"desc"`             //服务描述
	DeployIp   string `gorm:"column:deploy_ip" json:"deploy_ip"`   //服务部署主机
	Url        string `gorm:"column:url" json:"url"`               //请求URL地址
	DlType     int    `gorm:"column:dl_type" json:"dl_type"`       //下载类型，1-FTP，2-HTTP
	SavePath   string `gorm:"column:save_path" json:"save_path"`   //数据存放目录
	StoreType  int    `gorm:"column:store_type" json:"store_type"` //存储类型，1-数据库，2-硬盘，3-oss
	Remark     string `gorm:"column:remark" json:"remark"`         //备注
	Action     int    `gorm:"column:action" json:"action"`
	ActionTime int64  `gorm:"column:action_time" json:"action_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

func (DlRecords) TableName() string {
	return "t_dl_records"
}

// DlRecords 下载记录表
type DlRecords struct {
	Id         int64  `gorm:"column:id" json:"id"`
	Name       string `gorm:"name" json:"name"`       //服务名称
	Status     int    `gorm:"status" json:"status"`   //下载状态
	DlTime     int64  `gorm:"dl_time" json:"dl_time"` //下载时间
	DlDate     string `gorm:"dl_date" json:"dl_date"` //下载日期，yyyy-MM-dd HH:mm:ss
	Desc       string `gorm:"dsesc" json:"desc"`      //下载详情，错误描述
	Action     int    `gorm:"column:action" json:"action"`
	ActionTime int64  `gorm:"column:action_time" json:"action_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}
