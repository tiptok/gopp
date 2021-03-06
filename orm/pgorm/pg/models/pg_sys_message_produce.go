package models

// SysMessageProduce
type SysMessageProduce struct {
	//	消息ID
	Id int64 `gorm:"primaryKey"`
	//	主题
	Topic string
	//	分区信息
	Partition int
	//	消息内容
	Value string
	//	消息时间
	MsgTime int64
	//	状态
	Status int64
}

func (m *SysMessageProduce) TableName() string {
	return "sys_message_produce"
}
