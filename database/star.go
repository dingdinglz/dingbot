package database

func StarExist(repository string, sender string) bool {
	var i int64
	DB.Model(&StarTable{}).Where("repository = ?", repository).Where("sender = ?", sender).Count(&i)
	return i != 0
}

func StarAdd(repository string, sender string) {
	var i StarTable
	i.Repository = repository
	i.Sender = sender
	DB.Create(&i)
}
