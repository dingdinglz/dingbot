package database

func KeyWordHave(key string) bool {
	var i int64
	DB.Model(&KeyWordTable{}).Where("key = ?", key).Count(&i)
	return i != 0
}

func KeyWordDelete(key string) {
	if !KeyWordHave(key) {
		return
	}
	DB.Delete(&KeyWordTable{}, "key = ?", key)
}

func KeyWordCreate(key string, text string, precise bool) {
	if KeyWordHave(key) {
		KeyWordDelete(key)
	}
	var i KeyWordTable
	i.Key = key
	i.Text = text
	if precise {
		i.Precise = "true"
	} else {
		i.Precise = "false"
	}
	DB.Create(&i)
}

func KeyWordGet(key string) KeyWordTable {
	var i KeyWordTable
	DB.Model(&KeyWordTable{}).Where("key = ?", key).First(&i)
	return i
}

func KeyWordGetAll() []KeyWordTable {
	var i []KeyWordTable
	DB.Model(&KeyWordTable{}).Find(&i)
	return i
}
