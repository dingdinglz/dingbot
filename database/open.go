package database

func OpenHave(Uin uint32, TypeName string) bool {
	var i int64
	DB.Model(&OpenTable{}).Where("typename = ?", TypeName).Where("uin = ?", Uin).Count(&i)
	return i != 0
}

func OpenAdd(Uin uint32, TypeName string) {
	if OpenHave(Uin, TypeName) {
		return
	}
	var i OpenTable
	i.Uin = Uin
	i.Typename = TypeName
	DB.Create(&i)
}

func OpenDelete(Uin uint32, TypeName string) {
	DB.Delete(&OpenTable{}, "uin = ? AND typename = ?", Uin, TypeName)
}
