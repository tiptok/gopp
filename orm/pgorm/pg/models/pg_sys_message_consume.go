package models

// SysMessageConsume
type SysMessageConsume struct {
	//	消息ID
	Id int64 `gorm:"primaryKey"`
	//	主题
	Topic string
	//	分区信息
	Partition int
	//	消息偏移序号
	Offset int64
	//	键值
	Key string
	//	消息内容
	Value string
	//	消息时间
	MsgTime int64
	//	创建时间
	CreateAt int64
	//	状态
	Status int64
}

func (m *SysMessageConsume) TableName() string {
	return "sys_message_consume"
}
