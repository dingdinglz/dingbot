package database

func GithubWebhookExist(name string) bool {
	var i int64
	DB.Model(&GithubHookTable{}).Where("name = ?", name).Count(&i)
	return i != 0
}

func GithubWebhookCreate(name string, group string) {
	var i GithubHookTable
	if GithubWebhookExist(name) {
		DB.Model(&GithubHookTable{}).Where("name = ?", name).First(&i)
	}
	i.Name = name
	i.Group = group
	DB.Save(&i)
}

func GithubWebhookGet(name string) string {
	var i GithubHookTable
	DB.Model(&GithubHookTable{}).Where("name = ?", name).First(&i)
	return i.Group
}
