package database

// MessageInsert 把数据存表，以备查询用。
// typename 类型名group或者private
// messageText 消息内容
// messageSeq
func MessageInsert(typename string, messageText string, messageSeq uint32) {
	var i MessageTable
	i.Typename = typename
	i.Message = messageText
	i.Seq = messageSeq
	DB.Create(&i)
}

func MessageGetText(typename string, messageSeq uint32) string {
	var i MessageTable
	DB.Model(&MessageTable{}).Where("typename = ?", typename).Where("seq = ?", messageSeq).First(&i)
	return i.Message
}
